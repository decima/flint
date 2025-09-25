package servers

import (
	"flint/service/contracts"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewServerStorage,
	fx.Annotate(
		NewManager,
		fx.As(new(contracts.ServerCollectionManager)),
	),
)
