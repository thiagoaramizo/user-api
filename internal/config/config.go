package config

import (
    "github.com/spf13/viper"
    "log"
)

func LoadConfig() {
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error loading .env file")
    }
}