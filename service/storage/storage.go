package storage

import (
	"flint/config"
	"fmt"
)

type Storage[T any] interface {
	Load(entity *T) error
	Save(entity T) error

	Get() (T, error)
	Set(entity T) error
	Transaction(func(entity *T, loadError error) error) error
}

var NotFoundErr = fmt.Errorf("Entity not found")

func CreateStorage[T any](config *config.Config, domain string) Storage[T] {
	return NewFileYamlStorage[T](config.StoragePath, domain)
}
