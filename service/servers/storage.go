package servers

import (
	"flint/config"
	"flint/service/contracts"
	"flint/service/storage"
)

type ServerStorage struct {
	storage.Storage[contracts.ServerCollection]
}

func NewServerStorage(config *config.Config) *ServerStorage {
	initialStorage := storage.CreateStorage[contracts.ServerCollection](config, "servers")
	return &ServerStorage{Storage: initialStorage}
}
