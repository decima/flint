package middlewares

import (
	"flint/security"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const UserIDKey = "user_id"
const UserRoleKey = "user_role"

type PolicyMiddleware struct {
	logger *zerolog.Logger
}

func NewPolicyMiddleware(logger *zerolog.Logger) *PolicyMiddleware {
	return &PolicyMiddleware{
		logger: logger,
	}
}

func (p *PolicyMiddleware) Handler(policy *security.Policy) gin.HandlerFunc {
	return func(c *gin.Context) {
		if policy == nil {
			c.Next()
			return
		}

		passport := NewPassportFromContext(c)

		if !policy.Validate(passport) {
			UnauthorizedResponse(c, p.logger, "Policy validation failed")
			return
		}

		c.Next()
	}
}

func NewPassportFromContext(c *gin.Context) *security.Passport {
	userID := c.GetString(UserIDKey)
	role := c.GetString(UserRoleKey)
	if role == "" {
		role = string(security.Anon)
	}
	roleSlice := []security.Role{security.Role(role)}

	return &security.Passport{
		UserID: userID,
		Roles:  roleSlice,
		IP:     c.ClientIP(),
	}
}

type UserLoader interface {
	UserExists(username string) bool
}

type AuthMiddleware struct {
	logger *zerolog.Logger
	Loader UserLoader
	jwt    *security.Jwt
}

func NewAuthMiddleware(logger *zerolog.Logger, jwt *security.Jwt, Loader UserLoader) *AuthMiddleware {
	return &AuthMiddleware{
		logger: logger,
		jwt:    jwt,
		Loader: Loader,
	}
}

func (a *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Next()
			return
		}

		//remove Bearer prefix if exists
		if len(tokenString) > 7 && strings.ToLower(tokenString[:7]) == "bearer " {
			tokenString = tokenString[7:]
		}

		userId, role, err := a.jwt.ValidateToken(tokenString)
		if err != nil {
			UnauthorizedResponse(c, a.logger, "error parsing token")
			return
		}

		if !a.Loader.UserExists(userId) {
			UnauthorizedResponse(c, a.logger, "User from token not found")
			return
		}

		c.Set(UserIDKey, userId)
		c.Set(UserRoleKey, string(role))
		c.Next()
	}
}

func UnauthorizedResponse(c *gin.Context, logger *zerolog.Logger, message string) {
	logger.Debug().Msg(message)
	c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
}
