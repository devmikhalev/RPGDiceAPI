package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env     string `yaml:"env" env:"ENV" env-default:"prod"`
	Address string `yaml:"address" env:"ADDRESS" env-required:"true"`
}

func LoadConfig(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("Config file with path doesn't exist: %s", path)
	}
	var conf Config
	if err := cleanenv.ReadConfig(path, &conf); err != nil {
		log.Fatalf("Config file cannot be read: %s", err)
	}
	return &conf
}
