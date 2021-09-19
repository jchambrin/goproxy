package cache

import (
	"testing"
	"time"

	"github.com/jchambrin/goproxy/pkg/proxy"
)

func TestCacheMemoryTTL(t *testing.T) {
	memCache := NewMemoryCache(2 * time.Second)
	key := proxy.KeyCache{"/test", "HEAD"}
	memCache.Put(key, &proxy.CacheData{})
	if _, ok := memCache.Get(key); !ok {
		t.Error("cache should be filled with one entry")
	}

	time.Sleep(4 * time.Second)
	key2 := proxy.KeyCache{"/test", "GET"}
	memCache.Put(key2, &proxy.CacheData{})
	if _, ok := memCache.Get(key); ok {
		t.Error("cache should be empty now since TTL has expired")
	}
	if _, ok := memCache.Get(key2); !ok {
		t.Errorf("key : %v should be in cache", key2)
	}
}
