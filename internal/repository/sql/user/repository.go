package sql

import (
	"database/sql"

	user "github.com/fla-t/go-ing/internal/domain/user"
)

// Repository is a struct that holds the database connection
type Repository struct {
	tx *sql.Tx
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(tx *sql.Tx) *Repository {
	return &Repository{
		tx: tx,
	}
}

// GetByID returns a user by its id
func (r *Repository) GetByID(id string) (*user.User, error) {
	var u user.User
	query := "select id, name, email from users where id = $1"

	err := r.tx.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Save saves a user
func (r *Repository) Save(u *user.User) error {
	query := "insert into users (id, name, email) values ($1, $2, $3)"
	_, err := r.tx.Exec(query, u.ID, u.Name, u.Email)

	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a user
func (r *Repository) Delete(id string) error {
	query := "delete from users where id = $1"
	_, err := r.tx.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
