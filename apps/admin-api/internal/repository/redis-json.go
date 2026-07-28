package repository

import (
	"time"
)

func (h Handler) JSONSet(key string, value interface{}, expiration time.Duration) error {
	return h.redis.JSONSet(key, value, expiration)
}

func (h Handler) JSONGet(key string, result interface{}) error {
	return h.redis.JSONGet(key, result)
}
