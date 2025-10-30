package main

import (
	"fmt"
	"os"

	"github.com/21ess/animemaster/src/config"
	"github.com/21ess/animemaster/src/db"
	"github.com/21ess/animemaster/src/log"
	"github.com/21ess/animemaster/src/router"
)

func main() {
	config.LoadConfig()

	db.InitMongo()
	db.InitRedis()
	r := router.InitRouter()
	addr := fmt.Sprintf("%v:%v", os.Getenv("BASE_URL"), os.Getenv("PORT"))
	if err := r.Run(addr); err != nil {
		log.Log.Error(err.Error())
	}
}
