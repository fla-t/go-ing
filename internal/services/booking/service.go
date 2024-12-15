package booking

import (
	acl "github.com/fla-t/go-ing/internal/acl/user"
	domain "github.com/fla-t/go-ing/internal/domain/booking"
	"github.com/fla-t/go-ing/internal/uow"
)

// Service holds all the booking public methods
type Service struct {
	uow     uow.UnitOfWorkInterface
	userACL acl.UserACLInterface
}

// NewService creates a new booking service
func NewService(uow uow.UnitOfWorkInterface, userACL acl.UserACLInterface) *Service {
	return &Service{
		uow:     uow,
		userACL: userACL,
	}
}

// CreateBooking creates a new booking
func (s *Service) CreateBooking(userID string, source string, destination string, distance float64, cost float64) (*domain.Booking, error) {
	ride := domain.NewRide(source, destination, distance, cost)
	booking := domain.NewBooking(userID, *ride)

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
func (s *Service) GetBookingByID(id string) (*Booking, error) {
	b, err := s.uow.BookingRepository().GetBookingByID(id)
	if err != nil {
		return nil, err
	}

	user, err := s.userACL.GetUserByID(b.UserID)
	if err != nil {
		return nil, err
	}

	return &Booking{
		ID:   b.ID,
		Name: user.Name,
		Ride: &Ride{
			Source:      b.Ride.Source,
			Destination: b.Ride.Destination,
			Distance:    b.Ride.Distance,
			Cost:        b.Ride.Cost,
		},
		Time: b.Time,
	}, nil

}

// UpdateRide updates a ride
func (s *Service) UpdateRide(ride *domain.Ride) error {
	if err := s.uow.Begin(); err != nil {
		return err
	}

	if err := s.uow.BookingRepository().UpdateRide(ride); err != nil {
		s.uow.Rollback()
		return err
	}

	return s.uow.Commit()
}
