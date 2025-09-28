package users

import (
	"errors"
	"flint/service/contracts"
	"flint/service/model"
	"flint/service/storage"
	"testing"
)

// mockPasswordHasher is a mock implementation of the PasswordHasherInterface for testing.
type mockPasswordHasher struct {
	hash func(password string) (string, error)
}

func (m *mockPasswordHasher) Hash(password string) (string, error) {
	if m.hash != nil {
		return m.hash(password)
	}
	return "hashed_" + password, nil
}

func (m *mockPasswordHasher) Verify(password, hash string) bool {
	return "hashed_"+password == hash
}

// mockUserStorage is a mock implementation of the UserStorage for testing.
type mockUserStorage struct {
	users []model.User
	err   error
}

func (m *mockUserStorage) Get() ([]model.User, error) {
	if m.err != nil {
		return nil, m.err
	}
	if m.users == nil {
		return nil, storage.NotFoundErr
	}
	// Return a copy to prevent modification of the mock's internal state
	usersCopy := make([]model.User, len(m.users))
	copy(usersCopy, m.users)
	return usersCopy, nil
}

func (m *mockUserStorage) Set(users []model.User) error {
	if m.err != nil {
		return m.err
	}
	// Store a copy to prevent modification of the caller's slice
	usersCopy := make([]model.User, len(users))
	copy(usersCopy, users)
	m.users = usersCopy
	return nil
}

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		storage := &mockUserStorage{}
		hasher := &mockPasswordHasher{}
		manager := NewUserManager(hasher, storage)

		err := manager.CreateUser("testuser", "password", "user")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		users, _ := storage.Get()
		if len(users) != 1 {
			t.Fatalf("expected 1 user, got %d", len(users))
		}
		if users[0].Username != "testuser" {
			t.Errorf("expected username 'testuser', got '%s'", users[0].Username)
		}
		if users[0].HashedPassword != "hashed_password" {
			t.Errorf("expected hashed password 'hashed_password', got '%s'", users[0].HashedPassword)
		}
	})

	t.Run("user already exists", func(t *testing.T) {
		storage := &mockUserStorage{
			users: []model.User{{Username: "testuser"}},
		}
		hasher := &mockPasswordHasher{}
		manager := NewUserManager(hasher, storage)

		err := manager.CreateUser("testuser", "password", "user")
		if err == nil {
			t.Fatal("expected an error, but got nil")
		}
		if err.Error() != "User already exists" {
			t.Errorf("expected error 'User already exists', got '%v'", err)
		}
	})

	t.Run("password hashing fails", func(t *testing.T) {
		storage := &mockUserStorage{}
		hasher := &mockPasswordHasher{
			hash: func(password string) (string, error) {
				return "", errors.New("hash error")
			},
		}
		manager := NewUserManager(hasher, storage)

		err := manager.CreateUser("testuser", "password", "user")
		if err == nil {
			t.Fatal("expected an error, but got nil")
		}
		if err.Error() != "hash error" {
			t.Errorf("expected error 'hash error', got '%v'", err)
		}
	})
}

func TestUserExists(t *testing.T) {
	storage := &mockUserStorage{
		users: []model.User{{Username: "existinguser"}},
	}
	manager := NewUserManager(&mockPasswordHasher{}, storage)

	if !manager.UserExists("existinguser") {
		t.Error("expected user to exist, but it doesn't")
	}
	if manager.UserExists("nonexistentuser") {
		t.Error("expected user to not exist, but it does")
	}
}

func TestGetUser(t *testing.T) {
	storage := &mockUserStorage{
		users: []model.User{{Username: "testuser", Role: "admin"}},
	}
	manager := NewUserManager(&mockPasswordHasher{}, storage)

	t.Run("get existing user", func(t *testing.T) {
		user, err := manager.GetUser("testuser")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if user.Username != "testuser" {
			t.Errorf("expected username 'testuser', got '%s'", user.Username)
		}
	})

	t.Run("get non-existent user", func(t *testing.T) {
		_, err := manager.GetUser("nonexistent")
		if !errors.Is(err, contracts.NotFoundUserErr) {
			t.Errorf("expected error '%v', got '%v'", contracts.NotFoundUserErr, err)
		}
	})
}

func TestListUsers(t *testing.T) {
	t.Run("list with users", func(t *testing.T) {
		storage := &mockUserStorage{
			users: []model.User{{Username: "user1"}, {Username: "user2"}},
		}
		manager := NewUserManager(&mockPasswordHasher{}, storage)
		users, err := manager.ListUsers()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(users) != 2 {
			t.Errorf("expected 2 users, got %d", len(users))
		}
	})

	t.Run("list with no users", func(t *testing.T) {
		storage := &mockUserStorage{}
		manager := NewUserManager(&mockPasswordHasher{}, storage)
		users, err := manager.ListUsers()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(users) != 0 {
			t.Errorf("expected 0 users, got %d", len(users))
		}
	})
}

func TestUpdateUser(t *testing.T) {
	storage := &mockUserStorage{
		users: []model.User{{Username: "testuser", Role: "user", HashedPassword: "hashed_password"}},
	}
	hasher := &mockPasswordHasher{}
	manager := NewUserManager(hasher, storage)

	t.Run("update password", func(t *testing.T) {
		err := manager.UpdateUser("testuser", "newpassword", "")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		user, _ := manager.GetUser("testuser")
		if user.HashedPassword != "hashed_newpassword" {
			t.Errorf("expected new hashed password, got '%s'", user.HashedPassword)
		}
	})

	t.Run("update role", func(t *testing.T) {
		err := manager.UpdateUser("testuser", "", "admin")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		user, _ := manager.GetUser("testuser")
		if user.Role != "admin" {
			t.Errorf("expected role 'admin', got '%s'", user.Role)
		}
	})

	t.Run("update non-existent user", func(t *testing.T) {
		err := manager.UpdateUser("nonexistent", "password", "user")
		if !errors.Is(err, contracts.NotFoundUserErr) {
			t.Errorf("expected error '%v', got '%v'", contracts.NotFoundUserErr, err)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	storage := &mockUserStorage{
		users: []model.User{{Username: "testuser"}},
	}
	manager := NewUserManager(&mockPasswordHasher{}, storage)

	t.Run("delete existing user", func(t *testing.T) {
		err := manager.DeleteUser("testuser")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if manager.UserExists("testuser") {
			t.Error("user should have been deleted")
		}
	})

	t.Run("delete non-existent user", func(t *testing.T) {
		err := manager.DeleteUser("nonexistent")
		if !errors.Is(err, contracts.NotFoundUserErr) {
			t.Errorf("expected error '%v', got '%v'", contracts.NotFoundUserErr, err)
		}
	})
}