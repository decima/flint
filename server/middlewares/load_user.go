package middlewares

import (
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
		c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
	}
	c.Set(CurrentUserKey, &user)
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
