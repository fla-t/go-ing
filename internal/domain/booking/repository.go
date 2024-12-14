package booking

// RepositoryInterface for the Booking aggregate
type RepositoryInterface interface {
	CreateBooking(booking *Booking) error
	GetBookingByID(id string) (*Booking, error)
	UpdateRide(ride *Ride) error
}
