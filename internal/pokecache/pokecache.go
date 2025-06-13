package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{ 
		cache: make(map[string]cacheEntry),
	}
	// TODO: Handle a possible race condition here if the reapLoop conflicts with accessing the cache
	go c.reapLoop(interval) // Start the reap loop with a 5-minute interval
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cachE, ok := c.cache[key]
	return cachE.val, ok
}

func (c *Cache) reap(interval time.Duration) {

	timeSince := time.Now().UTC().Add(-interval)

	for k, v := range c.cache {
		if v.createdAt.Before(timeSince) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C { // Reap the cache every interval
		c.reap(interval)
	}
}
