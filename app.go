package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sdimitrenco/grammurrr/internal/config"
	"github.com/sdimitrenco/grammurrr/internal/entities"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/pkg/logrus"
)

func main() {
	// Устанавливаем часовой пояс на Europe/Berlin
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading location: %v\n", err)
		os.Exit(1)
	}
	time.Local = loc

	// Create base logger
	logrusLogger := logrus.NewLogrusLogger()
	log := logging.NewLogger(logrusLogger)

	cfg := config.GetConfig()

	fmt.Println(cfg.DB.Host)
	fmt.Println(cfg.DB.Password)

	// Create logger with field
	logWithField := log.WithField("key", "value")

	// Log error with field
	logWithField.Error("bad something")

	// Pass logger with field
	entities.Test(logWithField)

	// Pass logger with field
	test(logWithField)
}

func test(log *logging.Logger) {
	log.Info("Приложение запущено 2")
}
