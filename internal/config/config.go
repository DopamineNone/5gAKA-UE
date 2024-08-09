package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var (
	defaultFile = ".ue.yaml"
	config      *Config
)

type Config struct {
	SUPI    string
	Mcc     string
	Mnc     string
	Key     string
	Opc     string
	AmfAddr string
	AmfPort int
}

func ReadYamlConfig(path string) error {
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = home + defaultFile
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(config)

	return err
}

func GetConfig() *Config {
	return config
}

func init() {
	config = &Config{}
}
