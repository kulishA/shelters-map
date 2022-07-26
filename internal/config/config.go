package config

import (
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Database *DatabaseConfig
}

type DatabaseConfig struct {
	Name           string
	User           string
	Password       string
	Host           string
	Port           string
	MaxConnections string
	SslMode        string
	CtxLifeTime    time.Duration
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		&DatabaseConfig{
			Host:           getEnv("DB_HOST", ""),
			Port:           getEnv("DB_PORT", ""),
			Name:           getEnv("DB_NAME", ""),
			User:           getEnv("DB_USER", ""),
			Password:       getEnv("DB_PASSWORD", ""),
			MaxConnections: getEnv("MAX_DB_CONNECTIONS", ""),
			SslMode:        getEnv("DB_SSL_MODE", ""),
			CtxLifeTime:    getEnvTimeDuration("DB_CTX_LIFETIME", 5*time.Second),
		},
	}, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvTimeDuration(key string, defaultVal time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		t, err := time.ParseDuration(value)
		if err != nil {
			return defaultVal
		}

		return t
	}

	return defaultVal
}
