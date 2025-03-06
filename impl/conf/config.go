package conf

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

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
func InitConfig(configFile string) (Config, error) {

	log.Infof("Initializing configuration from %v", configFile)

	file, err := os.Open(configFile)
	if err != nil {
		log.Errorf("Error opening config file: %v", err)
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Errorf("Error parsing YAML config: %v", err)
		return Config{}, err
	}

	return config, nil

}
