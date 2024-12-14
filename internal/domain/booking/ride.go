package booking

import "github.com/google/uuid"

// Ride represents a ride in the system.
type Ride struct {
	ID          string
	Source      string
	Destination string
	Distance    float64 // in km, would be better if "km" was in the name
	Cost        float64
}

// NewRide creates a new ride.
func NewRide(source, destination string, distance, cost float64) *Ride {
	return &Ride{
		ID:          uuid.New().String(),
		Source:      source,
		Destination: destination,
		Distance:    distance,
		Cost:        cost,
	}
}
