package redis

import (
	"strings"
	"time"

	"bws/pkg/trace"

	"github.com/apex/log"
	"github.com/cenkalti/backoff/v4"
	"github.com/gomodule/redigo/redis"
)

type RedisClient interface {
	Ping() (err error)
	Keys(pattern string) (values []string, err error)
	AllKeys() (values []string, err error)
	Get(collection string) (value string, err error)
	Set(collection string, value string) (err error)
	Del(collection string) (err error)
	RPush(collection string, value interface{}) (err error)
	LPush(collection string, value interface{}) (err error)
	LPop(collection string) (value string, err error)
	RPop(collection string) (value string, err error)
	LLen(collection string) (count int, err error)
	BRPop(collection string, timeout int) (value string, err error)
	BLPop(collection string, timeout int) (value string, err error)
	HSet(collection string, key string, value string) (err error)
	HGet(collection string, key string) (value string, err error)
	HDel(collection string, key string) (err error)
	HScan(collection string) (results map[string]string, err error)
	HKeys(collection string) (results []string, err error)
	ZAdd(collection string, score float32, value interface{}) (err error)
	ZCount(collection string, min string, max string) (count int, err error)
	ZCountAll(collection string) (count int, err error)
	ZScan(collection string, pattern string, count int) (results []string, err error)
	ZPopMax(collection string, count int) (results []string, err error)
	ZPopMin(collection string, count int) (results []string, err error)
	ZPopMaxOne(collection string) (value string, err error)
	ZPopMinOne(collection string) (value string, err error)
	BZPopMax(collection string, timeout int) (value string, err error)
	BZPopMin(collection string, timeout int) (value string, err error)
	Lock(lockKey string) (value int64, err error)
	UnLock(lockKey string, value int64)
	MemoryStats() (stats map[string]int64, err error)
	SetBackoffMaxInterval(interval time.Duration)
	SetTimeout(timeout int)
}

type Client struct {
	backoffMaxInterval time.Duration
	timeout            int

	pool *redis.Pool
}

var client RedisClient

func NewRedisClient(opts ...Option) (client *Client, err error) {
	client = &Client{
		backoffMaxInterval: 20 * time.Second,
		pool:               NewRedisPool(),
	}

	for _, opt := range opts {
		opt(client)
	}

	if err := client.init(); err != nil {
		return nil, err
	}
	return client, nil
}

func GetRedisClient() (c RedisClient, err error) {
	if client != nil {
		return client, nil
	}
	c, err = NewRedisClient()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) init() (err error) {
	b := backoff.NewExponentialBackOff()
	b.MaxInterval = c.backoffMaxInterval
	if err := backoff.Retry(func() error {
		err := c.Ping()
		if err != nil {
			log.WithError(err).Warnf("waiting for redis pool active connection. will after %f seconds try again.", b.NextBackOff().Seconds())
		}
		return nil
	}, b); err != nil {
		return trace.TraceError(err)
	}
	return nil
}

func (c *Client) getLockKey(lockKey string) string {
	lockKey = strings.ReplaceAll(lockKey, ":", "-")
	return "nodes:lock:" + lockKey
}
