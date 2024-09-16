package config

import "github.com/caarlos0/env/v11"


type Config struct {
	Env string `env:"APP_ENV" envDefault:"development"`
	Url string `env:"APP_URL" envDefault:"http://localhost"`
	Port string `env:"APP_PORT" envDefault:":8080"`
	TimeZone string `env:"APP_TIMEZONE" envDefault:"Asia/Tokyo"`
	Log         string `env:"LOG" envDefault:"./logs/log/"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	LogRotation int    `env:"LOG_LEVEL" envDefault:"15"`
}

var config *Config = &Config{}

func GetConfig() (*Config, error) {
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}