package model

type Server struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Key      string `yaml:"key,omitempty" json:"key"`
	KeyPass  string `yaml:"keypass,omitempty" json:"keypass"`
	Password string `yaml:"password,omitempty" json:"password"`
	Username string `yaml:"username" json:"username"`
	WorkDir  string `yaml:"work_dir,omitempty" json:"work_dir"`
}
