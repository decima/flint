package servers

import (
	"errors"
	"flint/service/contracts"
	"flint/service/model"
	"flint/service/storage"
	"flint/utils/stringutils"
	"fmt"
)

type Manager struct {
	ServerStorage *ServerStorage
}

func (m *Manager) CreateServer(name string, hostOrIp string, port int, user string, sshKey string) (model.Server, error) {
	if !stringutils.IsValidSlug(name) {
		return model.Server{}, contracts.BadServerNameErr
	}

	var server model.Server
	err := m.ServerStorage.Transaction(func(servers *contracts.ServerCollection, err error) error {
		if err != nil && !errors.Is(err, storage.NotFoundErr) {
			return fmt.Errorf("failed to load servers: %w", err)
		}
		if *servers == nil {
			*servers = make(contracts.ServerCollection)
		}

		if _, exists := (*servers)[name]; exists {
			return contracts.DuplicateServerErr
		}

		server = model.Server{
			Host:     hostOrIp,
			Username: user,
			Port:     port,
			Key:      sshKey,
		}
		(*servers)[name] = server

		return nil
	})

	if err != nil {
		return model.Server{}, err
	}

	return server, nil
}

func (m *Manager) DeleteServer(name string) error {
	return m.ServerStorage.Transaction(func(servers *contracts.ServerCollection, err error) error {
		if err != nil {
			return fmt.Errorf("failed to load servers: %w", err)
		}

		if _, exists := (*servers)[name]; !exists {
			return contracts.ServerNotFoundErr
		}

		delete(*servers, name)
		return nil
	})
}

func (m *Manager) GetServer(name string) (model.Server, error) {
	servers, err := m.ServerStorage.Get()
	if err != nil {
		return model.Server{}, fmt.Errorf("failed to load servers: %w", err)
	}

	server, exists := servers[name]
	if !exists {
		return model.Server{}, contracts.ServerNotFoundErr
	}

	return server, nil
}

func (m *Manager) ListServers() (contracts.ServerCollection, error) {
	return m.ServerStorage.Get()
}

func NewManager(serverStorage *ServerStorage) *Manager {
	return &Manager{ServerStorage: serverStorage}
}
