package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Config - application config struct
type Config struct {
	Port    string `dotenv:"PORT"`
	GoEnv   string `dotenv:"GO_ENV"`
	Version string `dotenv:"VERSION"`
}

// NewConfig - creates the application config struct
func NewConfig() *Config {
	return &Config{
		Port:    GetEnv("PORT", "9000"),
		GoEnv:   GetEnv("GO_ENV", "local"),
		Version: GetEnv("Version", "local"),
	}
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		if len(defaultValue) > 0 {
			return defaultValue
		}
		logrus.Warn(fmt.Sprintf("Missing environment variable %s not set", key))
	}

	return value
}
