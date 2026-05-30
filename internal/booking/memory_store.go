package booking

type MemoryStore struct {
	//seats to booking
	bookings map[string]Booking
}

// constructor function
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	if _, exists := s.bookings[b.SeatId]; exists {
		return ErrorSeatAlreadyBooked
	}

	s.bookings[b.SeatId] = b

	return nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	var bookings []Booking

	for _, v := range s.bookings {
		if v.MovieId == movieID {
			bookings = append(bookings, v)
		}
	}

	return bookings
}
