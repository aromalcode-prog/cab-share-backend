package response

import "github.com/aromalcode-prog/cab-share-backend/internal/models"

type RideBookingResponseDTO struct {
	BookingID   uint                 `json:"booking_id"`
	RideID      uint                 `json:"ride_id"`
	PassengerID uint                 `json:"passenger_id"`
	SeatsBooked uint                 `json:"seats_booked"`
	Status      models.BookingStatus `json:"status"`
}
