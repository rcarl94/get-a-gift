package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Person struct {
	Name        string   `yaml:"name"`
	Descriptors []string `yaml:"descriptors"`
}

type GiftGrabConfig struct {
	People []Person `yaml:"people"`
}

func LoadConfig() (GiftGrabConfig, error) {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return GiftGrabConfig{}, fmt.Errorf("Error reading config file: %s", err)
	}

	var config GiftGrabConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return GiftGrabConfig{}, fmt.Errorf("Error parsing config file: %s", err)
	}

	return config, nil
}
