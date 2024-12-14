package domain

// RepositoryInterface for the User entity.
type RepositoryInterface interface {
	GetByID(id string) (*User, error)
	Save(user *User) error
	Delete(id string) error
}
