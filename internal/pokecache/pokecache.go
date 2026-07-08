package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.RWMutex
}

func NewCache(interval time.Duration) cache {
	var mut sync.RWMutex
	newCache := cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       &mut,
	}

	go newCache.reapLoop(interval)

	return newCache
}

func (c *cache) Add(key string, v []byte) {
	c.mu.Lock()

	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
	c.cacheMap[key] = entry
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()

	defer c.mu.Unlock()

	v, ok := c.cacheMap[key]

	if !ok {
		return nil, false
	}

	return v.val, true
}

func (c *cache) reapLoop(interval time.Duration) {
	c.mu.Lock()

	defer c.mu.Unlock()

	ticker := time.NewTicker(interval)

	for range ticker.C {
		for k, v := range c.cacheMap {
			startTime := time.Since(v.createdAt)
			if startTime < interval {
				delete(c.cacheMap, k)
			}
		}
	}

}
