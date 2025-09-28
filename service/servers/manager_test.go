package servers

import (
	"errors"
	"flint/service/contracts"
	"flint/service/storage"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockStorage is an in-memory implementation of the storage.Storage interface for testing.
type MockStorage struct {
	data contracts.ServerCollection
	err  error
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		data: make(contracts.ServerCollection),
	}
}

func (m *MockStorage) Load(entity *contracts.ServerCollection) error {
	if m.err != nil {
		return m.err
	}
	*entity = m.data
	return nil
}

func (m *MockStorage) Save(entity contracts.ServerCollection) error {
	if m.err != nil {
		return m.err
	}
	m.data = entity
	return nil
}

func (m *MockStorage) Get() (contracts.ServerCollection, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.data, nil
}

func (m *MockStorage) Set(entity contracts.ServerCollection) error {
	return m.Save(entity)
}

func (m *MockStorage) Transaction(tx func(entity *contracts.ServerCollection, loadError error) error) error {
	if m.err != nil && !errors.Is(m.err, storage.NotFoundErr) {
		return m.err
	}

	// The transaction gets a pointer to the data and the load error
	err := tx(&m.data, m.err)
	if err != nil {
		return err
	}

	return nil
}

func TestCreateServer(t *testing.T) {
	// Setup
	mockStorage := NewMockStorage()
	serverStorage := &ServerStorage{Storage: mockStorage}
	manager := NewManager(serverStorage)

	serverName := "test-server"
	host := "localhost"
	port := 22
	user := "testuser"
	key := "ssh-key"

	// --- Test Case 1: Successful server creation ---
	t.Run("should create a server successfully", func(t *testing.T) {
		// Execute
		createdServer, err := manager.CreateServer(serverName, host, port, user, key)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, host, createdServer.Host)
		assert.Equal(t, port, createdServer.Port)
		assert.Equal(t, user, createdServer.Username)
		assert.Equal(t, key, createdServer.Key)

		// Verify data in mock storage
		persistedServer, exists := mockStorage.data[serverName]
		require.True(t, exists, "server should exist in storage")
		assert.Equal(t, createdServer, persistedServer)
	})

	// --- Test Case 2: Duplicate server creation ---
	t.Run("should return an error if the server already exists", func(t *testing.T) {
		// Execute
		_, err := manager.CreateServer(serverName, host, port, user, key)

		// Assert
		require.Error(t, err)
		assert.Equal(t, contracts.DuplicateServerErr, err)
	})

	// --- Test Case 3: Invalid server name ---
	t.Run("should return an error for an invalid server name", func(t *testing.T) {
		// Execute
		_, err := manager.CreateServer("invalid name", host, port, user, key)

		// Assert
		require.Error(t, err)
		assert.Equal(t, contracts.BadServerNameErr, err)
	})
}