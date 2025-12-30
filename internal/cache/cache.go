package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]CacheEntry
	Interval time.Duration
	Mu       *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	entry := CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
	c.Mu.Lock()
	c.Entries[key] = entry
	c.Mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	entry, ok := c.Entries[key]
	defer c.Mu.Unlock()
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reap() {
	cutoff := time.Now().UTC().Add(-c.Interval)
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for k, v := range c.Entries {
		if v.createdAt.Before(cutoff) {
			delete(c.Entries, k)
		}
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap()
	}

}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Entries:  map[string]CacheEntry{},
		Mu:       &sync.Mutex{},
		Interval: interval,
	}
	go cache.reapLoop()
	return cache
}
