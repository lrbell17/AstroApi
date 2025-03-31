package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	LogLevel string `yaml:"log_level"`
	LogFile  string `yaml:"log_file"`
}

func initLogger(loggerConfig *Logger) error {

	level, err := logrus.ParseLevel(loggerConfig.LogLevel)
	if err != nil {
		return fmt.Errorf("invalid log level %v", loggerConfig.LogLevel)
	}

	logDir := filepath.Dir(loggerConfig.LogFile)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory %v: %v", logDir, err)
	}
	logFile, err := os.OpenFile(loggerConfig.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file %v: %v", loggerConfig.LogFile, err)
	}

	logrus.SetOutput(logFile)
	logrus.SetLevel(level)

	return nil
}
