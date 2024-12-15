package acl

import (
	"context"
	"fmt"
	"time"

	acl "github.com/fla-t/go-ing/internal/acl/user"
	pb "github.com/fla-t/go-ing/proto/user" // Import the generated gRPC code
	"google.golang.org/grpc"
)

// GRPCUserACL is a gRPC implementation of UserACL.
type GRPCUserACL struct {
	client pb.UserServiceClient
}

// NewGRPCUserACL creates a new GRPCUserACL.
func NewGRPCUserACL(conn *grpc.ClientConn) *GRPCUserACL {
	return &GRPCUserACL{
		client: pb.NewUserServiceClient(conn),
	}
}

// GetUserByID fetches user data from the gRPC User Service.
func (a *GRPCUserACL) GetUserByID(id string) (*acl.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Make the gRPC call
	resp, err := a.client.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("gRPC error: %w", err)
	}

	// Map the gRPC response to the internal User structure
	return &acl.User{
		ID:    resp.User.GetId(),
		Name:  resp.User.GetName(),
		Email: resp.User.GetEmail(),
	}, nil
}
