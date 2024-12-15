package booking

import (
	booking "github.com/fla-t/go-ing/internal/domain/booking"
	"github.com/fla-t/go-ing/internal/uow"
)

// Service holds all the booking public methods
type Service struct {
	uow uow.UnitOfWorkInterface
}

// NewService creates a new booking service
func NewService(uow uow.UnitOfWorkInterface) *Service {
	return &Service{uow: uow}
}

// CreateBooking creates a new booking
func (s *Service) CreateBooking(userID string, source string, destination string, distance float64, cost float64) (*booking.Booking, error) {
	ride := booking.NewRide(source, destination, distance, cost)
	booking := booking.NewBooking(userID, *ride)

	if err := s.uow.Begin(); err != nil {
		return nil, err
	}

	if err := s.uow.BookingRepository().CreateBooking(booking); err != nil {
		s.uow.Rollback()
		return nil, err
	}

	s.uow.Commit()
	return booking, nil
}

// GetBookingByID returns a booking by its id
func (s *Service) GetBookingByID(id string) (*booking.Booking, error) {
	return s.uow.BookingRepository().GetBookingByID(id)
}

// UpdateRide updates a ride
func (s *Service) UpdateRide(ride *booking.Ride) error {
	if err := s.uow.Begin(); err != nil {
		return err
	}

	if err := s.uow.BookingRepository().UpdateRide(ride); err != nil {
		s.uow.Rollback()
		return err
	}

	return s.uow.Commit()
}
