package booking

import (
	"context"

	domain "github.com/fla-t/go-ing/internal/domain/booking"
	"github.com/fla-t/go-ing/internal/services/booking"
	proto "github.com/fla-t/go-ing/proto/booking"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Service is the gRPC server for the booking service
type Service struct {
	proto.UnimplementedBookingServiceServer
	service *booking.Service
}

// NewBookingService creates a new BookingService
func NewBookingService(service *booking.Service) *Service {
	return &Service{service: service}
}

// CreateBooking creates a new booking
func (s *Service) CreateBooking(ctx context.Context, req *proto.CreateBookingRequest) (*proto.CreateBookingResponse, error) {
	b, err := s.service.CreateBooking(req.UserId, req.Ride.GetSource(), req.Ride.GetDestination(), req.Ride.GetDistance(), req.Ride.GetCost())

	if err != nil {
		return nil, err
	}

	return &proto.CreateBookingResponse{
		Booking: &proto.Booking{
			Id:     b.ID,
			UserId: b.UserID,
			RideId: b.Ride.ID,
			Time:   timestamppb.New(b.Time),
		}}, nil
}

// GetBooking returns a booking by its id
func (s *Service) GetBooking(ctx context.Context, req *proto.GetBookingRequest) (*proto.GetBookingResponse, error) {
	b, err := s.service.GetBookingByID(req.GetBookingId())

	if err != nil {
		return nil, err
	}

	return &proto.GetBookingResponse{
		Name: b.Name, // This should be user's name, inject acl here
		Ride: &proto.Ride{
			Source:      b.Ride.Source,
			Destination: b.Ride.Destination,
			Distance:    b.Ride.Distance,
			Cost:        b.Ride.Cost,
		},
	}, nil
}

// UpdateRide updates a ride
func (s *Service) UpdateRide(ctx context.Context, req *proto.UpdateRideRequest) (*proto.UpdateRideResponse, error) {
	r := &domain.Ride{
		ID:          req.GetRideId(),
		Source:      req.GetRide().GetSource(),
		Destination: req.GetRide().GetDestination(),
		Distance:    req.GetRide().GetDistance(),
		Cost:        req.GetRide().GetCost(),
	}

	if err := s.service.UpdateRide(r); err != nil {
		return nil, err
	}

	return &proto.UpdateRideResponse{Message: "Ride updated successfully"}, nil
}
