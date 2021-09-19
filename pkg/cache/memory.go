package cache

import (
	"sync"
	"time"

	"github.com/jchambrin/goproxy/pkg/proxy"
)

type memory struct {
	mu   sync.RWMutex
	m    map[proxy.KeyCache]*proxy.CacheData
	heap []heapEntry
}

type heapEntry struct {
	key      proxy.KeyCache
	creation int64
}

func NewMemoryCache(TTL time.Duration) *memory {
	res := &memory{
		m:    make(map[proxy.KeyCache]*proxy.CacheData),
		heap: make([]heapEntry, 0),
	}

	go func(TTL time.Duration) {
		ttlUnix := int64(TTL.Seconds())
		for now := range time.Tick(1 * time.Second) {
			res.mu.Lock()
			for _, v := range res.heap {
				if now.Unix()-v.creation > ttlUnix {
					delete(res.m, v.key)
					res.heap = res.heap[1:]
				} else {
					break
				}
			}
			res.mu.Unlock()
		}
	}(TTL)

	return res
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
	m.heap = append(m.heap, heapEntry{key, time.Now().Unix()})
}
