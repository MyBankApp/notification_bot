package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	token := os.Getenv("API_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates := bot.GetUpdatesChan(u)

	for update := range updates {
        if update.Message != nil {
            log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

            msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
            bot.Send(msg)
        }
    }
}
