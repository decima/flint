package servers

import (
	"flint/service/contracts"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewServerStorage,
	fx.Annotate(
		NewManager,
		fx.As(new(contracts.ServerCollectionManager), new(contracts.ServerValidator)),
	),
	fx.Annotate(
		NewActionMaker,
		fx.As(new(contracts.ServerActions)),
	),
)
