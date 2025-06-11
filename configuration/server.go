package configuration

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
