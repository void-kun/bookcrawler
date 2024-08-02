package errors

import (
	"fmt"
)

var (
	ErrorRedisInvalidType = NewRedisError("invalid type")
	ErrorRedisLocked      = NewRedisError("locked")
)

func NewRedisError(msg string) error {
	return fmt.Errorf("%s: %s", errorPrefixRedis, msg)
}
