package servers

import (
	"errors"
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/server/middlewares"
	"flint/service/contracts"
	"flint/utils/stringutils"
	"net/http"

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
	logger := middlewares.GetLogger(c)

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
	if payload.Workdir == "" {
		payload.Workdir = "."
	}

	newServer, err := csh.serverManager.CreateServer(payload.Name, payload.Host, payload.Port, payload.Username, payload.Workdir, payload.SSHKey, payload.SSHKeyPass, payload.Password)
	if errors.Is(err, contracts.DuplicateServerErr) {
		common.Err(c, http.StatusConflict, "server with this name already exists")
		return
	}
	if errors.Is(err, contracts.BadServerNameErr) {
		common.BadRequest(c, "invalid server name")
		return
	}

	if err != nil {
		logger.Error().Err(err).Msg("Failed to create server")
		common.InternalServerError(c, "failed to create server", err.Error())
		return
	}

	common.Ok(c, NewServerResponse(payload.Name, newServer), "Server created")
}
