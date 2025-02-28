package adapters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sdimitrenco/grammurrr/internal/controllers"
	"github.com/sdimitrenco/grammurrr/internal/domains"
)

type TelegramAdapter struct {
	bot        *tgbotapi.BotAPI
	controller controllers.BotController
}

type Message struct {
	Text string
}

func NewTelegramAdapter(token string, controller controllers.BotController) (bot *TelegramAdapter, err error) {

	tBot, err := tgbotapi.NewBotAPI(token)
	tBot.Debug = true

	if err != nil {
		return nil, err
	}

	return &TelegramAdapter{
		bot:        tBot,
		controller: controller,
	}, nil
}

func (t *TelegramAdapter) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message != nil {
			message := domains.Message{
				Message: update.Message,
			}
			response := t.controller.HandleUpdate(message)
			msg.Text = response.(string)

		}
		t.bot.Send(msg)
	}

	return nil
}
