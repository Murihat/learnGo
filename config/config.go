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
	// 1️⃣ Selalu baca ENV dari OS (Zeabur, Docker, dll)
	viper.AutomaticEnv()

	// 2️⃣ .env hanya untuk local (opsional)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// 3️⃣ Fail fast hanya untuk variable WAJIB
	if cfg.DBConn == "" {
		log.Fatal("DB_CONN is required but not set")
	}

	// 4️⃣ Default port (Zeabur biasanya inject PORT)
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg
}
