package contracts

import (
	"errors"
	"flint/service/model"
)

type ServerCollectionManager interface {
	CreateServer(
		name string,
		hostOrIp string,
		port int,
		user string,
		workDir string,
		sshKey string,
		keyPass string,
		password string,
	) (model.Server, error)
	DeleteServer(name string) error
	GetServer(name string) (model.Server, error)
	ListServers() (ServerCollection, error)
}

type ServerValidator interface {
	Validate(server model.Server) error
}

type ServerCollection map[string]model.Server

var DuplicateServerErr = errors.New("server with this name already exists")
var ServerNotFoundErr = errors.New("server with this name does not exist")
var BadServerNameErr = errors.New("bad server name")
