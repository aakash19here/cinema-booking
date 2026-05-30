package booking

import "sync"

// maps in golang are not concurrent safe
type ConcurrentStore struct {
	//seats to booking
	bookings map[string]Booking
	mu       sync.RWMutex
}

// constructor function
func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (s *ConcurrentStore) Book(b Booking) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.bookings[b.SeatId]; exists {
		return ErrorSeatAlreadyBooked
	}

	s.bookings[b.SeatId] = b

	return nil
}

func (s *ConcurrentStore) ListBookings(movieID string) []Booking {
	var bookings []Booking
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, v := range s.bookings {
		if v.MovieId == movieID {
			bookings = append(bookings, v)
		}
	}

	return bookings
}
