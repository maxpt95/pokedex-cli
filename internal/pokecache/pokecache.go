package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           *sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{cacheEntries: make(map[string]cacheEntry), mu: &sync.Mutex{}}
	go cache.reapLoop(interval)
	return &cache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for cutOff := range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cacheEntries {
			if cutOff.After(entry.createdAt) {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
