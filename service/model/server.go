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
