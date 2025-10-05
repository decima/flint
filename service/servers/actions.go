package servers

import (
	"flint/service/contracts"
	"flint/service/model"
	"io"
)

// todo : implement actions for servers:
// - get docker version
// - get running containers (docker)
// - get container logs (docker)
// - start/stop/restart containers (docker)
// - download docker images
// - send files to server
// - get system info (cpu, ram, disk, os, uptime, load avg)

type ActionMaker struct {
	remote contracts.RemoteAction
}

func NewActionMaker(remote contracts.RemoteAction) *ActionMaker {
	return &ActionMaker{remote: remote}
}

func (l *ActionMaker) DockerVersion(server model.Server) (string, error) {
	var output string
	err := l.remote.Execute(server, "docker version --format '{{.Server.Version}}'", func(out io.Reader) error {
		buf := make([]byte, 1024)
		n, err := out.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		output += string(buf[:n])
		return nil
	})
	if err != nil {
		return "", err
	}
	return output, nil
}
