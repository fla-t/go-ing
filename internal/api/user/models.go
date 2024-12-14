package user

import (
	"errors"

	domain "github.com/fla-t/go-ing/internal/domain/user"
)

// User represents a user of the system.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Validate validates the user.
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

// ConvertToDomain converts a user to a domain user.
func (u *User) ConvertToDomain() *domain.User {
	return &domain.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

// ConvertFromDomain converts a domain user to a user.
func ConvertFromDomain(u *domain.User) *User {
	return &User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
