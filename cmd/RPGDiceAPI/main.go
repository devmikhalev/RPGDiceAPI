package main

import (
	"RPRDiceAPI/internal/config"
	"RPRDiceAPI/internal/rolldice"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	//init config
	conf := config.LoadConfig(os.Getenv("CONFIG_PATH"))

	//create server
	router := gin.Default()
	router.GET("/roll/:faces/:amount/:mod", rolldice.MakeRoll)

	err := router.Run(conf.Address)
	if err != nil {
		log.Fatalf("Server cannot run: %s", err)
	}
}
