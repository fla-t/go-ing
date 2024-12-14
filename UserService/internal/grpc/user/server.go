package user

import (
	"context"

	domain "github.com/fla-t/go-ing/UserService/internal/domain/user"
	"github.com/fla-t/go-ing/UserService/internal/services/user"
	"github.com/fla-t/go-ing/UserService/proto"
)

// UserService is the gRPC server for the user service
type UserService struct {
	proto.UnimplementedUserServiceServer
	service *user.Service
}

// NewUserService creates a new UserService
func NewUserService(service *user.Service) *UserService {
	return &UserService{service: service}
}

// GetUser returns a user by its id
func (s *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	u, err := s.service.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		User: &proto.User{
			Id:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		},
	}, nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	u := &domain.User{
		ID:    req.GetId(),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	if err := s.service.CreateUser(u); err != nil {
		return nil, err
	}

	return &proto.CreateUserResponse{Message: "User created successfully"}, nil
}

// DeleteUser deletes a user by its id
func (s *UserService) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := s.service.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.DeleteUserResponse{Message: "User deleted successfully"}, nil
}
