package main

import (
	"github.com/aromalcode-prog/cab-share-backend/config"
	"github.com/aromalcode-prog/cab-share-backend/internal/router"
)

func main() {
	config := config.LoadConfig()
	router := router.SetupRouter()
	err := router.Run(":" + config.Port)
	if err != nil {
		panic(err)
	}
}
