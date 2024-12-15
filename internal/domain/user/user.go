package user

import "github.com/google/uuid"

// User represents a user of the system.
type User struct {
	ID    string
	Name  string
	Email string
}

// NewUser creates a new user.
func NewUser(name, email string) *User {
	return &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
}
