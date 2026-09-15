package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aromalcode-prog/cab-share-backend/internal/constants"
	"github.com/aromalcode-prog/cab-share-backend/internal/dto/request"
	"github.com/aromalcode-prog/cab-share-backend/internal/dto/response"
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
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrDepartureTimeInPast) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "ride created successfully"})
}

func (h *RideHandler) GetRideByID(c *gin.Context) {
	rideID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ride ID"})
		return
	}

	ride, err := h.rideService.GetRideByID(uint(rideID))
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrRideNotFound) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.FromRideModel(*ride))
}

func (h *RideHandler) GetAvailableRides(c *gin.Context) {
	rides, err := h.rideService.GetAvailableRides()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var rideResponses []response.RideResponseDTO

	for _, ride := range rides {
		rideResponses = append(
			rideResponses,
			response.FromRideModel(ride),
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"rides": rideResponses,
	})
}

func (h *RideHandler) SearchRides(c *gin.Context) {
	var req request.SearchRideRequestDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rides, err := h.rideService.SearchRides(req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, services.ErrSourceDestinationRequired) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	rideResponses := []response.RideResponseDTO{}
	for _, ride := range rides {
		rideResponses = append(rideResponses, response.FromRideModel(ride))
	}

	c.JSON(http.StatusOK, rideResponses)
}
