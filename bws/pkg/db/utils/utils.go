package utils

import (
	"io"

	"github.com/apex/log"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.WithError(err).Error("closing failed")
	}
}

func ContainsString(list []string, item string) bool {
	for _, d := range list {
		if d == item {
			return true
		}
	}
	return false
}
