package files

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/service/contracts"
	"path"

	"github.com/gin-gonic/gin"
)

func NewServerFileDeleteHandler(
	serverCollectionManager contracts.ServerCollectionManager,
	remoteActions contracts.RemoteAction,
) *ServerFileDeleteHandler {
	return &ServerFileDeleteHandler{
		serverCollectionManager: serverCollectionManager,
		remoteActions:           remoteActions,
	}
}

type ServerFileDeleteHandler struct {
	serverCollectionManager contracts.ServerCollectionManager
	remoteActions           contracts.RemoteAction
}

func (s ServerFileDeleteHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.DELETE, "/servers/:serverName/files/*filepath", security.UserOnly()
}

func (s ServerFileDeleteHandler) Do(c *gin.Context) {
	serverName := c.Param("serverName")
	filePath := c.Param("filepath")

	server, err := s.serverCollectionManager.GetServer(serverName)
	if err != nil {
		common.NotFound(c, "Invalid server ID", err.Error())
		return
	}

	filePath = path.Clean(filePath)

	err = s.remoteActions.DeleteFile(server, filePath)
	if err != nil {
		common.InternalServerError(c, "Failed to delete file", err.Error())
		return
	}

	common.Ok(c, gin.H{"message": "File deleted successfully"})
}
