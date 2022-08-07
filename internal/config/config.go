package config

import (
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Database *DatabaseConfig
	Server   *ServerConfig
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

type ServerConfig struct {
	Network string
	Port    string
	Host    string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		&DatabaseConfig{
			Host:           getEnv("POSTGRES_HOST", ""),
			Port:           getEnv("POSTGRES_PORT", ""),
			Name:           getEnv("POSTGRES_DB", ""),
			User:           getEnv("POSTGRES_USER", ""),
			Password:       getEnv("POSTGRES_PASSWORD", ""),
			MaxConnections: getEnv("MAX_POSTGRES_CONNECTIONS", ""),
			SslMode:        getEnv("POSTGRES_SSL_MODE", ""),
			CtxLifeTime:    getEnvTimeDuration("POSTGRES_CTX_LIFETIME", 5*time.Second),
		},
		&ServerConfig{
			Network: getEnv("API_NETWORK", ""),
			Host:    getEnv("API_HOST", ""),
			Port:    getEnv("API_PORT", ""),
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
