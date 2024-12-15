package booking

import (
	"testing"
	"time"

	domain "github.com/fla-t/go-ing/internal/domain/booking"
	repository "github.com/fla-t/go-ing/internal/repository/inmemory/booking"
	"github.com/google/uuid"
)

func TestCreateBooking(t *testing.T) {
	// Initialize in-memory repository
	repo := repository.NewInMemoryBookingRepository()

	// Test data
	ride := domain.Ride{
		ID:          uuid.New().String(),
		Source:      "Location A",
		Destination: "Location B",
		Distance:    50,
		Cost:        150,
	}
	booking := domain.Booking{
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
	ride := domain.Ride{
		ID:          uuid.New().String(),
		Source:      "Location A",
		Destination: "Location B",
		Distance:    50,
		Cost:        150,
	}
	booking := domain.Booking{
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

func TestUpdateRide(t *testing.T) {
	// Initialize in-memory repository
	repo := repository.NewInMemoryBookingRepository()

	// Test data
	ride := domain.Ride{
		ID:          uuid.New().String(),
		Source:      "Location A",
		Destination: "Location B",
		Distance:    50,
		Cost:        150,
	}
	booking := domain.Booking{
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

	// Update the ride
	updatedRide := domain.Ride{
		ID:          ride.ID,
		Source:      "Location A",
		Destination: "Location C",
		Distance:    75,
		Cost:        225,
	}

	// Perform the UpdateRide operation
	err = repo.UpdateRide(&updatedRide)
	if err != nil {
		t.Fatalf("failed to update ride: %v", err)
	}

	// Verify the updated ride
	b, err := repo.GetBookingByID(booking.ID)
	if err != nil {
		t.Fatalf("failed to get booking by id: %v", err)
	}

	if b.Ride.ID != updatedRide.ID {
		t.Errorf("expected ride id %s, got %s", updatedRide.ID, b.Ride.ID)
	}
	if b.Ride.Source != updatedRide.Source {
		t.Errorf("expected source %s, got %s", updatedRide.Source, b.Ride.Source)
	}
	if b.Ride.Destination != updatedRide.Destination {
		t.Errorf("expected destination %s, got %s", updatedRide.Destination, b.Ride.Destination)
	}
	if b.Ride.Distance != updatedRide.Distance {
		t.Errorf("expected distance %f, got %f", updatedRide.Distance, b.Ride.Distance)
	}
	if b.Ride.Cost != updatedRide.Cost {
		t.Errorf("expected cost %f, got %f", updatedRide.Cost, b.Ride.Cost)
	}
}
