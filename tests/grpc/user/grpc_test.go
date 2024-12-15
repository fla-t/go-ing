package user

import (
	"context"
	"net"
	"testing"

	userGRPC "github.com/fla-t/go-ing/internal/grpc/user"
	service "github.com/fla-t/go-ing/internal/services/user"
	uow "github.com/fla-t/go-ing/internal/uow/inmemory"
	proto "github.com/fla-t/go-ing/proto/user"
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
	service := service.NewService(uow)
	proto.RegisterUserServiceServer(s, userGRPC.NewUserService(service))
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewUserServiceClient(conn)

	req := &proto.CreateUserRequest{
		Id:    "1",
		Name:  "Test User",
		Email: "testuser@example.com",
	}

	resp, err := client.CreateUser(ctx, req)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	assert.Equal(t, resp.Message, "User created successfully")
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewUserServiceClient(conn)

	req := &proto.GetUserRequest{
		Id: "1",
	}

	resp, err := client.GetUser(ctx, req)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	assert.Equal(t, req.Id, resp.Id)
	assert.Equal(t, "Test User", resp.Name)
	assert.Equal(t, "testuser@example.com", resp.Email)
}

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewUserServiceClient(conn)

	req := &proto.DeleteUserRequest{
		Id: "1",
	}

	resp, err := client.DeleteUser(ctx, req)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}

	assert.Equal(t, req.Id, resp.Id)
}
