package router

import (
	"github.com/aromalcode-prog/cab-share-backend/config"
	"github.com/aromalcode-prog/cab-share-backend/internal/auth"
	"github.com/aromalcode-prog/cab-share-backend/internal/handlers"
	"github.com/aromalcode-prog/cab-share-backend/internal/middleware"
	"github.com/aromalcode-prog/cab-share-backend/internal/repositories"
	"github.com/aromalcode-prog/cab-share-backend/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.GET("/", handlers.HomeHandler)
	router.GET("/health", handlers.HealthHandler)

	jwtManager := auth.NewJWTManager(cfg)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, jwtManager)
	userHandler := handlers.NewUserHandler(userService)
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
	router.GET("/getusers", userHandler.GetAllUsers) // only for testing purpose

	authMiddleware := middleware.NewAuthMiddleware(jwtManager)

	rideRepo := repositories.NewRideRepository(db)
	rideService := services.NewRideService(rideRepo)
	rideHandler := handlers.NewRideHandler(rideService)
	router.GET("/getrides", rideHandler.GetAvailableRides)
	router.POST("/rides", authMiddleware.Authenticate(), rideHandler.CreateRide)

	rideBookingRepository := repositories.NewRideBookingRepository(db)
	rideBookingService := services.NewRideBookingService(db, rideRepo, rideBookingRepository)
	rideBookingHandler := handlers.NewRideBookingHandler(rideBookingService)
	router.GET("/bookings", authMiddleware.Authenticate(), rideBookingHandler.GetMyBookings)
	router.POST("/rides/:rideID/book", authMiddleware.Authenticate(), rideBookingHandler.BookRide)
	router.POST("/bookings/:bookingID/cancel", authMiddleware.Authenticate(), rideBookingHandler.CancelBooking)

	return router
}
