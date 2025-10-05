package servers

import (
	"errors"
	"flint/service/contracts"
	"flint/service/model"
	"flint/service/storage"
	"flint/utils/stringutils"
	"fmt"
	"log"
)

type Manager struct {
	ServerStorage *ServerStorage
	serverActions contracts.ServerActions
}

func (m *Manager) Validate(server model.Server) error {
	if m.serverActions == nil {
		return errors.New("remote action is not initialized")
	}
	version, err := m.serverActions.DockerVersion(server)
	if err != nil {
		return fmt.Errorf("failed to get docker version: %w", err)
	}
	if version == "" {
		return contracts.InvalidServerErr
	}
	log.Println("THE VERSION IS...", version)
	return nil
}

func (m *Manager) CreateServer(name string, hostOrIp string, port int, user string, workDir string, sshKey string, keyPass string, password string) (model.Server, error) {
	if !stringutils.IsValidSlug(name) {
		return model.Server{}, contracts.BadServerNameErr
	}
	server := model.Server{
		Host:     hostOrIp,
		Username: user,
		Port:     port,
		Key:      sshKey,
		KeyPass:  keyPass,
		Password: password,
		WorkDir:  workDir,
	}
	if err := m.Validate(server); err != nil {
		return model.Server{}, fmt.Errorf("server validation failed: %w", err)
	}

	err := m.ServerStorage.Transaction(func(servers *contracts.ServerCollection, err error) error {
		if *servers == nil {
			*servers = map[string]model.Server{}
		}
		if err != nil && !errors.Is(err, storage.NotFoundErr) {
			return fmt.Errorf("failed to load servers: %w", err)
		}

		if _, exists := (*servers)[name]; exists {
			return contracts.DuplicateServerErr
		}

		(*servers)[name] = server

		if err := m.ServerStorage.Set(*servers); err != nil {
			return err
		}
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

func NewManager(serverStorage *ServerStorage, serverActions contracts.ServerActions) *Manager {
	return &Manager{
		ServerStorage: serverStorage,
		serverActions: serverActions,
	}
}
