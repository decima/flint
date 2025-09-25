package model

import "flint/security"

type User struct {
	Username       string        `yaml:"username" json:"username"`
	Role           security.Role `yaml:"role" json:"role"`
	HashedPassword string        `yaml:"password" json:"-"`
}
