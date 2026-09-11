package handlers

import (
	"net/http"
	"strconv"

	"github.com/aromalcode-prog/cab-share-backend/internal/constants"
	"github.com/aromalcode-prog/cab-share-backend/internal/dto/request"
	"github.com/aromalcode-prog/cab-share-backend/internal/dto/response"
	"github.com/aromalcode-prog/cab-share-backend/internal/models"
	"github.com/aromalcode-prog/cab-share-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type RideBookingHandler struct {
	service services.RideBookingService
}

func NewRideBookingHandler(service services.RideBookingService) *RideBookingHandler {
	return &RideBookingHandler{service: service}
}

func (h *RideBookingHandler) BookRide(c *gin.Context) {
	rideIDParam := c.Param("rideID")

	rideID, err := strconv.ParseUint(rideIDParam, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ride ID",
		})
		return
	}

	userId, err := constants.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request request.RideBookingRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	rideBooking, err := h.service.BookRide(uint(rideID), userId, request.Seats)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := RideBookingModelToResponse(rideBooking)
	c.JSON(http.StatusCreated, response)

}

func (h *RideBookingHandler) GetMyBookings(c *gin.Context) {

	userId, err := constants.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	bookings, err := h.service.GetMyBookings(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := []response.RideBookingResponseDTO{}
	for _, booking := range bookings {
		response = append(response, *RideBookingModelToResponse(&booking))
	}
	c.JSON(http.StatusOK, response)

}

func RideBookingModelToResponse(rideBooking *models.RideBooking) *response.RideBookingResponseDTO {
	return &response.RideBookingResponseDTO{
		BookingID:   rideBooking.ID,
		RideID:      rideBooking.RideID,
		PassengerID: rideBooking.PassengerID,
		SeatsBooked: rideBooking.SeatsBooked,
	}
}
