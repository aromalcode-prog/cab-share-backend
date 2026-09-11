package response

type RideBookingResponseDTO struct {
	BookingID   uint `json:"booking_id"`
	RideID      uint `json:"ride_id"`
	PassengerID uint `json:"passenger_id"`
	SeatsBooked uint `json:"seats_booked"`
}
