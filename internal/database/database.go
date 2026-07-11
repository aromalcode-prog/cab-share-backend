package database

import (
	"fmt"

	"github.com/aromalcode-prog/cab-share-backend/config"
	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, err
}

func Execute(db *gorm.DB) {
	var FetchedUser models.User
	var FetchedUser1 models.User
	res := db.First(&FetchedUser, 1)
	fmt.Println(FetchedUser)
	fmt.Println(res.Error)
	res1 := db.First(&FetchedUser1, 100)
	fmt.Println(FetchedUser1)
	fmt.Println(res1.Error)
}
