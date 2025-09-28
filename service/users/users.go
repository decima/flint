package users

import (
	"flint/server/middlewares"
	"flint/service/contracts"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewUserStorage,
		fx.As(new(UserStorageInterface)),
	),
	fx.Annotate(
		NewUserManager,
		fx.As(new(contracts.UsersManagerInterface)),
		fx.As(new(middlewares.UserLoader)),
	),
)
