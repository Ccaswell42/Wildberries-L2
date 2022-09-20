package pkg

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port string `env:"PORT"`
}

func GetConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	return &Config{
		Port: os.Getenv("PORT"),
	}, nil
}
