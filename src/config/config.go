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
	// LogRotation int    `env:"LOG_ROTATION" envDefault:"15"`
    DbHost string `env:"DB_HOST" envDefault:"mysql"`
    DbPort string `env:"DB_PORT" envDefault:"3306"`
    DbUser string `env:"DB_USER" envDefault:"user"`
    DbPassword string `env:"DB_PASSWORD" envDefault:"password"`
    DbName string `env:"DB_NAME" envDefault:"db"`
    Dbms string `env:"DBMS" envDefault:"mysql"`
    DbMaxIdleConn int `env:"DB_MAX_IDLE_CONN" envDefault:"1"`
    DbMaxOpenConn int `env:"DB_MAX_OPEN_CONN" envDefault:"10"`
    DbConnMaxLifetime int `env:"DB_CONN_MAX_LIFETIME" envDefault:"1"`
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