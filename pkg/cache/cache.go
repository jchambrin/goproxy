package cache

import (
	"net/http"
)

type Storage interface {
	IsPresent()
	Store()
	Purge()
}

type Config struct {
	Enabled        bool
	AllowedMethods []string
}

func (c *Config) Send(r *http.Request) {
	if c.Enabled && contains(r.Method, c.AllowedMethods) {
		// TODO search into cache
	} else {
		// TODO proxy without cache
	}
}

func contains(s string, arr []string) bool {
	for _, a := range arr {
		if s == a {
			return true
		}
	}
	return false
}
