package config

import (
	"github.com/21ess/animemaster/src/log"
	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load("/Users/yuhao/Projects/acgccb/animemaster/.env"); err != nil {
		log.Log.Error(err.Error())
	}
	log.Log.Info("load config success")
}

type Config struct {
}
