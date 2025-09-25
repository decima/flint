package contracts

import (
	"errors"
	"flint/security"
	"flint/service/model"
)

type UsersManagerInterface interface {
	CreateUser(username string, password string, role security.Role) error
	DeleteUser(username string) error
	GetUser(username string) (model.User, error)
	ListUsers() ([]model.User, error)
}

var NotFoundUserErr = errors.New("User not found")
