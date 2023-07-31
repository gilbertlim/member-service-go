package config

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	RunMode      string `yaml:"runMode"`
	HttpPort     int    `yaml:"httpPort"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
}

type Database struct {
	Type string `yaml:"type"`
}
