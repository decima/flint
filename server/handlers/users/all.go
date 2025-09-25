package users

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/service/contracts"

	"github.com/gin-gonic/gin"
)

type GetAllHandler struct {
	userManager contracts.UsersManagerInterface
}

func NewGetAll(userManager contracts.UsersManagerInterface) *GetAllHandler {
	return &GetAllHandler{userManager: userManager}
}

func (p *GetAllHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.GET, "/users", security.UserOnly()
}

func (p *GetAllHandler) Do(c *gin.Context) {
	users, err := p.userManager.ListUsers()
	if err != nil {
		c.JSON(500, common.NewErrorResponse(err, "Error listing users"))
		return
	}

	c.JSON(200, common.NewResponse(users))
}
