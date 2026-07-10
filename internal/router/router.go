package router

import (
	"github.com/aromalcode-prog/cab-share-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", handlers.HomeHandler)
	router.GET("/health", handlers.HealthHandler)

	return router
}
