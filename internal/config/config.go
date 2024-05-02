package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
	"time"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

type Config struct {
	Env         string      `yaml:"env"`
	StoragePath StoragePath `yaml:"storage_path"`
	Grpc        Grpc        `yaml:"grpc"`
}

type StoragePath struct {
	Password string `yaml:"POSTGRES_PASSWORD"`
	Host     string `yaml:"POSTGRES_HOST"`
	User     string `yaml:"POSTGRES_USER"`
	Db       string `yaml:"POSTGRES_DB"`
	Port     string `yaml:"POSTGRES_PORT"`
}

type Grpc struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	fetchPath := fetchConfigPath()
	if fetchPath == "" {
		panic("Пустой файл конфигурации")
	}

	return MustLoadByPath(fetchPath)
}

func MustLoadByPath(fetchPath string) *Config {
	if _, err := os.Stat(fetchPath); os.IsNotExist(err) {
		panic("Не существует файл конфигурации: " + fetchPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(fetchPath, &cfg); err != nil {
		panic("Ошибка чтения конфига" + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "config path")
	flag.Parse()

	if res != "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

func SetupLoger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
