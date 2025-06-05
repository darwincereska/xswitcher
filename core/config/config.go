package config

import (
	"os"
	"path/filepath"
	"xswitcher/core"

	"gopkg.in/yaml.v3"
)

type Layout struct {
	Main      string `yaml:"main"`
	Secondary string `yaml:"secondary"`
}

type Config struct {
	Layouts Layout `yaml:"layouts"`
}

func Init(path string) {
	if path != "" {

	} else {

	}
}

func ParseConfig(file string) (Config, error) {
	var cfg Config
	var newFile string
	if file != "" {
		newFile = file
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return cfg, err
		}
		newFile = filepath.Join(home, core.CONFIG_DIR, core.FILENAME)
	}

	// Reads file
	data, err := os.ReadFile(newFile)
	if err != nil {
		return cfg, err
	}

	// Unmarshal YAML into struct
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
