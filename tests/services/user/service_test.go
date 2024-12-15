package user_test

import (
	"testing"

	"github.com/fla-t/go-ing/internal/domain/user"
	service "github.com/fla-t/go-ing/internal/services/user"
	"github.com/fla-t/go-ing/internal/uow/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	uow := inmemory.NewFakeUnitOfWork()
	service := service.NewService(uow)

	u := &user.User{ID: "1", Name: "John Doe", Email: "john@example.com"}
	err := service.CreateUser(u)
	assert.Nil(t, err)

	savedUser, err := service.GetUserByID("1")
	assert.Nil(t, err)
	assert.Equal(t, u, savedUser)
}

func TestDeleteUser(t *testing.T) {
	uow := inmemory.NewFakeUnitOfWork()
	service := service.NewService(uow)

	u := &user.User{ID: "1", Name: "John Doe", Email: "john@example.com"}
	service.CreateUser(u)

	err := service.DeleteUser("1")
	assert.Nil(t, err)

	_, err = service.GetUserByID("1")
	assert.NotNil(t, err)
}
