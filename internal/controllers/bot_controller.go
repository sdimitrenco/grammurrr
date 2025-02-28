package controllers

import (
	"github.com/sdimitrenco/grammurrr/internal/domains"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
)

type BotController interface {
	HandleUpdate(update domains.Message, answer domains.AnswerMessage) domains.AnswerMessage
}

type BotControllerImpl struct {
	log *logging.Logger
}

func NewBotController(log *logging.Logger) *BotControllerImpl {
	return &BotControllerImpl{
		log: log,
	}
}

func (b *BotControllerImpl) HandleUpdate(update domains.Message, answer domains.AnswerMessage) domains.AnswerMessage {
	switch update.Message.Text {
	case "/start":
		return b.Start(answer)
	default:
		return answer
	}
}
