package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	Url string
}

type CorsConfig struct {
	AllowOrigins string
}

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Cors     CorsConfig
}

func LoadConfig() (*Config, error) {
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	appConfig := AppConfig{
		Port: os.Getenv("APP_PORT"),
		Env:  os.Getenv("APP_ENV"),
	}

	databaseConfig := DatabaseConfig{
		Url: os.Getenv("DB_URL"),
	}

	corsConfig := CorsConfig{
		AllowOrigins: os.Getenv("CORS_ORIGINS"),
	}

	return &Config{
		App:      appConfig,
		Database: databaseConfig,
		Cors:     corsConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
