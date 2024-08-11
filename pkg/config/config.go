// config.go
package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    ServerPort string
    // other config fields
}

func LoadConfig() Config {
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error loading config file: %v", err)
    }

    return Config{
        ServerPort: viper.GetString("SERVER_PORT"),
        // other config values
    }
}
