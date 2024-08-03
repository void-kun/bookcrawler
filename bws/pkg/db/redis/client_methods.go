package redis

import (
	"reflect"
	"time"

	"bws/pkg/db/errors"
	"bws/pkg/db/utils"
	"bws/pkg/trace"

	"github.com/apex/log"
	"github.com/gomodule/redigo/redis"
)

func (c *Client) Ping() error {
	client := c.pool.Get()
	defer utils.Close(client)

	if _, err := redis.String(client.Do("PING")); err != nil {
		if err != redis.ErrNil {
			return trace.TraceError(err)
		}
		return err
	}
	return nil
}

func (c *Client) Keys(pattern string) (values []string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	values, err = redis.Strings(client.Do("KEYS", pattern))
	if err != nil {
		return nil, trace.TraceError(err)
	}
	return values, nil
}

func (c *Client) AllKeys() (values []string, err error) {
	return c.Keys("*")
}

func (c *Client) Get(collection string) (value string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	value, err = redis.String(client.Do("GET", collection))
	if err != nil {
		return "", trace.TraceError(err)
	}
	return value, nil
}

func (c *Client) Set(collection string, value string) (err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	_, err = client.Do("SET", collection, value)
	if err != nil {
		return trace.TraceError(err)
	}
	return nil
}

func (c *Client) Del(collection string) error {
	client := c.pool.Get()
	defer utils.Close(client)

	_, err := client.Do("DEL", collection)
	if err != nil {
		return trace.TraceError(err)
	}
	return nil
}

func (c *Client) RPush(collection string, value any) error {
	client := c.pool.Get()
	defer utils.Close(client)

	_, err := client.Do("RPUSH", collection, value)
	if err != nil {
		return trace.TraceError(err)
	}
	return nil
}

func (c *Client) LPush(collection string, value any) error {
	client := c.pool.Get()
	defer utils.Close(client)

	_, err := client.Do("LPUSH", collection, value)
	if err != nil {
		if err != redis.ErrNil {
			return trace.TraceError(err)
		}
		return err
	}
	return nil
}

func (c *Client) LPop(collection string) (string, error) {
	client := c.pool.Get()
	defer utils.Close(client)

	value, err := redis.String(client.Do("LPOP", collection))
	if err != nil {
		if err != redis.ErrNil {
			return "", trace.TraceError(err)
		}
		return "", nil
	}
	return value, nil
}

func (c *Client) RPop(collection string) (string, error) {
	client := c.pool.Get()
	defer utils.Close(client)

	value, err := redis.String(client.Do("RPOP", collection))
	if err != nil {
		if err != redis.ErrNil {
			return value, trace.TraceError(err)
		}
		return value, err
	}
	return value, nil
}

func (c *Client) LLen(collection string) (int, error) {
	client := c.pool.Get()
	defer utils.Close(client)

	length, err := redis.Int(client.Do("LLEN", collection))
	if err != nil {
		return 0, trace.TraceError(err)
	}
	return length, nil
}

func (c *Client) BRPop(collection string, timeout int) (value string, err error) {
	if timeout <= 0 {
		timeout = 60
	}
	client := c.pool.Get()
	defer utils.Close(client)

	values, err := redis.Strings(client.Do("BRPOP", collection, timeout))
	if err != nil {
		if err != redis.ErrNil {
			return value, trace.TraceError(err)
		}
		return value, err
	}
	return values[1], nil
}

func (c *Client) BLPop(collection string, timeout int) (value string, err error) {
	if timeout <= 0 {
		timeout = 60
	}
	client := c.pool.Get()
	defer utils.Close(client)

	values, err := redis.Strings(client.Do("BLPOP", collection, timeout))
	if err != nil {
		if err != redis.ErrNil {
			return value, trace.TraceError(err)
		}
		return value, err
	}
	return values[1], nil
}

func (c *Client) HSet(collection string, key string, value string) error {
	client := c.pool.Get()
	defer utils.Close(client)

	_, err := client.Do("HSET", collection, key, value)
	if err != nil {
		if err != redis.ErrNil {
			return trace.TraceError(err)
		}
		return err
	}
	return nil
}

func (c *Client) HGet(collection string, key string) (string, error) {
	client := c.pool.Get()
	defer utils.Close(client)

	value, err := redis.String(client.Do("HGET", collection, key))
	if err != nil {
		if err != redis.ErrNil {
			return value, trace.TraceError(err)
		}
		return value, err
	}
	return value, nil
}

func (c *Client) HDel(collection string, key string) error {
	client := c.pool.Get()
	defer utils.Close(client)

	_, err := redis.String(client.Do("HDEL", collection, key))
	if err != nil {
		return trace.TraceError(err)
	}
	return nil
}

func (c *Client) HScan(collection string) (results map[string]string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	var (
		cursor int64
		items  []string
	)

	results = map[string]string{}
	for {
		values, err := redis.Values(client.Do("HSCAN", collection, cursor))
		if err != nil {
			if err != redis.ErrNil {
				return nil, trace.TraceError(err)
			}
			return nil, err
		}

		values, err = redis.Scan(values, &cursor, &items)
		if err != nil {
			if err != redis.ErrNil {
				return nil, trace.TraceError(err)
			}
			return nil, err
		}
		for i := 0; i < len(values); i += 2 {
			key := items[i]
			value := items[i+1]
			results[key] = value
		}
		if cursor == 0 {
			break
		}
	}
	return results, nil
}

