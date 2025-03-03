package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sdimitrenco/grammurrr/internal/adapters"
	"github.com/sdimitrenco/grammurrr/internal/config"
	"github.com/sdimitrenco/grammurrr/internal/controllers"
	"github.com/sdimitrenco/grammurrr/internal/domains"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/pkg/logrus"
)

func main() {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading location: %v\n", err)
		os.Exit(1)
	}
	time.Local = loc

	logrusLogger := logrus.NewLogrusLogger()
	log := logging.NewLogger(logrusLogger)

	cfg := config.GetConfig()

	bot, err := adapters.NewTelegramAdapter(cfg.TelegramBot.Token, controllers.NewBotController(log))
	if err != nil {
		log.WithField("start bot", "can't create bot").Fatal(err)
	}

	go bot.Start()

	a := 1
	b := 2

	fmt.Println(domains.Ternary(a > b, a, b))

	waitForShutdown(log)


}

func waitForShutdown(log *logging.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Info("Received shutdown signal, exiting...")
}
