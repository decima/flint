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
		c.JSON(500, gin.H{"error": "Failed to retrieve servers"})
		return
	}

	c.JSON(200, common.NewResponse(NewServerListResponse(servers)))
}
