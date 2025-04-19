package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]CacheEntry
	mu           sync.RWMutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheEntries: map[string]CacheEntry{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mu.Lock()
	cache.cacheEntries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	cache.mu.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	res, exists := cache.cacheEntries[key]
	if exists {
		cache.mu.RUnlock()
		return res.val, true
	} else {
		cache.mu.RUnlock()
		return nil, false
	}
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for entry, val := range cache.cacheEntries {
				limit := val.createdAt.Add(interval)
				if limit.Before(time.Now()) {
					cache.mu.Lock()
					delete(cache.cacheEntries, entry)
					cache.mu.Unlock()
				}
			}
		}
	}

}
