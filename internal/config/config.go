package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Telegram TelegramConfig
}

type TelegramConfig struct {
	Token   string
	BotName string
	AuthId  int64
	ChatId  int64
}

func InitConfig() (*Config, error) {
	token := os.Getenv("TELEGRAM_TOKEN")
	botName := os.Getenv("TELEGRAM_BOT_NAME")
	chatIdStr := os.Getenv("TELEGRAM_CHAT_ID")
	authIdStr := os.Getenv("TELEGRAM_AUTH_ID")

	if authIdStr == "" || token == "" || botName == "" || chatIdStr == "" {
		logrus.Fatal("missing required telegram configuration", token, botName, chatIdStr, authIdStr)
	}

	authId, err := strconv.ParseInt(authIdStr, 10, 64)

	if err != nil {
		logrus.Fatal("telegram configuration incorrect")
	}

	chatId, err := strconv.ParseInt(chatIdStr, 10, 64)

	if err != nil {
		logrus.Fatal("telegram configuration incorrect")
	}

	cfg := &Config{
		Telegram: TelegramConfig{
			Token:   token,
			AuthId:  authId,
			ChatId:  chatId,
			BotName: botName,
		},
	}
	return cfg, nil
}
