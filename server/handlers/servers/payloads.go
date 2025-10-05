package servers

import (
	"flint/service/contracts"
	"flint/service/model"
)

type ServerListResponse struct {
	Servers []ServerResponse `json:"servers"`
}

func NewServerListResponse(servers contracts.ServerCollection) ServerListResponse {
	payload := ServerListResponse{
		Servers: make([]ServerResponse, 0, len(servers)),
	}

	for id, server := range servers {
		payload.Servers = append(payload.Servers, NewServerResponse(id, server))
	}

	return payload
}

type ServerCreatePayload struct {
	Name       string `json:"name,omitempty"`
	Port       int    `json:"port" default:"22"`
	Host       string `json:"host" binding:"required"`
	Username   string `json:"username" binding:"required"`
	SSHKey     string `json:"ssh_key,omitempty"`
	SSHKeyPass string `json:"ssh_key_pass,omitempty"`
	Password   string `json:"password,omitempty"`
	Workdir    string `json:"workdir,omitempty"`
}

type ServerResponse struct {
	ID       string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Workdir  string `json:"workdir,omitempty"`
}

func NewServerResponse(id string, server model.Server) ServerResponse {
	return ServerResponse{
		ID:       id,
		Host:     server.Host,
		Port:     server.Port,
		Username: server.Username,
		Workdir:  server.WorkDir,
	}
}
