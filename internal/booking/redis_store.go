package booking

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const defaultHoldTTL = 2 * time.Minute

type RedisStore struct {
	rdb *redis.Client
}

func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{
		rdb,
	}
}

func sessionKey(id string) string {
	return fmt.Sprintf("session:%s", id)
}

func holdKey(movieID, seatID string) string {
	return fmt.Sprintf("seat:%s:%s", movieID, seatID)
}

func (s *RedisStore) Book(b Booking) error {

	session, err := s.hold(b)

	if err != nil {
		return err
	}

	log.Printf("session %v \n", session)

	return nil
}

func (s *RedisStore) ListBookings(movieID string) []Booking {
	return []Booking{}
}

func (s *RedisStore) hold(b Booking) (Booking, error) {
	id := uuid.New().String()
	now := time.Now()

	ctx := context.Background()

	booking := Booking{Id: id,
		MovieId:   b.MovieId,
		SeatId:    b.SeatId,
		UserId:    b.UserId,
		Status:    "held",
		ExpiresAt: now.Add(defaultHoldTTL)}

	data, _ := json.Marshal(booking)

	cmd := s.rdb.SetArgs(ctx, holdKey(b.MovieId, b.SeatId), data, redis.SetArgs{
		Mode: "NX",
		TTL:  defaultHoldTTL,
	})

	if cmd.Val() != "OK" {
		return Booking{}, ErrorSeatAlreadyBooked
	}

	cmd = s.rdb.Set(ctx, sessionKey(id), holdKey(b.MovieId, b.SeatId), defaultHoldTTL)

	if cmd.Val() != "OK" {
		return Booking{}, errors.New("something went wrong")
	}

	return booking, nil
}
