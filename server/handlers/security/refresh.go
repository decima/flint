package security

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/server/middlewares"
	"flint/service/contracts"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	jwt         *security.Jwt
	userManager contracts.UsersManagerInterface
}

func (r RefreshTokenHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.POST, "/login/refresh", security.AnonymousOnly()
}

func (r RefreshTokenHandler) Do(c *gin.Context) {
	refreshTokenPayload := RefreshTokenPayload{}
	logger := middlewares.GetLogger(c)

	if err := c.ShouldBindJSON(&refreshTokenPayload); err != nil {
		logger.Debug().Err(err).Msg("invalid refresh token payload")
		common.BadRequest(c, "invalid refresh token payload", err.Error())
		return
	}

	username, err := r.jwt.ValidateRefreshToken(refreshTokenPayload.RefreshToken)
	if err != nil {
		middlewares.UnauthorizedResponse(c, logger, "cannot validate refresh token: "+err.Error())
		return
	}

	user, err := r.userManager.GetUser(username)
	if err != nil {
		middlewares.UnauthorizedResponse(c, logger, "cannot find user: "+err.Error())
		return
	}

	token, err := r.jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		middlewares.UnauthorizedResponse(c, logger, "Unable to generate token: "+err.Error())
		return
	}

	refresh, err := r.jwt.GenerateRefreshToken(user.Username)
	if err != nil {
		middlewares.UnauthorizedResponse(c, logger, "Unable to generate refresh token: "+err.Error())
		return
	}

	common.Ok(c, AuthResponsePayload{
		Token:        token,
		RefreshToken: refresh,
	})
}

func NewRefreshTokenHandler(jwt *security.Jwt, userManager contracts.UsersManagerInterface) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		jwt:         jwt,
		userManager: userManager,
	}
}
