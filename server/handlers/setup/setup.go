package setup

import (
	"flint/security"
	"flint/server/common"
	"flint/server/handlers/utils"
	"flint/service/contracts"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type SetupHandler struct {
	userManager contracts.UsersManagerInterface
	logger      *zerolog.Logger
}

func (s SetupHandler) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.POST, "/setup", security.
		AnonymousOnly().
		WithCustomValidator(func() bool {
			users, err := s.userManager.ListUsers()

			if err != nil {
				s.logger.Error().Err(err).Msg("Error listing users, disabling setup endpoint")
				return false
			}
			// this is for the future when we have more roles
			// for now, if there is any user, we disable setup
			//for _, user := range users {
			//	if user.Role == security.User{
			//		return false
			//	}
			//}
			//
			return len(users) == 0
		})
}

func (s SetupHandler) Do(c *gin.Context) {
	payload := SetupPayload{}
	if err := c.ShouldBindJSON(&payload); err != nil {

		c.JSON(400, common.NewInvalidPayloadResponse(err))
		return
	}

	if err := s.userManager.CreateUser(payload.Username, payload.Password, security.User); err != nil {
		c.JSON(500, common.NewErrorResponse(err, "Error creating user"))
		return
	}

	c.JSON(200, common.NewResponse(Response{
		Username: payload.Username,
	}, "Setup completed", "You can now log in with your user"))

}

func NewSetupHandler(userManager contracts.UsersManagerInterface, logger *zerolog.Logger) *SetupHandler {
	return &SetupHandler{userManager: userManager, logger: logger}
}
