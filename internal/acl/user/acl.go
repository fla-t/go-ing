package acl

// User represents the User data structure.
type User struct {
	ID    string
	Name  string
	Email string
}

// UserACL defines the contract for fetching user data.
type UserACLInterface interface {
	GetUserByID(id string) (*User, error)
}
