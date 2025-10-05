package files

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/service/contracts"
	"io"
	"path"

	"github.com/gin-gonic/gin"
)

func NewServerFilePutHandler(
	serverCollectionManager contracts.ServerCollectionManager,
	remoteActions contracts.RemoteAction,
) *ServerFilePutHandler {
	return &ServerFilePutHandler{
		serverCollectionManager: serverCollectionManager,
		remoteActions:           remoteActions,
	}
}

type ServerFilePutHandler struct {
	serverCollectionManager contracts.ServerCollectionManager
	remoteActions           contracts.RemoteAction
}

func (s ServerFilePutHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.PUT, "/servers/:serverName/files/*filepath", security.UserOnly()
}

func (s ServerFilePutHandler) Do(c *gin.Context) {
	serverName := c.Param("serverName")
	filePath := c.Param("filepath")

	server, err := s.serverCollectionManager.GetServer(serverName)
	if err != nil {
		common.NotFound(c, "Invalid server ID", err.Error())
		return
	}

	filePath = path.Clean(filePath)

	fileContent, err := io.ReadAll(c.Request.Body)
	if err != nil {
		common.BadRequest(c, "Failed to read file content", err.Error())
		return
	}

	err = s.remoteActions.WriteFile(server, filePath, fileContent)
	if err != nil {
		common.InternalServerError(c, "Failed to write file", err.Error())
		return
	}

	common.Ok(c, gin.H{"message": "File uploaded successfully"})
}
