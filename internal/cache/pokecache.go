package pokeCache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.mu.Lock()
		for name, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, name)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()

	val, ok := c.entries[key]

	c.mu.Unlock()
	if ok {
		return val.val, true
	}
	return nil, false
}
