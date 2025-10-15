package handlers

import (
	"flint/server/handlers/security"
	"flint/server/handlers/servers"
	"flint/server/handlers/servers/files"
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

	// Remote File handlers
	AsRoute(files.NewServerFileGetHandler),
	AsRoute(files.NewServerFilePutHandler),
	AsRoute(files.NewServerFileDeleteHandler),

	AsRoute(NewHomepage),
	AsRoute(users.NewGetAll),
	AsRoute(NewWebsocketSSHRoute),
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(utils.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
