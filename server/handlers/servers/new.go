package servers

import (
	"errors"
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/server/middlewares"
	"flint/service/contracts"
	"flint/utils/stringutils"

	"github.com/gin-gonic/gin"
)

type CreateServerHandler struct {
	serverManager contracts.ServerCollectionManager
}

func NewCreateServerHandler(serverManager contracts.ServerCollectionManager) *CreateServerHandler {
	return &CreateServerHandler{serverManager: serverManager}
}

func (csh CreateServerHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.POST, "/servers", security.UserOnly()
}

func (csh CreateServerHandler) Do(c *gin.Context) {
	var payload ServerCreatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, common.NewInvalidPayloadResponse(err))
		return
	}

	if payload.Name == "" {
		payload.Name = stringutils.RandomString(14)
	}

	if payload.Port == 0 {
		payload.Port = 22
	}

	newServer, err := csh.serverManager.CreateServer(payload.Name, payload.Host, payload.Port, payload.Username, payload.SSHKey)
	if errors.Is(err, contracts.DuplicateServerErr) {
		c.JSON(409, gin.H{"error": "ServerResponse with this name already exists"})
		return
	}
	if errors.Is(err, contracts.BadServerNameErr) {
		c.JSON(400, gin.H{"error": "invalid server name"})
		return
	}

	if err != nil {
		middlewares.GetLogger(c).Error().Err(err).Msg("Failed to create server")
		c.JSON(500, gin.H{"error": "Failed to create server"})
		return
	}

	c.JSON(201, common.NewResponse(NewServerResponse(payload.Name, newServer), "Server created"))
}
