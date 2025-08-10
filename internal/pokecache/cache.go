package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mu    sync.RWMutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		cache: make(map[string]CacheEntry),
	}
	go newCache.reapLoop(interval)
	return &newCache
}

func (a *Cache) Add(key string, val []byte) {
	newEntry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.cache[key] = newEntry
}

func (g *Cache) Get(key string) ([]byte, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	entry, ok := g.cache[key]
	if ok {
		return entry.val, ok
	} else {
		return nil, false
	}
}

func (r *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			r.mu.Lock()
			for key, entry := range r.cache {
				if time.Since(entry.createdAt) > interval {
					delete(r.cache, key)
				}
			}
			r.mu.Unlock()
		}

	}

}
