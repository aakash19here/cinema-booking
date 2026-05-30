package redis

import (
	"context"
	"log"

	goredis "github.com/redis/go-redis/v9"
)

func NewClient(addr string) *goredis.Client {
	rdb := goredis.NewClient(&goredis.Options{Addr: addr})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis ping: %v", err)
	}

	log.Println("connected to redis at", addr)

	return rdb
}
