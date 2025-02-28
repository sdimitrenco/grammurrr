package tgbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotAPI struct {
	Bot        *tgbotapi.BotAPI
	NewMessage tgbotapi.MessageConfig
	NewUpdate  tgbotapi.UpdateConfig
}

func NewBotAPI(token string) *BotAPI {
	tBot, err := tgbotapi.NewBotAPI(token)

	tBot.Debug = true

	if err != nil {
		panic(err)
	}

	return &BotAPI{
		Bot: tBot,
	}
}

func (b *BotAPI) SetNewUpdate(offset int) *BotAPI {
	b.NewUpdate = tgbotapi.NewUpdate(offset)
	return b
}

func (b *BotAPI) SetNewMessage(chatID int64, text string) *BotAPI {
	b.NewMessage = tgbotapi.NewMessage(chatID, text)
	return b
}
