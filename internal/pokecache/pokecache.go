package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type cache struct {
	CacheMap map[string]cacheEntry
	Mu       *sync.RWMutex
}

var mut sync.RWMutex

func NewCache(interval time.Duration) *cache {
	newCache := cache{
		CacheMap: make(map[string]cacheEntry),
		Mu:       &mut,
	}

	go newCache.reapLoop(interval)

	return &newCache
}

func (c *cache) Add(key string, v []byte) {
	c.Mu.Lock()

	defer c.Mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
	c.CacheMap[key] = entry
	fmt.Println("Successful add")
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()

	defer c.Mu.Unlock()

	v, ok := c.CacheMap[key]

	if !ok {
		fmt.Println("Error: key does not exist")
		return nil, false
	}

	fmt.Println("Successful get")
	return v.val, true
}

func (c *cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.Mu.Lock()

		for k, v := range c.CacheMap {
			startTime := time.Since(v.createdAt)
			if startTime > interval {
				delete(c.CacheMap, k)
			}
		}
		c.Mu.Unlock()
	}
}
