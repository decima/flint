package security

import "flint/security"

type CredentialsPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type AuthResponsePayload struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type CurrentUserResponse struct {
	Username string        `json:"username"`
	Role     security.Role `json:"role"`
}
