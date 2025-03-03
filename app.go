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
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/storage"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/storage/postgres"
	"github.com/sdimitrenco/grammurrr/internal/usecases"
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

	db := storage.NewPostgresDB()

	groupRepo := postgres.NewGroupRepositoryPostgres(db)
	wordUseCase := usecases.NewGroupUseCase(groupRepo, log)
	controllers := controllers.NewBotController(log, wordUseCase)

	bot, err := adapters.NewTelegramAdapter(cfg.TelegramBot.Token, controllers)
	
	if err != nil {
		log.WithField("start bot", "can't create bot").Fatal(err)
	}

	go bot.Start()

	waitForShutdown(log)


}

func waitForShutdown(log *logging.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Info("Received shutdown signal, exiting...")
}
