package servers

import (
	"encoding/json"
	"flint/service/model"
	"fmt"
	"io"
	"strings"
)

// todo : implement actions for servers:
// - get docker version
// - get running containers (docker)
// - get container logs (docker)
// - start/stop/restart containers (docker)
// - download docker images
// - send files to server
// - get system info (cpu, ram, disk, os, uptime, load avg)

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
	return strings.Trim(output, "\n"), nil
}

func (l *ActionMaker) DockerComposeVersion(server model.Server) (string, error) {
	var output string
	err := l.remote.Execute(server, "docker compose version --short", func(out io.Reader) error {
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
	return strings.Trim(output, "\n"), nil
}

func (l *ActionMaker) DockerInfo(server model.Server) (model.DockerInfo, error) {
	format := `{
	"containers": {
		"total":{{.Containers}},
		"running":{{.ContainersRunning}},
		"paused":{{.ContainersPaused}},
		"stopped":{{.ContainersStopped}}
	},
	"images":{{.Images}},
	"server": {
		"operating_system":"{{.OperatingSystem}}",
		"architecture":"{{.Architecture}}",
		"server_version":"{{.ServerVersion}}",
		"kernel_version":"{{.KernelVersion}}"
	},
	"client":{
		"version":"{{.ClientInfo.Version}}",
		"api_version":"{{.ClientInfo.DefaultAPIVersion}}",
		"architecture":"{{.ClientInfo.Arch}}",
		"operating_system":"{{.ClientInfo.Os}}"
	}
}`

	cmd := fmt.Sprintf("docker info -f '%s'", format)

	var output = model.DockerInfo{}

	err := l.remote.Execute(server, cmd, func(out io.Reader) error {
		buf := make([]byte, 10*1024)
		n, err := out.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		err = json.Unmarshal(buf[:n], &output)
		return err
	})

	return output, err

}
