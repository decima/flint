package security

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/server/middlewares"

	"github.com/gin-gonic/gin"
)

type WhoAmIHandler struct {
}

func (w WhoAmIHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.GET, "/whoami", security.UserOnly()
}

func NewWhoAmIHandler() *WhoAmIHandler {
	return &WhoAmIHandler{}
}

func (w WhoAmIHandler) Do(c *gin.Context) {
	user, exists := middlewares.GetCurrentUser(c)
	if !exists {
		common.InternalServerError(c, "user not found in context")
		return
	}

	common.Ok(c, CurrentUserResponse{
		Username: user.Username,
		Role:     user.Role,
	})
}
