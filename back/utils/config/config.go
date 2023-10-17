package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"nixietech/internal/config"
	"sync"
)

var instance *config.Config
var once sync.Once

func GetConfig(configName string) *config.Config {
	once.Do(func() {
		slog.Info("read application configuration")
		instance = &config.Config{}
		if err := cleanenv.ReadConfig(configName, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			slog.Error(help)
			slog.Error(err.Error())
		}
	})
	return instance
}
