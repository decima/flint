package server

import (
	"context"
	"flint/server/handlers/utils"
	"flint/server/middlewares"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type GinEngineParams struct {
	fx.In

	PolicyMiddleware *middlewares.PolicyMiddleware
	AuthMiddleware   *middlewares.AuthMiddleware
	LoadUser         *middlewares.LoadUserMiddleware
	Routes           []utils.Route            `group:"routes"`
	Middlewares      []middlewares.Middleware `group:"middlewares"`
}

func NewGinEngine(gep GinEngineParams,
) *gin.Engine {
	r := gin.New()

	for _, m := range gep.Middlewares {
		r.Use(m.Do)
	}

	for _, route := range gep.Routes {
		method, path, securityPolicy := route.Route()
		group := r.Group("")
		group.
			Use(gep.AuthMiddleware.Handler()).
			Use(gep.PolicyMiddleware.Handler(securityPolicy)).
			Use(gep.LoadUser.Do).
			Handle(string(method), string(path), route.Do)
	}

	return r
}

func NewHTTPServer(lc fx.Lifecycle, engine *gin.Engine, logger *zerolog.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: engine}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			logger.Info().Msgf("Starting HTTP server at %v", srv.Addr)
			fmt.Println("http://0.0.0.0:8080")
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
