package booking

import "github.com/google/uuid"

// Booking is the aggregate root for rides in the system
type Booking struct {
	ID     string
	userID string
	Ride   Ride
	time   string
}

// NewBooking creates a new booking
func NewBooking(userID string, ride Ride, time string) *Booking {
	return &Booking{
		ID:     uuid.New().String(),
		userID: userID,
		Ride:   ride,
		time:   time,
	}
}
