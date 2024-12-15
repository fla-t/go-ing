package booking

import "time"

// Ride here is the service model for ride
type Ride struct {
	Source      string
	Destination string
	Distance    float64
	Cost        float64
}

// Booking here is the service model for booking
type Booking struct {
	ID   string
	Name string // This should be user's name, inject acl here
	Ride *Ride
	Time time.Time // time of booking in UTC
}
