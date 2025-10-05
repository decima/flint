package remote

import (
	"flint/service/contracts"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewSSHClient,
		fx.As(new(contracts.RemoteAction)),
	),
)
