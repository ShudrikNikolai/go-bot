package config

import (
	"os"
	"strconv"

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
	AuthId  int64
}

func LoadConfig() (*Config, error) {
	// todo
	authIdStr := os.Getenv("TELEGRAM_AUTH_ID")
	token := os.Getenv("TELEGRAM_TOKEN")

	authId, _ := strconv.ParseInt(authIdStr, 10, 64)

	cfg := &Config{
		Logs: LogConfig{
			Style: os.Getenv("LOG_STYLE"),
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: SQLiteConfig{
			hz: os.Getenv("LOG_LEVEL"),
		},
		Telegram: TelegramConfig{
			Token:   token,
			BotName: os.Getenv("TELEGRAM_BOT_NAME"),
			AuthId:  authId,
		},
	}
	return cfg, nil
}
