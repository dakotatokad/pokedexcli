package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	// Maps aren't safe for concurrent use, so we need to handle synchronization
	// You could also use channels
	cache map[string]cacheEntry
	mux   *sync.Mutex // You need to lock the cache when accessing it from multiple goroutines
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{}, // Initialize the mutex
	}
	go c.reapLoop(interval) // Start the reap loop with a 5-minute interval
	return c
}

func (c *Cache) Add(key string, val []byte) {
	// Assuming no other routine has mutex locked, this will gain exclusive access to mutex until it is unlocked
	c.mux.Lock()         // Lock the cache for writing
	defer c.mux.Unlock() // Ensure the mutex is unlocked after this function returns
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()         // Lock the cache for writing
	defer c.mux.Unlock() // Ensure the mutex is unlocked after this function returns
	cachE, ok := c.cache[key]
	return cachE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C { // Reap the cache every interval
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()         // Lock the cache for writing
	defer c.mux.Unlock() // Ensure the mutex is unlocked after this function returns
	timeSince := time.Now().UTC().Add(-interval)

	for k, v := range c.cache {
		if v.createdAt.Before(timeSince) {
			delete(c.cache, k)
		}
	}
}
