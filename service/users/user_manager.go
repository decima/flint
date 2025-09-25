package users

import (
	"errors"
	"flint/security"
	"flint/service/contracts"
	"flint/service/model"
	"flint/service/storage"
)

type Manager struct {
	PasswordHasher security.PasswordHasherInterface
	userStorage    *UserStorage
}

func (u *Manager) UserExists(username string) bool {
	users, err := u.ListUsers()
	if err != nil {
		return false
	}
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}

func (u *Manager) CreateUser(username string, password string, role security.Role) error {
	users, err := u.userStorage.Get()
	if err != nil && !errors.Is(err, storage.NotFoundErr) {
		return err
	}
	for _, user := range users {
		if user.Username == username {
			return errors.New("User already exists")
		}
	}

	passwordHashed, err := u.PasswordHasher.Hash(password)
	if err != nil {
		return err
	}
	users = append(users, model.User{
		Username:       username,
		Role:           role,
		HashedPassword: passwordHashed,
	})
	return u.userStorage.Set(users)
}

func (u *Manager) UpdateUser(username string, password string, role security.Role) error {
	user, err := u.GetUser(username)
	if err != nil {
		return err
	}
	if password != "" {
		newPassword, err := u.PasswordHasher.Hash(password)
		if err != nil {
			return err
		}
		user.HashedPassword = newPassword
	}

	if role != "" {
		user.Role = role
	}

	return nil
}

func (u *Manager) DeleteUser(username string) error {
	//TODO implement me
	panic("implement me")
}

func (u *Manager) GetUser(username string) (model.User, error) {
	users, err := u.ListUsers()
	if err != nil {
		return model.User{}, err
	}
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return model.User{}, contracts.NotFoundUserErr
}

func (u *Manager) ListUsers() ([]model.User, error) {
	users, err := u.userStorage.Get()
	if err != nil && !errors.Is(err, storage.NotFoundErr) {
		return nil, err
	}
	return users, nil
}

func NewUserManager(passwordHasher security.PasswordHasherInterface, userStorage *UserStorage) *Manager {
	return &Manager{passwordHasher, userStorage}
}
