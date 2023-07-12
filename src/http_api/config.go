package http_api

import "github.com/spf13/viper"

type Config struct {
	Port     int    // The port to bind HTTP application API server to
	LogLevel string // The level of logging
}

func InitConfig() (*Config, error) {
	config := &Config{
		Port:     viper.GetInt("API.HTTP.Port"),
		LogLevel: viper.GetString("Log.Level"),
	}
	if config.Port == 0 {
		config.Port = 8080
	}
	return config, nil
}
