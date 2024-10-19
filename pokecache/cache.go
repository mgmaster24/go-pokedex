package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mu      *sync.Mutex
}

type CacheEntry struct {
	createAt time.Time
	Val      []byte
}

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]CacheEntry),
		mu:      &sync.Mutex{},
	}

	go cache.reapLoop(duration)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	entry := CacheEntry{
		createAt: time.Now(),
		Val:      val,
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.entries[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, found := cache.entries[key]
	return entry.Val, found
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

func (cache *Cache) reap(now time.Time, last time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	for k, v := range cache.entries {
		if v.createAt.Before(now.Add(-last)) {
			delete(cache.entries, k)
		}
	}
}
