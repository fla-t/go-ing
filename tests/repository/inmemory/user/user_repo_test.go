package user

import (
	"testing"

	"github.com/fla-t/go-ing/internal/domain/user"
	repository "github.com/fla-t/go-ing/internal/repository/inmemory/user"
	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {
	// Initialize in-memory repository
	repo := repository.NewInMemoryUserRepository()

	// Test data
	newUser := user.User{
		ID:    uuid.New().String(),
		Name:  "testuser",
		Email: "testuser@example.com",
	}

	// Perform the Save operation
	err := repo.Save(&newUser)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Verify the record in the repository
	u, err := repo.GetByID(newUser.ID)
	if err != nil {
		t.Fatalf("failed to get user by id: %v", err)
	}

	if u.ID != newUser.ID {
		t.Errorf("expected user id %s, got %s", newUser.ID, u.ID)
	}
	if u.Name != newUser.Name {
		t.Errorf("expected name %s, got %s", newUser.Name, u.Name)
	}
	if u.Email != newUser.Email {
		t.Errorf("expected email %s, got %s", newUser.Email, u.Email)
	}
}

func TestGetUserByID(t *testing.T) {
	// Initialize in-memory repository
	repo := repository.NewInMemoryUserRepository()

	// Test data
	newUser := user.User{
		ID:    uuid.New().String(),
		Name:  "testuser",
		Email: "testuser@example.com",
	}

	// Perform the Save operation
	err := repo.Save(&newUser)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Perform the GetByID operation
	u, err := repo.GetByID(newUser.ID)
	if err != nil {
		t.Fatalf("failed to get user by id: %v", err)
	}

	// Verify the returned user
	if u.ID != newUser.ID {
		t.Errorf("expected user id %s, got %s", newUser.ID, u.ID)
	}
	if u.Name != newUser.Name {
		t.Errorf("expected name %s, got %s", newUser.Name, u.Name)
	}
	if u.Email != newUser.Email {
		t.Errorf("expected email %s, got %s", newUser.Email, u.Email)
	}
}
