package servers

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/service/contracts"

	"github.com/gin-gonic/gin"
)

type GetAllHandler struct {
	serverCollectionManager contracts.ServerCollectionManager
}

func NewGetAllHandler(serverCollectionManager contracts.ServerCollectionManager) *GetAllHandler {
	return &GetAllHandler{serverCollectionManager: serverCollectionManager}
}

func (g GetAllHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.GET, "/servers", security.UserOnly()
}

func (g GetAllHandler) Do(c *gin.Context) {
	servers, err := g.serverCollectionManager.ListServers()
	if err != nil {
		common.InternalServerError(c, "Failed to retrieve servers", err.Error())
		return
	}

	common.Ok(c, NewServerListResponse(servers))
}
