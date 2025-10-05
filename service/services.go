package service

import (
	"flint/service/remote"
	"flint/service/servers"
	"flint/service/users"

	"go.uber.org/fx"
)

var Module = fx.Module("service",
	remote.Module,
	users.Module,
	servers.Module,
)
