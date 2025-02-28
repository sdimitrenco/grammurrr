package adapters

import (
	"github.com/sdimitrenco/grammurrr/internal/controllers"
	"github.com/sdimitrenco/grammurrr/internal/domains"
	"github.com/sdimitrenco/grammurrr/pkg/tgbot"
)

type TelegramAdapter struct {
	bot        *tgbot.BotAPI
	controller controllers.BotController
}

type Message struct {
	Text string
}

func NewTelegramAdapter(token string, controller controllers.BotController) (bot *TelegramAdapter, err error) {

	tBot := tgbot.NewBotAPI(token)

	return &TelegramAdapter{
		bot:        tBot,
		controller: controller,
	}, nil
}

func (t *TelegramAdapter) Start() error {

	u := t.bot.SetNewUpdate(0).NewUpdate
	u.Timeout = 60

	updates := t.bot.Bot.GetUpdatesChan(u)

	for update := range updates {
		answer := domains.AnswerMessage{
			Message: t.bot.SetNewMessage(update.Message.Chat.ID, "").NewMessage,
		}

		if update.Message != nil {
			message := domains.Message{
				Message: update.Message,
			}
			answer.Message = t.controller.HandleUpdate(message, answer).Message

		}
		t.bot.Bot.Send(answer.Message)
	}

	return nil
}
