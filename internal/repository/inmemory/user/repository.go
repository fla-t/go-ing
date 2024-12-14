package repository

import (
	"errors"

	user "github.com/fla-t/go-ing/internal/domain/user"
)

// Repository is a simple in-memory repository for users for testing
type Repository struct {
	users map[string]*user.User
}

// NewInMemoryUserRepository creates a new UserRepository
func NewInMemoryUserRepository() *Repository {
	return &Repository{
		users: make(map[string]*user.User),
	}
}

// GetByID returns a user by its id
func (r *Repository) GetByID(id string) (*user.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("User Not Found")
	}
	return u, nil
}

// Save saves a user
func (r *Repository) Save(u *user.User) error {
	r.users[u.ID] = u
	return nil
}

// Delete deletes a user
func (r *Repository) Delete(id string) error {
	delete(r.users, id)
	return nil
}
