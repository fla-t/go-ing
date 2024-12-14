package booking

import (
	"time"

	"github.com/google/uuid"
)

// Booking is the aggregate root for rides in the system
type Booking struct {
	ID     string
	UserID string
	Ride   Ride
	Time   time.Time // time of booking in UTC
}

// NewBooking creates a new booking
func NewBooking(userID string, ride Ride) *Booking {
	return &Booking{
		ID:     uuid.New().String(),
		UserID: userID,
		Ride:   ride,
		Time:   time.Now().UTC(),
	}
}
