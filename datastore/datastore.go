package datastore

import (
	"github.com/go-redis/redis"
	"time"
)
type store struct {
	client *redis.Client
}

func New() *store {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	return &store{client: redisClient}
}

func (s *store) IsEntryPresent(key string) bool {
	_, err := s.client.Get(key).Result()
	if err != nil {
		return false
	}
	return true
}

func (s *store) UpsertEntry(key string, ttl_in_seconds int) {
	if s.IsEntryPresent(key) {
		pipeline := s.client.Pipeline()
		pipeline.Get(key)
		pipeline.Expire(key, time.Duration(ttl_in_seconds)*time.Second)
		pipeline.Exec()
	} else {
		s.client.Set(key, "", time.Duration(ttl_in_seconds)*time.Second)
	}

}
