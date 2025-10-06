package server

import (
	"context"
	"flint/server/handlers/utils"
	"flint/server/middlewares"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/gin-contrib/static"
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

	// Serve static files from the frontend
	r.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))

	// API routes
	api := r.Group("/api")
	api.Use(gep.AuthMiddleware.Handler())
	api.Use(gep.LoadUser.Do)

	for _, route := range gep.Routes {
		method, path, securityPolicy := route.Route()
		api.Handle(
			string(method),
			string(path),
			gep.PolicyMiddleware.Handler(securityPolicy),
			route.Do,
		)
	}

	r.NoRoute(func(c *gin.Context) {
		if _, err := os.Stat("./frontend/dist/index.html"); err == nil {
			c.File("./frontend/dist/index.html")
			return
		}

		reverseProxy := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.URL.Scheme = "http"
				req.URL.Path = c.Request.URL.Path
				req.URL.Host = "localhost:5173"
				req.Host = "localhost:5173"
			},
		}
		reverseProxy.ServeHTTP(c.Writer, c.Request)
	})

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
