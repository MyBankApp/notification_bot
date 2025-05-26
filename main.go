package main

import (
	"log"
	"notification_telegram_bot/bot"
	"notification_telegram_bot/config"
	"notification_telegram_bot/db"
	"notification_telegram_bot/repository"
	"notification_telegram_bot/service"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.LoadConfig()
	
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	botAPI, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	db := db.Connect(cfg.DatabaseURL)

	repository := repository.NewUserRepository(db)
	userService := service.NewUserService(*repository)
	
	if err != nil {
		log.Panic(err)
	}

	botHandler := bot.NewBotHandler(botAPI, userService)

	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	botHandler.Start()
}
