package booking

import (
	"context"
	"net"
	"testing"

	aclModels "github.com/fla-t/go-ing/internal/acl/user"
	acl "github.com/fla-t/go-ing/internal/acl/user/inmemory"
	bookingGRPC "github.com/fla-t/go-ing/internal/grpc/booking"
	service "github.com/fla-t/go-ing/internal/services/booking"
	uow "github.com/fla-t/go-ing/internal/uow/inmemory"
	proto "github.com/fla-t/go-ing/proto/booking"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	uow := uow.NewFakeUnitOfWork()
	userACL := acl.NewInMemoryUserACL()
	service := service.NewService(uow, userACL)

	userACL.AddUser(&aclModels.User{ID: "1", Name: "Test User", Email: "test@example.com"})
	proto.RegisterBookingServiceServer(s, bookingGRPC.NewBookingService(service))
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateBooking(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewBookingServiceClient(conn)

	req := &proto.CreateBookingRequest{
		UserId: "1",
		Ride: &proto.Ride{
			Source:      "A",
			Destination: "B",
			Distance:    10.0,
			Cost:        100.0,
		},
	}

	resp, err := client.CreateBooking(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.UserId, resp.Booking.UserId)
}

func TestGetBooking(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewBookingServiceClient(conn)

	// Create a booking first
	createReq := &proto.CreateBookingRequest{
		UserId: "1",
		Ride: &proto.Ride{
			Source:      "A",
			Destination: "B",
			Distance:    10.0,
			Cost:        100.0,
		},
	}
	createResp, err := client.CreateBooking(ctx, createReq)
	assert.Nil(t, err)
	assert.NotNil(t, createResp)

	req := &proto.GetBookingRequest{
		Id: createResp.Booking.Id,
	}

	resp, err := client.GetBooking(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestUpdateRide(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewBookingServiceClient(conn)

	// Create a booking first
	createReq := &proto.CreateBookingRequest{
		UserId: "1",
		Ride: &proto.Ride{
			Source:      "A",
			Destination: "B",
			Distance:    10.0,
			Cost:        100.0,
		},
	}
	createResp, err := client.CreateBooking(ctx, createReq)
	assert.Nil(t, err)
	assert.NotNil(t, createResp)

	req := &proto.UpdateRideRequest{
		RideId: createResp.Booking.RideId,
		Ride: &proto.Ride{
			Source:      "A",
			Destination: "B",
			Distance:    15.0,
			Cost:        150.0,
		},
	}

	resp, err := client.UpdateRide(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Message, "Ride updated successfully")
}
