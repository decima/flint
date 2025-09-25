package security

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/server/middlewares"
	"flint/service/contracts"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type LoginHandler struct {
	userManager    contracts.UsersManagerInterface
	passwordHasher security.PasswordHasherInterface
	jwt            *security.Jwt
	logger         *zerolog.Logger
}

func NewLoginHandler(userManager contracts.UsersManagerInterface, logger *zerolog.Logger, hasher security.PasswordHasherInterface, jwt *security.Jwt) *LoginHandler {
	return &LoginHandler{
		logger:         logger,
		userManager:    userManager,
		passwordHasher: hasher,
		jwt:            jwt,
	}
}

func (l LoginHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return "POST", "/login", security.AnonymousOnly()
}

func (l LoginHandler) Do(c *gin.Context) {
	payload := CredentialsPayload{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		middlewares.UnauthorizedResponse(c, middlewares.GetLogger(c), "invalid login payload: "+err.Error())
		return
	}

	user, err := l.userManager.GetUser(payload.Username)
	if err != nil || !l.passwordHasher.Verify(user.HashedPassword, payload.Password) {
		middlewares.UnauthorizedResponse(c, middlewares.GetLogger(c), "invalid username or password")
		return
	}

	token, err := l.jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		middlewares.UnauthorizedResponse(c, middlewares.GetLogger(c), "Unable to generate token: "+err.Error())
		return
	}

	refresh, err := l.jwt.GenerateRefreshToken(user.Username)
	if err != nil {
		middlewares.UnauthorizedResponse(c, middlewares.GetLogger(c), "Unable to generate refresh token: "+err.Error())
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	c.JSON(200, common.NewResponse(
		AuthResponsePayload{
			Token:        token,
			RefreshToken: refresh,
		},
	))
	return

}
