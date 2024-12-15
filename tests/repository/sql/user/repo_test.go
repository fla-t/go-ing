package user

import (
	"testing"

	"github.com/fla-t/go-ing/internal/domain/user"
	"github.com/fla-t/go-ing/internal/uow/sql"
	"github.com/fla-t/go-ing/tests"
	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {
	// Setup the test database
	db, cleanup, err := tests.SetupTestDatabase()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	defer cleanup()

	// Initialize db unit of work
	uow := sql.NewDbUnitOfWork(db)
	uow.Begin()

	// Test data
	newUser := user.User{
		ID:    uuid.New().String(),
		Name:  "testuser",
		Email: "testuser@example.com",
	}

	// Perform the CreateUser operation
	err = uow.UserRepository().Save(&newUser)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Commit the transaction
	err = uow.Commit()
	if err != nil {
		t.Fatalf("failed to commit transaction: %v", err)
	}

	// Verify the record in the database
	var count int
	err = db.QueryRow("select count(*) from users where id = $1", newUser.ID).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query database: %v", err)
	}

	if count != 1 {
		t.Errorf("expected 1 user, got %d", count)
	}
}

func TestGetUserByID(t *testing.T) {
	// Setup the test database
	db, cleanup, err := tests.SetupTestDatabase()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	defer cleanup()

	// Initialize db unit of work
	uow := sql.NewDbUnitOfWork(db)
	uow.Begin()

	// Test data
	newUser := user.User{
		ID:    uuid.New().String(),
		Name:  "testuser",
		Email: "testuser@example.com",
	}

	// Perform the CreateUser operation
	err = uow.UserRepository().Save(&newUser)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Perform the GetUserByID operation
	u, err := uow.UserRepository().GetByID(newUser.ID)
	if err != nil {
		t.Fatalf("failed to get user by id: %v", err)
	}

	// Commit the transaction
	err = uow.Commit()
	if err != nil {
		t.Fatalf("failed to commit transaction: %v", err)
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

func TestSaveUser(t *testing.T) {
	// Setup the test database
	db, cleanup, err := tests.SetupTestDatabase()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	defer cleanup()

	// Initialize db unit of work
	uow := sql.NewDbUnitOfWork(db)
	uow.Begin()

	// Test data
	testUser := user.User{
		ID:    uuid.New().String(),
		Name:  "Test User",
		Email: "testuser@example.com",
	}

	// Perform the Save operation
	err = uow.UserRepository().Save(&testUser)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	// Commit the transaction
	err = uow.Commit()
	if err != nil {
		t.Fatalf("failed to commit transaction: %v", err)
	}

	// Verify the record in the database
	var count int
	err = db.QueryRow("select count(*) from users where id = $1", testUser.ID).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query database: %v", err)
	}

	if count != 1 {
		t.Errorf("expected 1 user, got %d", count)
	}
}

func TestGetByID(t *testing.T) {
	// Setup the test database
	db, cleanup, err := tests.SetupTestDatabase()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	defer cleanup()

	// Initialize db unit of work
	uow := sql.NewDbUnitOfWork(db)
	uow.Begin()

	// Test data
	testUser := user.User{
		ID:    uuid.New().String(),
		Name:  "Test User",
		Email: "testuser@example.com",
	}

	// Save the test user
	err = uow.UserRepository().Save(&testUser)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	// Perform the GetByID operation
	retrievedUser, err := uow.UserRepository().GetByID(testUser.ID)
	if err != nil {
		t.Fatalf("failed to get user by id: %v", err)
	}

	// Commit the transaction
	err = uow.Commit()
	if err != nil {
		t.Fatalf("failed to commit transaction: %v", err)
	}

	// Verify the returned user
	if retrievedUser.ID != testUser.ID {
		t.Errorf("expected user id %s, got %s", testUser.ID, retrievedUser.ID)
	}
	if retrievedUser.Name != testUser.Name {
		t.Errorf("expected name %s, got %s", testUser.Name, retrievedUser.Name)
	}
	if retrievedUser.Email != testUser.Email {
		t.Errorf("expected email %s, got %s", testUser.Email, retrievedUser.Email)
	}
}
