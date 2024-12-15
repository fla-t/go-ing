package repository

import (
	"errors"

	booking "github.com/fla-t/go-ing/internal/domain/booking"
)

// Repository is a simple in-memory repository for bookings for testing
type Repository struct {
	bookings map[string]*booking.Booking
}

// NewInMemoryBookingRepository creates a new BookingRepository
func NewInMemoryBookingRepository() *Repository {
	return &Repository{
		bookings: make(map[string]*booking.Booking),
	}
}

// CreateBooking creates a new booking
func (r *Repository) CreateBooking(b *booking.Booking) error {
	r.bookings[b.ID] = b
	return nil
}

// GetBookingByID returns a booking by its id
func (r *Repository) GetBookingByID(id string) (*booking.Booking, error) {
	b, ok := r.bookings[id]
	if !ok {
		return nil, errors.New("Booking Not Found")
	}
	return b, nil
}

// UpdateRide updates a ride
func (r *Repository) UpdateRide(ride *booking.Ride) error {
	for _, b := range r.bookings {
		if b.Ride.ID == ride.ID {
			// Update the Ride object
			b.Ride = *ride
			return nil
		}
	}

	return errors.New("Ride Not Found")
}
