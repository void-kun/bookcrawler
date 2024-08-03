package redis

import (
	"time"
)

type Option func(c RedisClient)

func WithBackoffMaxInterval(interval time.Duration) Option {
	return func(c RedisClient) {
		c.SetBackoffMaxInterval(interval)
	}
}

func WithTimeout(timeout int) Option {
	return func(c RedisClient) {
		c.SetTimeout(timeout)
	}
}
