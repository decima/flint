package contracts

import (
	"errors"
	"flint/service/model"
	"io"
)

type RemoteAction interface {
	Execute(server model.Server, command string, f func(out io.Reader) error) error
}

type ServerActions interface {
	DockerVersion(server model.Server) (string, error)
}

var InvalidServerErr = errors.New("invalid server")
