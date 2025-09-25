package service

import (
	"flint/service/servers"
	"flint/service/users"

	"go.uber.org/fx"
)

var Module = fx.Module("service",
	users.Module,
	servers.Module,
)
