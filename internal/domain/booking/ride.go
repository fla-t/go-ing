package booking

import "github.com/google/uuid"

// Ride represents a ride in the system.
type Ride struct {
	ID          string
	source      string
	destination string
	distance    float64 // in km, would be better if "km" was in the name
	cost        float64
}

// NewRide creates a new ride.
func NewRide(source, destination string, distance, cost float64) *Ride {
	return &Ride{
		ID:          uuid.New().String(),
		source:      source,
		destination: destination,
		distance:    distance,
		cost:        cost,
	}
}
