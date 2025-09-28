package middlewares

import (
	"flint/server/common"
	"flint/service/contracts"
	"flint/service/model"

	"github.com/gin-gonic/gin"
)

const CurrentUserKey = "current_user"

type LoadUserMiddleware struct {
	userManager contracts.UsersManagerInterface
}

func (l *LoadUserMiddleware) Do(c *gin.Context) {
	c.Get(UserIDKey)
	userID := c.GetString(UserIDKey)
	if userID == "" {
		c.Next()
		return
	}

	user, err := l.userManager.GetUser(userID)
	if err != nil {
		common.NotFound(c, "user not found")
		return
	}
	c.Set(CurrentUserKey, &user)
	c.Next()
}

func GetCurrentUser(c *gin.Context) (*model.User, bool) {
	user, exists := c.Get(CurrentUserKey)
	if !exists {
		return nil, false
	}

	userCasted, ok := user.(*model.User)
	return userCasted, ok
}

func NewLoadUserMiddleware(userManager contracts.UsersManagerInterface) *LoadUserMiddleware {
	return &LoadUserMiddleware{userManager: userManager}
}
