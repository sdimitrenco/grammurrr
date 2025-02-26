package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/pkg/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	IsDebug     bool   `mapstructure:"is_debug"`
	WorkingDir  string `mapstructure:"working_dir"`
	Listen      ListenConfig `mapstructure:"listen"`
	DB          DBConfig `mapstructure:"db"`
	TelegramBot TelegramBotConfig `mapstructure:"telegram_bot"`
}

type ListenConfig struct {
	Type   string `mapstructure:"type"`
	BindIp string `mapstructure:"bind_ip"`
	Port   string `mapstructure:"port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbSsl    string `mapstructure:"db_ssl"`
	Type     string `mapstructure:"type"`
}

type TelegramBotConfig struct {
	Token string `mapstructure:"token"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logrusLogger := logrus.NewLogrusLogger()
		logger := logging.NewLogger(logrusLogger)
		instance = &Config{}

		v := viper.New()

		workingDir, err := os.Getwd()
		logger.Info("workingDir", workingDir)
		if err != nil {
			logger.Fatal(err)
		}

		workingDir = strings.Replace(workingDir, "/cmd/main", "", 1)
		instance.WorkingDir = workingDir

		// Load .env file if it exists
		envPath := filepath.Join(workingDir, ".env")
		if err := godotenv.Load(envPath); err != nil {
			logger.Warn(".env file not found. Skipping loading from .env file.")
		}

		configPath := filepath.Join(workingDir, "config.yml")
		v.SetConfigFile(configPath)
		v.SetConfigType("yaml")

		// Set default values
		v.SetDefault("is_debug", false)
		v.SetDefault("listen.type", "port")
		v.SetDefault("listen.bind_ip", "127.0.0.1")
		v.SetDefault("listen.port", "8080")
		v.SetDefault("db.host", "localhost")
		v.SetDefault("db.port", "5432")
		v.SetDefault("db.name", "grammurrr")
		v.SetDefault("db.user", "root")
		v.SetDefault("db.password", "12369874")
		v.SetDefault("db.db_ssl", "disable")
		v.SetDefault("db.type", "postgres")

		// Add env variables
		v.SetEnvPrefix("")
		v.AutomaticEnv()
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// Expand env variables in config file
		v.Set("db.host", os.ExpandEnv(v.GetString("db.host")))
		v.Set("db.port", os.ExpandEnv(v.GetString("db.port")))
		v.Set("db.name", os.ExpandEnv(v.GetString("db.name")))
		v.Set("db.user", os.ExpandEnv(v.GetString("db.user")))
		v.Set("db.password", os.ExpandEnv(v.GetString("db.password")))
		v.Set("db.db_ssl", os.ExpandEnv(v.GetString("db.db_ssl")))
		v.Set("db.type", os.ExpandEnv(v.GetString("db.type")))
		v.Set("telegram_bot.token", os.ExpandEnv(v.GetString("telegram_bot.token")))
		v.Set("listen.bind_ip", os.ExpandEnv(v.GetString("listen.bind_ip")))
		v.Set("listen.port", os.ExpandEnv(v.GetString("listen.port")))

		// Read the config file
		if err := v.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				logger.Warn("Config file not found. Using default and environment variables.")
			} else {
				logger.Fatal("Error reading config file: ", err)
			}
		}

		// Unmarshal the config into the struct
		if err := v.Unmarshal(instance); err != nil {
			logger.Fatal("Error unmarshalling config: ", err)
		}

		// Validate config
		if err := instance.validate(); err != nil {
			logger.Fatal("Config validation error: ", err)
		}

		logger.Info("Config successfully loaded")
	})
	return instance
}

func (c *Config) validate() error {
	if c.TelegramBot.Token == "" {
		return fmt.Errorf("Telegram bot token is required")
	}
	if c.DB.Type == "" {
		return fmt.Errorf("db type is required")
	}
	return nil
}
