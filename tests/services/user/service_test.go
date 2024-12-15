package user_test

import (
	"testing"

	service "github.com/fla-t/go-ing/internal/services/user"
	"github.com/fla-t/go-ing/internal/uow/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	uow := inmemory.NewFakeUnitOfWork()
	service := service.NewService(uow)

	id, err := service.CreateUser("John Doe", "john@example.com")
	assert.Nil(t, err)

	savedUser, err := service.GetUserByID(id)
	assert.Nil(t, err)
	assert.Equal(t, savedUser.ID, id)
}

func TestDeleteUser(t *testing.T) {
	uow := inmemory.NewFakeUnitOfWork()
	service := service.NewService(uow)

	id, err := service.CreateUser("John Doe", "john@example.com")
	assert.Nil(t, err)

	err = service.DeleteUser(id)
	assert.Nil(t, err)

	_, err = service.GetUserByID(id)
	assert.NotNil(t, err)
}