func (c *Client) HKeys(collection string) (results []string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	values, err := redis.Strings(client.Do("HKEYS", collection))
	if err != nil {
		if err != redis.ErrNil {
			return nil, trace.TraceError(err)
		}
		return nil, err
	}
	return values, nil
}

func (c *Client) ZAdd(collection string, score float32, value any) (err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	_, err = client.Do("ZADD", collection, score, value)
	if err != nil {
		return trace.TraceError(err)
	}
	return nil
}

func (c *Client) ZCount(collection string, min string, max string) (count int, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	count, err = redis.Int(client.Do("ZCOUNT", collection, min, max))
	if err != nil {
		return 0, trace.TraceError(err)
	}
	return count, nil
}

func (c *Client) ZCountAll(collection string) (count int, err error) {
	return c.ZCount(collection, "-inf", "+inf")
}

func (c *Client) ZScan(collection string, pattern string, count int) (values []string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	values, err = redis.Strings(client.Do("ZSCAN", collection, 0, pattern, count))
	if err != nil {
		if err != redis.ErrNil {
			return nil, trace.TraceError(err)
		}
		return nil, err
	}
	return values, nil
}

func (c *Client) ZPopMax(collection string, count int) (results []string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	results = []string{}

	values, err := redis.Strings(client.Do("ZPOPMAX", collection, count))
	if err != nil {
		if err != redis.ErrNil {
			return nil, trace.TraceError(err)
		}
		return nil, err
	}

	for i := 0; i < len(values); i += 2 {
		v := values[i]
		results = append(results, v)
	}

	return results, nil
}

func (c *Client) ZPopMin(collection string, count int) (results []string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	results = []string{}

	values, err := redis.Strings(client.Do("ZPOPMIN", collection, count))
	if err != nil {
		if err != redis.ErrNil {
			return nil, trace.TraceError(err)
		}
		return nil, err
	}

	for i := 0; i < len(values); i += 2 {
		v := values[i]
		results = append(results, v)
	}

	return results, nil
}

func (c *Client) ZPopMaxOne(collection string) (value string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	values, err := c.ZPopMax(collection, 1)
	if err != nil {
		return "", err
	}
	if values == nil || len(values) == 0 {
		return "", nil
	}
	return values[0], nil
}

func (c *Client) ZPopMinOne(collection string) (value string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	values, err := c.ZPopMin(collection, 1)
	if err != nil {
		return "", err
	}
	if values == nil || len(values) == 0 {
		return "", nil
	}
	return values[0], nil
}

func (c *Client) BZPopMax(collection string, timeout int) (value string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	values, err := redis.Strings(client.Do("BZPOPMAX", collection, timeout))
	if err != nil {
		if err != redis.ErrNil {
			return "", trace.TraceError(err)
		}
		return "", err
	}
	if len(values) < 3 {
		return "", trace.TraceError(errors.ErrorRedisInvalidType)
	}
	return values[1], nil
}

func (c *Client) BZPopMin(collection string, timeout int) (value string, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	values, err := redis.Strings(client.Do("BZPOPMIN", collection, timeout))
	if err != nil {
		if err != redis.ErrNil {
			return "", trace.TraceError(err)
		}
		return "", err
	}
	if len(values) < 3 {
		return "", trace.TraceError(errors.ErrorRedisInvalidType)
	}
	return values[1], nil
}

func (c *Client) Lock(lockKey string) (value int64, err error) {
	client := c.pool.Get()
	defer utils.Close(client)

	lockKey = c.getLockKey(lockKey)

	ts := time.Now().Unix()
	ok, err := client.Do("SET", lockKey, ts, "NX", "PX", 30000)
	if err != nil {
		if err != redis.ErrNil {
			return value, trace.TraceError(err)
		}
		return value, err
	}
	if ok == nil {
		return value, trace.TraceError(errors.ErrorRedisLocked)
	}
	return value, nil
}

func (c *Client) UnLock(lockKey string, value int64) {
	client := c.pool.Get()
	defer utils.Close(client)
	lockKey = c.getLockKey(lockKey)

	getVal, err := redis.Int64(client.Do("GET", lockKey))
	if err != nil {
		log.Errorf("get lockkey error: %s", err.Error())
		return
	}

	if getVal != value {
		log.Errorf("the lockKey value diff: %d, %d", value, getVal)
		return
	}

	v, err := redis.Int64(client.Do("DEL", lockKey))
	if err != nil {
		log.Errorf("unlock failed, error: %s", err.Error())
		return
	}

	if v == 0 {
		log.Infof("unlock failed, lockKey: %s", lockKey)
		return
	}
}

func (c *Client) MemoryStats() (stats map[string]int64, err error) {
	stats = map[string]int64{}
	client := c.pool.Get()
	defer utils.Close(client)

	values, err := redis.Values(client.Do("MEMORY", "STATS"))
	for i, v := range values {
		typeV := reflect.TypeOf(v)
		if typeV.Kind() == reflect.Slice {
			vc, _ := redis.String(v, err)
			if utils.ContainsString(MemoryStatsMetrics, vc) {
				stats[vc], _ = redis.Int64(values[i+1], err)
			}
		}
	}

	if err != nil {
		if err != redis.ErrNil {
			return stats, trace.TraceError(err)
		}
		return stats, err
	}
	return stats, nil
}

func (c *Client) SetBackoffMaxInterval(interval time.Duration) {
	c.backoffMaxInterval = interval
}

func (c *Client) SetTimeout(timeout int) {
	c.timeout = timeout
}

func (c *Client) GetTimeout(timeout int) (res int) {
	if timeout == 0 {
		return c.timeout
	}
	return timeout
}
