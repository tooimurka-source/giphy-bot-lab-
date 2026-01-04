package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	GiphyKey string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		GiphyKey: os.Getenv("GIPHY_API_KEY"),
	}, nil
}
