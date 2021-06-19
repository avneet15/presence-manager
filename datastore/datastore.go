package datastore

import (
	"github.com/go-redis/redis"
)
type store struct {
	client redis.client
}

func New() *store {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	return &store{client: redisClient}
}

func (s *store) isEntryPresent(key string) bool {
	_, err := s.client.Get(key).Result()
	if err != nil {
		return false
	}
	return true
}
