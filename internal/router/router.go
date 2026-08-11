package router

import (
	"github.com/aromalcode-prog/cab-share-backend/config"
	"github.com/aromalcode-prog/cab-share-backend/internal/auth"
	"github.com/aromalcode-prog/cab-share-backend/internal/handlers"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
	"github.com/aromalcode-prog/cab-share-backend/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.GET("/", handlers.HomeHandler)
	router.GET("/health", handlers.HealthHandler)

	userRepo := repositories.NewUserRepository(db)
	jwtManager := auth.NewJWTManager(cfg)
	userService := services.NewUserService(userRepo, jwtManager)
	userHandler := handlers.NewUserHandler(userService)
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	return router
}
