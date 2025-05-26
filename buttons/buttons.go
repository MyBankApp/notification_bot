package buttons

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetRegisterButton() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Register"),
		),
	)

	return keyboard
}
