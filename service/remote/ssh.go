package remote

import (
	"flint/service/model"
	"fmt"
	"io"
	"strings"

	"github.com/melbahja/goph"
)

type SSHClient struct{}

func NewSSHClient() *SSHClient {
	return &SSHClient{}
}

func (r *SSHClient) getAuthMethod(server model.Server) (goph.Auth, error) {
	if server.Key != "" {
		return goph.RawKey(server.Key, server.KeyPass)
	}
	return goph.Password(server.Password), nil
}

func (r *SSHClient) Execute(server model.Server, command string, f func(out io.Reader) error) error {
	workDir := server.WorkDir
	if server.WorkDir == "" {
		workDir = "."
	}

	auth, err := r.getAuthMethod(server)
	if err != nil {
		return err
	}

	client, err := goph.NewUnknown(server.Username, server.Host, auth)
	if err != nil {
		return err
	}
	defer client.Close()
	fullCommand := fmt.Sprintf("sh -c 'mkdir -p %s && cd %s && { %s; }'", workDir, workDir, strings.ReplaceAll(command, "'", "'\\''"))

	cmd, err := client.Command(fullCommand)
	fmt.Println(cmd)
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	go f(io.MultiReader(stdout, stderr))

	return cmd.Run()
}
