package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	DBConn string
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	cfg := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	if cfg.DBConn == "" {
		log.Fatal("DB_CONN is empty — check .env")
	}

	return cfg
}
