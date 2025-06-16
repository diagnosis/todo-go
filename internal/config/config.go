package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port     string
	Database struct {
		URL string
	}
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	cfg := &Config{
		Port: os.Getenv("PORT"),
	}
	cfg.Database.URL = os.Getenv("DATABASE_URL")
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	if cfg.Database.URL == "" {
		return nil, os.ErrInvalid
	}
	return cfg, nil
}
