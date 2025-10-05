package remote

import (
	"flint/service/model"
	"fmt"
	"io"
	"path"
	"strings"

	"github.com/melbahja/goph"
	"github.com/pkg/sftp"
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

func (r *SSHClient) getClient(server model.Server) (*goph.Client, error) {
	auth, err := r.getAuthMethod(server)
	if err != nil {
		return nil, err
	}

	client, err := goph.NewUnknown(server.Username, server.Host, auth)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *SSHClient) Execute(server model.Server, command string, f func(out io.Reader) error) error {
	workDir := server.WorkDir

	client, err := r.getClient(server)
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

func (r *SSHClient) WriteFile(server model.Server, fileName string, content []byte) error {
	client, err := r.getClient(server)
	if err != nil {
		return err
	}
	defer client.Close()

	workDir := server.WorkDir
	if server.WorkDir == "" {
		workDir = "."
	}

	sftp, err := client.NewSftp()
	if err != nil {
		return err
	}

	fullPath := path.Join(workDir, fileName)

	sftp.MkdirAll(path.Dir(fullPath))

	file, err := sftp.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(content)

	return nil
}

func (r *SSHClient) GetFileContent(server model.Server, fileName string) ([]byte, error) {
	client, err := r.getClient(server)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	workDir := server.WorkDir
	if server.WorkDir == "" {
		workDir = "."
	}
	sftp, err := client.NewSftp()
	if err != nil {
		return nil, err
	}
	file, err := sftp.Open(path.Join(workDir, fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (r *SSHClient) DeleteFile(server model.Server, fileName string) error {
	client, err := r.getClient(server)
	if err != nil {
		return err
	}
	defer client.Close()

	workDir := server.WorkDir
	if server.WorkDir == "" {
		workDir = "."
	}
	sftp, err := client.NewSftp()
	if err != nil {
		return err
	}
	fullPath := path.Join(workDir, fileName)

	if path.Clean(fullPath) == path.Clean(workDir) {
		return fmt.Errorf("cannot delete the working directory")
	}

	err = r.deleteFileRecursively(sftp, path.Join(workDir, fileName))
	if err != nil {
		return err
	}
	return nil
}

func (r *SSHClient) deleteFileRecursively(sftp *sftp.Client, fullPath string) error {
	fileInfo, err := sftp.Lstat(fullPath)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		files, err := sftp.ReadDir(fullPath)
		if err != nil {
			return err
		}
		for _, file := range files {
			err = r.deleteFileRecursively(sftp, path.Join(fullPath, file.Name()))
			if err != nil {
				return err
			}
		}
	}
	return sftp.Remove(fullPath)
}

func (r *SSHClient) ListFiles(server model.Server, dir string) ([]model.File, error) {
	client, err := r.getClient(server)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	workDir := server.WorkDir
	if server.WorkDir == "" {
		workDir = "."
	}
	sftp, err := client.NewSftp()
	if err != nil {
		return nil, err
	}
	filesInfo, err := sftp.ReadDir(path.Join(workDir, dir))
	if err != nil {
		return nil, err
	}

	files := make([]model.File, 0, len(filesInfo))
	for _, fileInfo := range filesInfo {
		files = append(files, model.File{
			Name:  fileInfo.Name(),
			IsDir: fileInfo.IsDir(),
			Size:  fileInfo.Size(),
			Mode:  fileInfo.Mode().String(),
			MTime: fileInfo.ModTime(),
		})
	}
	return files, nil

}

func (r *SSHClient) GetFileInfo(server model.Server, fileName string) (model.File, error) {
	client, err := r.getClient(server)
	if err != nil {
		return model.File{}, err
	}
	defer client.Close()

	workDir := server.WorkDir
	if server.WorkDir == "" {
		workDir = "."
	}
	sftp, err := client.NewSftp()
	if err != nil {
		return model.File{}, err
	}
	fileInfo, err := sftp.Stat(path.Join(workDir, fileName))
	if err != nil {
		return model.File{}, err
	}

	return model.File{
		Name:  fileInfo.Name(),
		IsDir: fileInfo.IsDir(),
		Size:  fileInfo.Size(),
		Mode:  fileInfo.Mode().String(),
		MTime: fileInfo.ModTime(),
	}, nil
}
