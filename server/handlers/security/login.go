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
	logger := middlewares.GetLogger(c)

	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Debug().Err(err).Msg("invalid login payload")
		common.BadRequest(c, "invalid login payload", err.Error())
		return
	}

	user, err := l.userManager.GetUser(payload.Username)
	if err != nil || !l.passwordHasher.Verify(user.HashedPassword, payload.Password) {
		middlewares.UnauthorizedResponse(c, logger, "invalid username or password")
		return
	}

	token, err := l.jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		middlewares.UnauthorizedResponse(c, logger, "Unable to generate token: "+err.Error())
		return
	}

	refresh, err := l.jwt.GenerateRefreshToken(user.Username)
	if err != nil {
		middlewares.UnauthorizedResponse(c, logger, "Unable to generate refresh token: "+err.Error())
		return
	}

	common.Ok(c, AuthResponsePayload{
		Token:        token,
		RefreshToken: refresh,
	})
}
