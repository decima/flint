package servers

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/service/contracts"
	"sync"

	"github.com/gin-gonic/gin"
)

type SummaryHandler struct {
	serverCollectionManager contracts.ServerCollectionManager
	serverActions           contracts.ServerActions
}

func (s SummaryHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.GET, "/servers/:serverName/summary", security.UserOnly()
}

func (s SummaryHandler) Do(c *gin.Context) {
	serverName := c.Param("serverName")
	server, err := s.serverCollectionManager.GetServer(serverName)
	if err != nil {
		common.NotFound(c, "Server not found", err.Error())
		return
	}
	mux := sync.Mutex{}
	wg := sync.WaitGroup{}

	summary := map[string]any{
		"server": server,
		"docker": nil,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		dockerInfo, err := s.serverActions.DockerInfo(server)
		if err != nil {
			return
		}

		mux.Lock()
		summary["docker"] = dockerInfo

		mux.Unlock()
	}()

	wg.Wait()

	common.Ok(c, summary)
}

func NewSummaryHandler(serverCollectionManager contracts.ServerCollectionManager, serverActions contracts.ServerActions) *SummaryHandler {
	return &SummaryHandler{serverCollectionManager: serverCollectionManager, serverActions: serverActions}
}
