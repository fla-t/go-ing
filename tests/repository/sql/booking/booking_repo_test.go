package booking

import (
	"testing"
	"time"

	"github.com/fla-t/go-ing/internal/domain/booking"
	domain "github.com/fla-t/go-ing/internal/domain/booking"
	"github.com/fla-t/go-ing/internal/uow/sql"
	"github.com/fla-t/go-ing/tests"
	"github.com/google/uuid"
)

func TestCreateBooking(t *testing.T) {
	// Setup the test database
	db, cleanup, err := tests.SetupTestDatabase()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	defer cleanup()

	// Initialize db unit of work
	uow := sql.NewDbUnitOfWork(db)
	uow.Begin()

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
	err = uow.BookingRepository().CreateBooking(&booking)
	if err != nil {
		t.Fatalf("failed to create booking: %v", err)
	}

	// Commit the transaction
	err = uow.Commit()
	if err != nil {
		t.Fatalf("failed to commit transaction: %v", err)
	}

	// Verify the record in the database
	var count int
	err = db.QueryRow("select count(*) from bookings where id = $1", booking.ID).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query database: %v", err)
	}

	if count != 1 {
		t.Errorf("expected 1 booking, got %d", count)
	}
}

func TestGetBookingByID(t *testing.T) {
	// Setup the test database
	db, cleanup, err := tests.SetupTestDatabase()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	defer cleanup()

	// Initialize db unit of work
	uow := sql.NewDbUnitOfWork(db)
	uow.Begin()

	// Test data
	ride := booking.Ride{
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
	err = uow.BookingRepository().CreateBooking(&booking)
	if err != nil {
		t.Fatalf("failed to create booking: %v", err)
	}

	// Perform the GetBookingByID operation
	b, err := uow.BookingRepository().GetBookingByID(booking.ID)
	if err != nil {
		t.Fatalf("failed to get booking by id: %v", err)
	}

	// Commit the transaction
	err = uow.Commit()
	if err != nil {
		t.Fatalf("failed to commit transaction: %v", err)
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
	// Setup the test database
	db, cleanup, err := tests.SetupTestDatabase()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	defer cleanup()

	// Initialize db unit of work
	uow := sql.NewDbUnitOfWork(db)
	uow.Begin()

	// Test data
	ride := booking.Ride{
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
	err = uow.BookingRepository().CreateBooking(&booking)
	if err != nil {
		t.Fatalf("failed to create booking: %v", err)
	}

	// Update the ride
	ride.Cost = 200
	err = uow.BookingRepository().UpdateRide(&ride)
	if err != nil {
		t.Fatalf("failed to update ride: %v", err)
	}

	// Commit the transaction
	err = uow.Commit()
	if err != nil {
		t.Fatalf("failed to commit transaction: %v", err)
	}

	// Verify the updated record in the database
	var cost float64
	err = db.QueryRow("select cost from rides where id = $1", ride.ID).Scan(&cost)
	if err != nil {
		t.Fatalf("failed to query database: %v", err)
	}

	if cost != 200 {
		t.Errorf("expected cost 200, got %f", cost)
	}
}
