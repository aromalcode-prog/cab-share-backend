package router

import (
	"github.com/aromalcode-prog/cab-share-backend/internal/handlers"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
	"github.com/aromalcode-prog/cab-share-backend/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/", handlers.HomeHandler)
	router.GET("/health", handlers.HealthHandler)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	router.POST("/register", userHandler.Register)

	return router
}
