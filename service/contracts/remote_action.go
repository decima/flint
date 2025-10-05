package contracts

import (
	"errors"
	"flint/service/model"
	"io"
)

type RemoteAction interface {
	Execute(server model.Server, command string, f func(out io.Reader) error) error
	WriteFile(server model.Server, fileName string, content []byte) error
	GetFileContent(server model.Server, fileName string) ([]byte, error)
	DeleteFile(server model.Server, fileName string) error
	ListFiles(server model.Server, dir string) ([]model.File, error)
	GetFileInfo(server model.Server, fileName string) (model.File, error)
}

type ServerActions interface {
	DockerVersion(server model.Server) (string, error)
}

var InvalidServerErr = errors.New("invalid server")
