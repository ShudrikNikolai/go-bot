package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Logs     LogConfig
	DB       SQLiteConfig
	Telegram TelegramConfig
}

type LogConfig struct {
	Style string
	Level string
}

type SQLiteConfig struct {
	hz string
}

type TelegramConfig struct {
	Token   string
	BotName string
	AuthId  string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Logs: LogConfig{
			Style: os.Getenv("LOG_STYLE"),
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: SQLiteConfig{
			hz: os.Getenv("LOG_LEVEL"),
		},
		Telegram: TelegramConfig{
			Token:   os.Getenv("TELEGRAM_TOKEN"),
			BotName: os.Getenv("TELEGRAM_BOT_NAME"),
			AuthId:  os.Getenv("TELEGRAM_AUTH_ID"),
		},
	}
	return cfg, nil
}
