package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasherInterface interface {
	Hash(password string) (string, error)
	Verify(hash string, password string) bool
}

type PasswordHasher struct {
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (p *PasswordHasher) Hash(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (p *PasswordHasher) Verify(hash string, password string) bool {
	fmt.Println("verifying password :", password, " against hash: ", hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
