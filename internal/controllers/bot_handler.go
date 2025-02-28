package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sdimitrenco/grammurrr/internal/domains"
)

func (b *BotControllerImpl) Start(answer domains.AnswerMessage) domains.AnswerMessage {
	text := "Привет! Я бот для заучивания слов.\n" +
		"/addgroup [название] — создать группу\n" +
		"/addword [группа] [иностранное слово] - [перевод] — добавить слово\n" +
		"/train [группа] — начать тренировку"

	answer.Message.Text = text
	answer.Message.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/addgroup"),
			tgbotapi.NewKeyboardButton("/addword"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/train"),
			tgbotapi.NewKeyboardButton("/start"),
		),
	)

	return answer

}
