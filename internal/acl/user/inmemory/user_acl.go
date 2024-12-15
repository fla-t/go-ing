package acl

import (
	"errors"

	acl "github.com/fla-t/go-ing/internal/acl/user"
)

// InMemoryUserACL is an in-memory implementation of UserACL for testing.
type InMemoryUserACL struct {
	users map[string]*acl.User
}

// NewInMemoryUserACL creates a new InMemoryUserACL.
func NewInMemoryUserACL() *InMemoryUserACL {
	return &InMemoryUserACL{
		users: make(map[string]*acl.User),
	}
}

// AddUser adds a user to the in-memory ACL.
func (a *InMemoryUserACL) AddUser(user *acl.User) {
	a.users[user.ID] = user
}

// GetUserByID fetches a user by ID from the in-memory map.
func (a *InMemoryUserACL) GetUserByID(id string) (*acl.User, error) {
	user, exists := a.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
