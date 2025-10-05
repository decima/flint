package handlers

import (
	"flint/server/handlers/security"
	"flint/server/handlers/servers"
	"flint/server/handlers/setup"
	"flint/server/handlers/users"
	"flint/server/handlers/utils"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	AsRoute(setup.NewSetupHandler),

	AsRoute(security.NewLoginHandler),
	AsRoute(security.NewRefreshTokenHandler),
	AsRoute(security.NewWhoAmIHandler),

	AsRoute(servers.NewGetAllHandler),
	AsRoute(servers.NewCreateServerHandler),
	AsRoute(servers.NewSummaryHandler),

	AsRoute(NewHomepage),
	AsRoute(users.NewGetAll),
	AsRoute(NewWebsocketRoute),
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(utils.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
