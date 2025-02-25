package main

import (
	"github.com/sdimitrenco/grammurrr/internal/entities"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/pkg/logrus"
)

func main() {
	log := logging.NewLogger(logrus.NewLogrusLogger())

	
	entities.Test(log)

	test(log)
}

func test(log *logging.Logger) {
	log.Info("Приложение запущено 2")
}
