package request

type RideBookingRequestDTO struct {
	Seats uint `json:"seats" binding:"required,min=1"`
}
