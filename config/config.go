package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramToken string
	DatabaseURL   string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		TelegramToken: os.Getenv("API_TOKEN"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
	}, nil
}
