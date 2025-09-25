package handlers

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/server/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Homepage struct{}

func (h *Homepage) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.GET, "/", security.AnonymousOrUser()
}

func NewHomepage() *Homepage {
	return &Homepage{}

}

func (h *Homepage) Do(c *gin.Context) {
	log := middlewares.GetLogger(c)

	log.Debug().Msgf("Homepage")

	c.JSON(http.StatusOK, common.NewResponse(nil, "Welcome to Flint!", "Land of stones."))

}
