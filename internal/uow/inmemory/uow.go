package uow

import (
	booking "github.com/fla-t/go-ing/internal/domain/booking"
	user "github.com/fla-t/go-ing/internal/domain/user"
	bookingRepo "github.com/fla-t/go-ing/internal/repository/inmemory/booking"
	userRepo "github.com/fla-t/go-ing/internal/repository/inmemory/user"
)

// FakeUnitOfWork is a struct that holds all the repositories
type FakeUnitOfWork struct {
	userRepository    user.RepositoryInterface
	bookingRepository booking.RepositoryInterface
}

// NewFakeUnitOfWork creates a new FakeUnitOfWork
func NewFakeUnitOfWork() *FakeUnitOfWork {
	return &FakeUnitOfWork{
		userRepository:    userRepo.NewInMemoryUserRepository(),
		bookingRepository: bookingRepo.NewInMemoryBookingRepository(),
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

// BookingRepository returns the user repository
func (u *FakeUnitOfWork) BookingRepository() booking.RepositoryInterface {
	return u.bookingRepository
}
