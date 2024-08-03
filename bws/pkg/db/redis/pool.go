package redis

import (
	"time"

	"bws/pkg/trace"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

func NewRedisPool() *redis.Pool {
	address := viper.GetString("redis.address")
	port := viper.GetString("redis.port")
	database := viper.GetString("redis.database")
	password := viper.GetString("redis.password")

	// normalize params
	if address == "" {
		address = DEFAULT_ADDRESS
	}
	if port == "" {
		port = DEFAULT_PORT
	}
	if database == "" {
		database = DEFAULT_DATABASE
	}

	var url string
	if password == "" {
		url = "redis://" + address + ":" + port + "/" + database
	} else {
		url = "redis://x:" + password + "@" + address + ":" + port + "/" + database
	}
	return &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.DialURL(
				url,
				redis.DialConnectTimeout(time.Second*10),
				redis.DialReadTimeout(time.Second*600),
				redis.DialWriteTimeout(time.Second*10),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return trace.TraceError(err)
		},
		MaxIdle:         10,
		MaxActive:       0,
		IdleTimeout:     300 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
