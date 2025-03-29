package conf

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Database   Database   `yaml:"database"`
		Datasource Datasource `yaml:"datasource"`
		Cache      Cache      `yaml:"cache"`
		Api        Api        `yaml:"api"`
		Logger     Logger     `yaml:"logger"`
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
		if initErr != nil {
			return
		}
		initErr = initLogger(&configInstance.Logger)
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

	file, err := os.Open(configFile)
	if err != nil {
		return &Config{}, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	var config *Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return &Config{}, fmt.Errorf("error parsing YAML config: %v", err)
	}

	return config, nil

}
