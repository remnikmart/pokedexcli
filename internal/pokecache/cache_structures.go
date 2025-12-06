package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	ttl         time.Duration
	checkPeriod time.Duration
	data        map[string]cacheEntry
	mu          sync.RWMutex
}

func NewCache(ttl int) *Cache {
	cache := Cache{}
	cache.data = make(map[string]cacheEntry)
	cache.checkPeriod = time.Duration(ttl) * time.Second
	cache.ttl = time.Duration(ttl) * time.Second
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	_, exists := c.data[key]
	if !exists {
		c.data[key] = cacheEntry{createdAt: time.Now(), val: val}
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, exists := c.data[key]
	if exists {
		return val.val, true
	}
	return nil, false
}

func (c *Cache) removeKey(key string) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

func (c *Cache) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.checkPeriod)
	for range ticker.C {
		c.clearStaleData()
	}
}

func (c *Cache) clearStaleData() {
	for k, v := range c.data {
		if time.Since(v.createdAt) > c.ttl {
			c.removeKey(k)
		}
	}
}
