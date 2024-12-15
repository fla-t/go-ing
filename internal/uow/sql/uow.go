package sql

import (
	"database/sql"

	booking "github.com/fla-t/go-ing/internal/domain/booking"
	user "github.com/fla-t/go-ing/internal/domain/user"
	bookingRepo "github.com/fla-t/go-ing/internal/repository/sql/booking"
	userRepo "github.com/fla-t/go-ing/internal/repository/sql/user"
	"github.com/fla-t/go-ing/internal/uow"
)

// DbUnitOfWork is a struct that holds all the repositories
type DbUnitOfWork struct {
	db          *sql.DB
	tx          *sql.Tx
	userRepo    user.RepositoryInterface
	bookingRepo booking.RepositoryInterface
}

// NewDbUnitOfWork starts a transaction
func NewDbUnitOfWork(db *sql.DB) uow.UnitOfWorkInterface {
	return &DbUnitOfWork{db: db}
}

// Begin starts a transaction
func (uow *DbUnitOfWork) Begin() error {
	tx, err := uow.db.Begin()
	if err != nil {
		return err
	}

	uow.tx = tx
	uow.userRepo = userRepo.NewUserRepository(uow.tx)
	uow.bookingRepo = bookingRepo.NewBookingRepository(uow.tx)

	return nil
}

// Commit commits the transaction
func (uow *DbUnitOfWork) Commit() error {
	return uow.tx.Commit()
}

// Rollback rolls back the transaction
func (uow *DbUnitOfWork) Rollback() error {
	return uow.tx.Rollback()
}

// UserRepository returns the user repository
func (uow *DbUnitOfWork) UserRepository() user.RepositoryInterface {
	return uow.userRepo
}

// BookingRepository returns the user repository
func (uow *DbUnitOfWork) BookingRepository() booking.RepositoryInterface {
	return uow.bookingRepo
}
