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
	if err := c.ShouldBindJSON(&refreshTokenPayload); err != nil {
		middlewares.UnauthorizedResponse(c, middlewares.GetLogger(c), "invalid login payload: "+err.Error())
		return
	}

	username, err := r.jwt.ValidateRefreshToken(refreshTokenPayload.RefreshToken)
	if err != nil {
		middlewares.UnauthorizedResponse(c, middlewares.GetLogger(c), "cannot validate refresh token: "+err.Error())
		return
	}

	user, err := r.userManager.GetUser(username)
	if err != nil {
		middlewares.UnauthorizedResponse(c, middlewares.GetLogger(c), "cannot find user: "+err.Error())
		return
	}

	token, err := r.jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		middlewares.UnauthorizedResponse(c, middlewares.GetLogger(c), "Unable to generate token: "+err.Error())
		return
	}

	refresh, err := r.jwt.GenerateRefreshToken(user.Username)
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

func NewRefreshTokenHandler(jwt *security.Jwt, userManager contracts.UsersManagerInterface) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		jwt:         jwt,
		userManager: userManager,
	}
}
