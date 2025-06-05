package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"xswitcher/core"
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
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Can not locate User Home Directory: %s\n", err.Error())
		}

		data := `layouts:
	main: "us"
	secondary: "us"
		`

		filePath := home + "/" + core.CONFIG_DIR + "/" + core.FILENAME

		// Create file
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Failed to create file: %s\n", err.Error())
		}
		defer file.Close()

		// Write data
		_, err = file.WriteString(data)
		if err != nil {
			fmt.Printf("Failed to write data to file: %s\n", err.Error())
		}
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
