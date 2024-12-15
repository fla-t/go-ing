package tests

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// SetupTestDatabase sets up container for PostgreSQL
func SetupTestDatabase() (*sql.DB, func(), error) {
	ctx := context.Background()

	// Start PostgreSQL container
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15"),
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpassword"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start container: %w", err)
	}

	// Get the connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get connection string: %w", err)
	}

	// Retry connecting until the database is ready
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Create schema
	_, err = db.Exec(`
		create table rides (
			id uuid primary key,
			source text not null,
			destination text not null,
			distance double precision not null,
			cost double precision not null
		);
		
		create table bookings (
			id uuid primary key,
			user_id text not null,
			ride_id uuid references rides (id) not null,
			time timestamptz not null
		);
			
		create table users (
			id uuid primary key,
			name text not null,
			email text not null
		);
	`)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create schema: %w", err)
	}

	// Return the database and cleanup function
	cleanup := func() {
		_ = db.Close()
		_ = pgContainer.Terminate(ctx)
	}
	return db, cleanup, nil
}
