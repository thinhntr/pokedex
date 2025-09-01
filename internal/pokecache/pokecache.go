package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	res, ok := c.data[key]
	c.mu.RUnlock()
	if !ok {
		return []byte{}, ok
	}
	return res.val, ok
}

func (c Cache) reapLoop(duration time.Duration, ticker *time.Ticker) {
	for currentTime := range ticker.C {
		c.mu.Lock()
		for k, v := range c.data {
			if currentTime.Add(-duration).After(v.createdAt) {
				delete(c.data, k)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.RWMutex{},
	}
	ticker := time.NewTicker(interval)
	go cache.reapLoop(interval, ticker)
	return cache
}
