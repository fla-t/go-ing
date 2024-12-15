package booking

import (
	"testing"
	"time"

	"github.com/fla-t/go-ing/internal/domain/booking"
	repository "github.com/fla-t/go-ing/internal/repository/inmemory/booking"
	"github.com/google/uuid"
)

func TestCreateBooking(t *testing.T) {
	// Initialize in-memory repository
	repo := repository.NewInMemoryBookingRepository()

	// Test data
	ride := booking.Ride{
		ID:          uuid.New().String(),
		Source:      "Location A",
		Destination: "Location B",
		Distance:    50,
		Cost:        150,
	}
	booking := booking.Booking{
		ID:     uuid.New().String(),
		UserID: "user123",
		Ride:   ride,
		Time:   time.Now().UTC(),
	}

	// Perform the CreateBooking operation
	err := repo.CreateBooking(&booking)
	if err != nil {
		t.Fatalf("failed to create booking: %v", err)
	}

	// Verify the record in the repository
	b, err := repo.GetBookingByID(booking.ID)
	if err != nil {
		t.Fatalf("failed to get booking by id: %v", err)
	}

	if b.ID != booking.ID {
		t.Errorf("expected booking id %s, got %s", booking.ID, b.ID)
	}
	if b.UserID != booking.UserID {
		t.Errorf("expected user id %s, got %s", booking.UserID, b.UserID)
	}
	if b.Ride.ID != booking.Ride.ID {
		t.Errorf("expected ride id %s, got %s", booking.Ride.ID, b.Ride.ID)
	}
	if !b.Time.Truncate(time.Microsecond).Equal(booking.Time.Truncate(time.Microsecond)) {
		t.Errorf("expected time %v, got %v", booking.Time, b.Time)
	}
}

func TestGetBookingByID(t *testing.T) {
	// Initialize in-memory repository
	repo := repository.NewInMemoryBookingRepository()

	// Test data
	ride := booking.Ride{
		ID:          uuid.New().String(),
		Source:      "Location A",
		Destination: "Location B",
		Distance:    50,
		Cost:        150,
	}
	booking := booking.Booking{
		ID:     uuid.New().String(),
		UserID: "user123",
		Ride:   ride,
		Time:   time.Now().UTC(),
	}

	// Perform the CreateBooking operation
	err := repo.CreateBooking(&booking)
	if err != nil {
		t.Fatalf("failed to create booking: %v", err)
	}

	// Perform the GetBookingByID operation
	b, err := repo.GetBookingByID(booking.ID)
	if err != nil {
		t.Fatalf("failed to get booking by id: %v", err)
	}

	// Verify the returned booking
	if b.ID != booking.ID {
		t.Errorf("expected booking id %s, got %s", booking.ID, b.ID)
	}
	if b.UserID != booking.UserID {
		t.Errorf("expected user id %s, got %s", booking.UserID, b.UserID)
	}
	if b.Ride.ID != booking.Ride.ID {
		t.Errorf("expected ride id %s, got %s", booking.Ride.ID, b.Ride.ID)
	}
	if !b.Time.Truncate(time.Microsecond).Equal(booking.Time.Truncate(time.Microsecond)) {
		t.Errorf("expected time %v, got %v", booking.Time, b.Time)
	}
}
