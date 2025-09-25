package users

import (
	"flint/server/middlewares"
	"flint/service/contracts"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewUserStorage,
	fx.Annotate(
		NewUserManager,
		fx.As(new(contracts.UsersManagerInterface)),
		fx.As(new(middlewares.UserLoader)),
	),
)
