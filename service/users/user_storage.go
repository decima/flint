package users

import (
	"flint/config"
	"flint/service/model"
	"flint/service/storage"
)

type UserStorage struct {
	storage.Storage[[]model.User]
}

func NewUserStorage(config *config.Config) *UserStorage {
	initialStorage := storage.CreateStorage[[]model.User](config, "users")
	return &UserStorage{Storage: initialStorage}
}
