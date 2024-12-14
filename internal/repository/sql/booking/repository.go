package repository

import (
	"database/sql"

	booking "github.com/fla-t/go-ing/internal/domain/booking"
)

// Repository is a struct that holds the database connection
type Repository struct {
	tx *sql.Tx
}

// NewBookingRepository creates a new BookingRepository
func NewBookingRepository(tx *sql.Tx) *Repository {
	return &Repository{
		tx: tx,
	}
}

// CreateBooking creates a new booking
func (r *Repository) CreateBooking(b *booking.Booking) error {
	bookingQuery := `
		insert into bookings (id, user_id, ride_id, time) 
		values ($1, $2, $3, $4);
	`
	_, err := r.tx.Exec(bookingQuery, b.ID, b.UserID, b.Ride.ID, b.Time)

	if err != nil {
		return err
	}

	query := `
		insert into rides (id, source, destination, distance, cost) 
		values ($1, $2, $3, $4, $5);
	`
	_, err = r.tx.Exec(query, b.Ride.ID, b.Ride.Source, b.Ride.Destination, b.Ride.Distance, b.Ride.Cost)

	if err != nil {
		return err
	}

	return nil
}

// GetBookingByID returns a booking by its id
func (r *Repository) GetBookingByID(bookingID string) (*booking.Booking, error) {
	var b booking.Booking
	var ride booking.Ride

	err := r.tx.QueryRow(`
		select 
			b.id, b.user_id, b.time, 
			r.ride_id, r.source, r.destination, r.distance, r.cost
		from bookings b
		join rides r on b.ride_id = r.ride_id
		where b.booking_id = ?;
	`, bookingID).
		Scan(&b.ID, &b.UserID, &b.Time, &ride.ID, &ride.Source, &ride.Destination, &ride.Distance, &ride.Cost)

	if err != nil {
		return nil, err
	}

	b.Ride = ride
	return &b, nil
}

// UpdateRide updates a ride
func (r *Repository) UpdateRide(b *booking.Ride) error {
	query := `
		update rides
		set source = $1, destination = $2, distance = $3, cost = $4
		where ride_id = $5;
	`
	_, err := r.tx.Exec(query, b.Source, b.Destination, b.Distance, b.Cost, b.ID)

	if err != nil {
		return err
	}

	return nil
}
