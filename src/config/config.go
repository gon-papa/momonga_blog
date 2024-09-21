package config

import (
	"sync"

	"github.com/caarlos0/env/v11"
)


type Config struct {
	Env string `env:"APP_ENV" envDefault:"development"`
	Url string `env:"APP_URL" envDefault:"http://localhost"`
	Port string `env:"APP_PORT" envDefault:":8080"`
	TimeZone string `env:"APP_TIMEZONE" envDefault:"Asia/Tokyo"`
	Log         string `env:"LOG" envDefault:"./logs/log/"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	LogRotation int    `env:"LOG_LEVEL" envDefault:"15"`
}

var (
    ConfigInstance *Config
    once           sync.Once
    initErr        error
)


func GetConfig() (*Config, error) {
    once.Do(func() {
        ConfigInstance = &Config{}
        initErr = env.Parse(ConfigInstance)
    })
    if initErr != nil {
        return nil, initErr
    }
    return ConfigInstance, nil
}