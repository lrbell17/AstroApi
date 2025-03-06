package conf

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const fileName = "config.yaml"

type (
	Config struct {
		Database struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
			User string `yaml:"user"`
			Pass string `yaml:"password"`
		} `yaml:"database"`
	}
)

// Initialize configuration from file
func InitConfig() (Config, error) {

	logrus.Infof("Initializing configuration from %v", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		logrus.Errorf("Error opening config file: %v", err)
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		logrus.Errorf("Error parsing YAML config: %v", err)
		return Config{}, err
	}

	return config, nil

}
