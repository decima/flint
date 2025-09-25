package security

import "testing"

func TestHash(t *testing.T) {
	hasher := NewPasswordHasher()
	password := "my_secure_password"

	hash, err := hasher.Hash(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if !hasher.Verify(password, hash) {
		t.Errorf("Password verification failed")
	}

	if hasher.Verify("wrong_password", hash) {
		t.Errorf("Password verification should have failed for wrong password")
	}
}
