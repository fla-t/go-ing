package booking_test

import (
	"testing"

	aclModels "github.com/fla-t/go-ing/internal/acl/user"
	acl "github.com/fla-t/go-ing/internal/acl/user/inmemory"
	"github.com/fla-t/go-ing/internal/domain/booking"
	service "github.com/fla-t/go-ing/internal/services/booking"
	"github.com/fla-t/go-ing/internal/uow/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	uow := inmemory.NewFakeUnitOfWork()
	userACL := acl.NewInMemoryUserACL()
	service := service.NewService(uow, userACL)

	userACL.AddUser(&aclModels.User{ID: "1", Name: "John Doe", Email: "john@example.com"})
	b, err := service.CreateBooking("1", "A", "B", 10.0, 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, b)

	savedBooking, err := service.GetBookingByID(b.ID)
	assert.Nil(t, err)
	assert.Equal(t, b.ID, savedBooking.ID)
	assert.Equal(t, b.Ride.Source, savedBooking.Ride.Source)
	assert.Equal(t, b.Ride.Destination, savedBooking.Ride.Destination)
	assert.Equal(t, b.Ride.Distance, savedBooking.Ride.Distance)
	assert.Equal(t, b.Ride.Cost, savedBooking.Ride.Cost)
	assert.Equal(t, b.Time, savedBooking.Time)
}

func TestUpdateRide(t *testing.T) {
	uow := inmemory.NewFakeUnitOfWork()
	userACL := acl.NewInMemoryUserACL()
	service := service.NewService(uow, userACL)

	userACL.AddUser(&aclModels.User{ID: "1", Name: "John Doe", Email: "john@example.com"})
	b, err := service.CreateBooking("1", "A", "B", 10.0, 100.0)
	if err != nil {
		t.Fatal(err)
	}

	ride := booking.Ride{ID: b.Ride.ID, Source: "A", Destination: "C", Distance: 15.0, Cost: 150.0}

	err = service.UpdateRide(&ride)
	assert.Nil(t, err)

	updatedBooking, err := service.GetBookingByID(b.ID)
	assert.Nil(t, err)
	assert.Equal(t, ride.Source, updatedBooking.Ride.Source)
	assert.Equal(t, ride.Destination, updatedBooking.Ride.Destination)
	assert.Equal(t, ride.Distance, updatedBooking.Ride.Distance)
	assert.Equal(t, ride.Cost, updatedBooking.Ride.Cost)
}
