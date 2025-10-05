package files

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/service/contracts"
	"path"

	"github.com/gin-gonic/gin"
)

func NewServerFileGetHandler(
	serverCollectionManager contracts.ServerCollectionManager,
	remoteActions contracts.RemoteAction,
) *ServerFileGetHandler {
	return &ServerFileGetHandler{
		serverCollectionManager: serverCollectionManager,
		remoteActions:           remoteActions,
	}
}

type ServerFileGetHandler struct {
	serverCollectionManager contracts.ServerCollectionManager
	remoteActions           contracts.RemoteAction
}

func (s ServerFileGetHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.GET, "/servers/:serverName/files/*filepath", security.UserOnly()
}

func (s ServerFileGetHandler) Do(c *gin.Context) {
	serverName := c.Param("serverName")
	filePath := c.Param("filepath")

	server, err := s.serverCollectionManager.GetServer(serverName)
	if err != nil {
		common.NotFound(c, "Invalid server ID", err.Error())
		return
	}
	filePath = path.Clean(filePath)

	file, err := s.remoteActions.GetFileInfo(server, filePath)
	if err != nil {
		common.InternalServerError(c, "Failed to list file", err.Error())
		return
	}

	if file.IsDir {
		fileList, err := s.remoteActions.ListFiles(server, filePath)
		if err != nil {
			common.InternalServerError(c, "Failed to list files", err.Error())
			return
		}

		common.Ok(c, fileList)
		return
	}

	// If it's a file, return content
	content, err := s.remoteActions.GetFileContent(server, filePath)
	if err != nil {
		common.InternalServerError(c, "Failed to get file content", err.Error())
		return
	}

	c.Data(200, "application/octet-stream", content)

}
