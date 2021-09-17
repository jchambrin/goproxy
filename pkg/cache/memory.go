package cache

import (
	"sync"
	"time"

	"github.com/jchambrin/goproxy/pkg/proxy"
)

type memory struct {
	mu sync.RWMutex
	m  map[proxy.KeyCache]*proxy.CacheData
}

func NewMemoryCache(TTL time.Duration) *memory {
	return &memory{
		m: make(map[proxy.KeyCache]*proxy.CacheData),
	}
}

func (m *memory) Get(key proxy.KeyCache) (*proxy.CacheData, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	res, ok := m.m[key]
	return res, ok
}

func (m *memory) Put(key proxy.KeyCache, data *proxy.CacheData) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = data
}
