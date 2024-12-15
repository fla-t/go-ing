package booking_test

import (
	"testing"

	"github.com/fla-t/go-ing/internal/domain/booking"
	service "github.com/fla-t/go-ing/internal/services/booking"
	"github.com/fla-t/go-ing/internal/uow/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	uow := inmemory.NewFakeUnitOfWork()
	service := service.NewService(uow)

	b, err := service.CreateBooking("1", "A", "B", 10.0, 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, b)

	savedBooking, err := service.GetBookingByID(b.ID)
	assert.Nil(t, err)
	assert.Equal(t, b, savedBooking)
}

func TestUpdateRide(t *testing.T) {
	uow := inmemory.NewFakeUnitOfWork()
	service := service.NewService(uow)

	b, err := service.CreateBooking("1", "A", "B", 10.0, 100.0)
	if err != nil {
		t.Fatal(err)
	}

	ride := booking.Ride{ID: b.Ride.ID, Source: "A", Destination: "C", Distance: 15.0, Cost: 150.0}

	err = service.UpdateRide(&ride)
	assert.Nil(t, err)

	updatedBooking, err := service.GetBookingByID(b.ID)
	assert.Nil(t, err)
	assert.Equal(t, ride, updatedBooking.Ride)
}
