package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Middleware interface {
	Do(c *gin.Context)
}

var Module = fx.Provide(
	AsMiddleware(NewContextualLoggerMiddleware),
	AsMiddleware(NewAccessLogger),
	NewLoadUserMiddleware,
	NewPolicyMiddleware,
	NewAuthMiddleware,
)

func AsMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Middleware)),
		fx.ResultTags(`group:"middlewares"`),
	)
}
