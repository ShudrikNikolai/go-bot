package main

import (
	"fmt"
	"go-bot/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, _ := config.LoadConfig()
	fmt.Print(cfg.Telegram.AuthId + 9999)
}
