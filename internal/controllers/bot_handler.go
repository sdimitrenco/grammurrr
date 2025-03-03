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
	answer.Message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить группу", "/addgroup"),
			tgbotapi.NewInlineKeyboardButtonData("Добавить слово", "/addword"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Тренировка", "/train"),
			tgbotapi.NewInlineKeyboardButtonData("Старт", "/start"),
		),
	)

	return answer

}

func (b *BotControllerImpl) AddGroup(message domains.Message, answer domains.AnswerMessage) domains.AnswerMessage {

	answer.Message.Text = "Группа добавлена"
	return answer
}

/*
func (b *BotControllerImpl) AddWord(message domains.Message, answer domains.AnswerMessage) domains.AnswerMessage {

	word := domains.Word{
		Word: message.Message.Text,
		Lang: "en",
	}

	b.wordService.AddWord(word)
	answer.Message.Text = "Слово добавлено"

	return answer
}
*/

func (b *BotControllerImpl) Train(message domains.Message, answer domains.AnswerMessage) domains.AnswerMessage {
	answer.Message.Text = "Тренировка начата"
	return answer
}

func (b *BotControllerImpl) Default(message domains.Message, answer domains.AnswerMessage) domains.AnswerMessage {

	answer.Message.Text = "Неизвестная команда"
	return answer
}
