package main

import (
	"fmt"

	"github.com/aromalcode-prog/cab-share-backend/config"
	"github.com/aromalcode-prog/cab-share-backend/internal/database"
	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"github.com/aromalcode-prog/cab-share-backend/internal/router"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println("Config loaded")

	db, err := database.ConnectDB(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	// database.Execute(db)
	router := router.SetupRouter(db, cfg)
	err = router.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
