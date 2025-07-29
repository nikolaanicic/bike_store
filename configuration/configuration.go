package configuration

import (
	"bike_store/log"
	"os"

	"gopkg.in/yaml.v3"
)

var configPath string = os.Getenv("CONFIG_PATH")

var CentralServerHost string = ""

type Config struct {
	Database          Database `yaml:"database"`
	Server            Server   `yaml:"server"`
	CentralServerHost string   `yaml:"central_server_host"`
}

func readFile(path string) (string, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func Get() (Config, error) {
	var content []byte
	var err error
	var cfg Config

	log.Info("reading configuration")
	if content, err = os.ReadFile(configPath); err != nil {
		return cfg, err
	}

	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return cfg, err
	}

	if pass, err := readFile(cfg.Database.Password); err != nil {
		log.Fatalf("unable to read db password field file...shuting down: %v", err)
		return Config{}, err
	} else {
		cfg.Database.Password = pass
	}

	CentralServerHost = cfg.CentralServerHost

	return cfg, nil
}
