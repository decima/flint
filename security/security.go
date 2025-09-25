package security

import (
	"go.uber.org/fx"
)

var Module = fx.Module("security",
	fx.Provide(
		fx.Annotate(
			NewPasswordHasher,
			fx.As(new(PasswordHasherInterface)),
		),
		NewPasswordHasher,
		NewJwt,
	),
)
