package main

import (
	"github.com/sdimitrenco/grammurrr/internal/entities"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/pkg/logrus"
)

func main() {
	log := logging.NewLogger(logrus.NewLogrusLogger())

	log.Info("Приложение запущено")
	log.Debug("Приложение запущено")
	log.Warn("Приложение запущено")
	log.Error("Приложение запущено")
	entities.Test(log)

	test(log)
}

func test(log *logging.Logger) {
	log.Info("Приложение запущено 2")
}
