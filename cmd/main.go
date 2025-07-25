package main

import (
	"go-bot/internal/adapters"
	"go-bot/internal/config"
	"go-bot/internal/logger"
)

func main() {
	log := logger.InitLogger()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tgAdapter, err := adapters.InitTg(cfg.Telegram.Token)
	if err != nil {
		log.Fatal("Error init telegram")
	}

	defer tgAdapter.TgUpdates(cfg.Telegram.AuthId)
}
