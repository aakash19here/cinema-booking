package booking

// Store is implemented by the booking backends (in-memory, concurrent, ...).
type Store interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) Book(b Booking) error {
	return s.store.Book(b)
}

func (s *Service) ListBookings(movieID string) []Booking {
	return s.store.ListBookings(movieID)
}
