package controllers

import (
	"github.com/sdimitrenco/grammurrr/internal/domains"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/internal/usecases"
)

type BotController interface {
	HandleUpdate(message domains.Message, answer domains.AnswerMessage) domains.AnswerMessage
}

type BotControllerImpl struct {
	log         *logging.Logger
	wordService *usecases.GroupUseCase
}

func NewBotController(log *logging.Logger, servise *usecases.GroupUseCase) *BotControllerImpl {
	return &BotControllerImpl{
		log:         log,
		wordService: servise,
	}
}

func (b *BotControllerImpl) HandleUpdate(message domains.Message, answer domains.AnswerMessage) domains.AnswerMessage {
	switch message.Message.Text {
	case "/start":
		return b.Start(answer)
	case "/addgroup":
		return b.AddGroup(message, answer)
	case "/train":
		return b.Train(message, answer)
	default:
		return b.Default(message, answer)
	}
}
