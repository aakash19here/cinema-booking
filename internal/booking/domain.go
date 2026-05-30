package booking

import "errors"

var (
	ErrorSeatAlreadyBooked = errors.New("seat already taken")
)

type Booking struct {
	Id      string
	MovieId string
	SeatId  string
	UserId  string
	Status  string
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}
