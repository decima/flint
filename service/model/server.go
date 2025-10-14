package model

type Server struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Key      string `yaml:"key,omitempty" json:"-"`
	KeyPass  string `yaml:"keypass,omitempty" json:"-"`
	Password string `yaml:"password,omitempty" json:"-"`
	Username string `yaml:"username" json:"username"`
	WorkDir  string `yaml:"work_dir,omitempty" json:"work_dir"`
}

type DockerInfo struct {
	Containers struct {
		Total   int `json:"total"`
		Running int `json:"running"`
		Paused  int `json:"paused"`
		Stopped int `json:"stopped"`
	} `json:"containers"`
	Images int `json:"images"`
	Server struct {
		OperatingSystem string `json:"operating_system"`
		Architecture    string `json:"architecture"`
		ServerVersion   string `json:"server_version"`
		KernelVersion   string `json:"kernel_version"`
	} `json:"server"`
	Client struct {
		Version         string `json:"version"`
		APIVersion      string `json:"api_version"`
		Architecture    string `json:"architecture"`
		OperatingSystem string `json:"operating_system"`
	} `json:"client"`
}
