package conf

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Database   Database   `yaml:"database"`
		Datasource Datasource `yaml:"datasource"`
	}
)

var (
	configInstance *Config
	once           sync.Once
	initErr        error
)

// Initialize config as singleton
func InitConfig(configFile string) error {

	once.Do(func() {
		configInstance, initErr = loadConfig(configFile)
	})
	return initErr
}

// Returns singleton instance of config
func GetConfig() (*Config, error) {
	if configInstance == nil {
		return nil, initErr
	}
	return configInstance, nil
}

// Initialize configuration from file
func loadConfig(configFile string) (*Config, error) {

	log.Infof("Loading configuration from %v", configFile)

	file, err := os.Open(configFile)
	if err != nil {
		log.Errorf("Error opening config file: %v", err)
		return &Config{}, err
	}
	defer file.Close()

	var config *Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Errorf("Error parsing YAML config: %v", err)
		return &Config{}, err
	}

	return config, nil

}
