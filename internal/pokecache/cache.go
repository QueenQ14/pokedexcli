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
	storage  map[string]cacheEntry
	interval time.Duration
	mu       *sync.Mutex
}

func NewCache(time time.Duration) Cache {
	c := Cache{
		storage:  map[string]cacheEntry{},
		interval: time,
		mu:       &sync.Mutex{},
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cEntry, ok := c.storage[key]
	return cEntry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.storage {
			c.reap(key, entry.createdAt)
		}
		c.mu.Unlock()
	}
}

func (c *Cache) reap(key string, entryTime time.Time) error {
	if time.Since(entryTime) >= c.interval {
		delete(c.storage, key)
	}
	return nil
}
