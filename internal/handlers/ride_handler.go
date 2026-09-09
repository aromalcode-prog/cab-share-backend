package handlers

import (
	"net/http"

	"github.com/aromalcode-prog/cab-share-backend/internal/constants"
	"github.com/aromalcode-prog/cab-share-backend/internal/dto/request"
	"github.com/aromalcode-prog/cab-share-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type RideHandler struct {
	rideService services.RideService
}

func NewRideHandler(rideService services.RideService) *RideHandler {
	return &RideHandler{
		rideService: rideService,
	}
}

func (h *RideHandler) CreateRide(c *gin.Context) {
	var req request.CreateRideRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := constants.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve authenticated user"})
		return
	}

	if err := h.rideService.CreateRide(req, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "ride created successfully"})
}
