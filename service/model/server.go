package model

type Server struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Key      string `yaml:"key" json:"key"`
	Username string `yaml:"username" json:"username"`
}
