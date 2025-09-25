package main

import (
	"flint/config"
	"flint/security"
	"flint/server"
	"flint/service"
	"os"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		security.Module,
		server.Module,
		service.Module,
		fx.Provide(
			config.NewConfig,
			NewLogger,
		),
		//		fx.NopLogger,
	).Run()
}

func NewLogger() *zerolog.Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().
		Str("version", "0.1.0").
		Logger()
	return &logger
}
