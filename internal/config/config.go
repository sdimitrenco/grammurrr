package config

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
)

type Config struct {
	IsDebug     *bool  `yaml:"is_debug" env:"IS_DEBUG" env-default:"false"`
	WorkingDir string  `yaml:"working_dir"`
	Listen      struct {
		Type   string `yaml:"type"  env-default:"port"`
		BindIp string `yaml:"bindIp" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	DB struct {
		Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
		Port     string `yaml:"port" env:"DB_PORT" env-default:"27017"`
		Name   string `yaml:"name" env:"DB_NAME" env-default:""`
		User     string `yaml:"user" env:"DB_USER" env-default:""`
		Password string `yaml:"password" env:"DB_PASSWORD" env-default:""`
		DbSsl string `yaml:"db_ssl" env:"DB_SSL" env-default:""`
		Type string `yaml:"type" env:"DB_TYPE" env-default:""`
	} `yaml:"db"`
	TelegramBot struct {
		Token string `yaml:"token" env:"TELEGRAM_BOT_TOKEN"`
	} `yaml:"telegram_bot"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logging.Logger) *Config {
	once.Do(func() {
		logger.Info("Get application configuration")
		instance = &Config{}

		workingDir, err := os.Getwd()
		logger.Info("workingDir", workingDir)
		if err != nil {
			logger.Fatal(err)
		}

		workingDir = strings.Replace(workingDir, "/cmd/main", "", 1)

		instance.WorkingDir = workingDir

		configPath := filepath.Join(workingDir, "config.yml")

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			errorDescription, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(errorDescription)
			logger.Fatal(err)
		}
	})
	return instance
}
