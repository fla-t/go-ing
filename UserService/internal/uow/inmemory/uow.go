package uow

import (
	user "github.com/fla-t/go-ing/UserService/internal/domain/user"
	inmemory "github.com/fla-t/go-ing/UserService/internal/repository/inmemory/user"
)

// FakeUnitOfWork is a struct that holds all the repositories
type FakeUnitOfWork struct {
	userRepository user.RepositoryInterface
}

// NewFakeUnitOfWork creates a new FakeUnitOfWork
func NewFakeUnitOfWork() *FakeUnitOfWork {
	return &FakeUnitOfWork{
		userRepository: inmemory.NewInMemoryUserRepository(),
	}
}

// Begin simulates a begin of a transaction
func (u *FakeUnitOfWork) Begin() error { return nil }

// Commit simulates a commit of a transaction
func (u *FakeUnitOfWork) Commit() error { return nil }

// Rollback simulates a rollback of a transaction
func (u *FakeUnitOfWork) Rollback() error { return nil }

// UserRepository returns the user repository
func (u *FakeUnitOfWork) UserRepository() user.RepositoryInterface {
	return u.userRepository
}
