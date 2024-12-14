package user

import (
	"errors"

	user "github.com/fla-t/go-ing/internal/domain/user"
	"github.com/fla-t/go-ing/internal/uow"
)

// Service holds all the user public methods
type Service struct {
	uow uow.UnitOfWorkInterface
}

// NewService creates a new user service
func NewService(uow uow.UnitOfWorkInterface) *Service {
	return &Service{uow: uow}
}

// CreateUser creates a new user
func (s *Service) CreateUser(u *user.User) error {
	if u.ID == "" || u.Name == "" || u.Email == "" {
		return errors.New("invalid user data")
	}

	if err := s.uow.Begin(); err != nil {
		return err
	}

	if err := s.uow.UserRepository().Save(u); err != nil {
		s.uow.Rollback()
		return err
	}

	return s.uow.Commit()
}

// GetUserByID returns a user by its id
func (s *Service) GetUserByID(id string) (*user.User, error) {
	return s.uow.UserRepository().GetByID(id)
}

// DeleteUser deletes a user by its id
func (s *Service) DeleteUser(id string) error {
	if err := s.uow.Begin(); err != nil {
		return err
	}

	if err := s.uow.UserRepository().Delete(id); err != nil {
		s.uow.Rollback()
	}

	return s.uow.Commit()
}
