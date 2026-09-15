package models

import "time"

type BookingStatus string

const (
	BookingStatusActive    BookingStatus = "active"
	BookingStatusCancelled BookingStatus = "cancelled"
)

type RideBooking struct {
	ID          uint          `gorm:"primaryKey"`
	RideID      uint          `gorm:"not null;index"`
	PassengerID uint          `gorm:"not null;index"`
	SeatsBooked uint          `gorm:"not null;default:1"`
	Status      BookingStatus `gorm:"not null;default:active"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
