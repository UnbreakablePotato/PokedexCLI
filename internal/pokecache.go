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
	lock     sync.RWMutex
}

func NewCache(t time.Duration) cache {

	return cache{}
}
