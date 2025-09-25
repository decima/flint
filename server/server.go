package server

import (
	"flint/server/handlers"
	"flint/server/middlewares"
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Module("server",
	fx.Provide(
		fx.Annotate(
			NewGinEngine,
		),
		NewHTTPServer,
	),
	handlers.Module,
	middlewares.Module,
	fx.Invoke(func(*http.Server) {}),
)
