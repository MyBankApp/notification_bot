package bot

import (
	"notification_telegram_bot/buttons"
	"notification_telegram_bot/dto"
	"notification_telegram_bot/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type RegistrationState struct {
	Step          string
	CreateUserDto *dto.CreateUserDto
}

var registrationStates = make(map[int64]*RegistrationState)

type BotHandler struct {
	bot         *tgbotapi.BotAPI
	userService *service.UserService
}

func NewBotHandler(bot *tgbotapi.BotAPI, us *service.UserService) *BotHandler {
	return &BotHandler{
		bot:         bot,
		userService: us,
	}
}

func (h *BotHandler) Init() tgbotapi.UpdatesChannel {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := h.bot.GetUpdatesChan(updateConfig)

	return updates
}

func (h *BotHandler) Start() {
	for update := range h.Init() {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		messageText := update.Message.Text

		if state, exists := registrationStates[chatID]; exists {
			switch state.Step {
			case "username":
				state.CreateUserDto.Username = messageText
				state.Step = "password"
				msg := tgbotapi.NewMessage(chatID, "Введите пароль:")
				h.bot.Send(msg)
				continue

			case "password":
				state.CreateUserDto.Password = messageText
				userDto := h.userService.Create(state.CreateUserDto)
				
				if userDto == nil {
					msg := tgbotapi.NewMessage(chatID, "❌ Ошибка регистрации")
					h.bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(chatID, "✅ Регистрация завершена!")
					h.bot.Send(msg)
				}

				delete(registrationStates, chatID)
				continue
			}
		}

		if update.Message.IsCommand() {
			switch messageText {
			case "/start":
				msg := tgbotapi.NewMessage(chatID, "Добро пожаловать! Нажмите кнопку для регистрации.")
				keyboard := buttons.GetRegisterButton()
				msg.ReplyMarkup = keyboard
				h.bot.Send(msg)
			}
		} else {
			switch messageText {
			case "Register":
				registrationStates[chatID] = &RegistrationState{
					Step: "username",
					CreateUserDto: &dto.CreateUserDto{
						Telegram_id: update.Message.From.UserName,
					},
				}
				msg := tgbotapi.NewMessage(chatID, "Введите имя пользователя:")
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				h.bot.Send(msg)
			}
		}
	}
}
