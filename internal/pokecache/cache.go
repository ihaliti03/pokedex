package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	// The cache will be shared across goroutines eventually. Even now, Add/Get 
	// from multiple goroutines would race on the map. Go maps are not safe for 
	// concurrent read/write — a concurrent map write panics. The mutex protects it.
	mu			sync.Mutex
	store		map[string]cacheEntry
	interval 	time.Duration
}

type cacheEntry struct {
    createdAt 	time.Time
    val       	[]byte
}
//Don't put a global CacheMap inside pokecache. That makes pokecache own a singleton — every consumer gets the same cache, and you can't test it in isolation. Bad.
//var CacheMap = NewCache(5 * time.Second)

func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        store:    make(map[string]cacheEntry),
        interval: interval,
    }
    go c.reapLoop()
    return c
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.store[key] = cacheEntry{
        createdAt: time.Now(),
        val:       val,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    entry, ok := c.store[key]
    if !ok {
        return nil, false
    }
    return entry.val, true
}

func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    defer ticker.Stop()
    for range ticker.C {
        c.reap()
    }
}

func (c *Cache) reap() {
    c.mu.Lock()
    defer c.mu.Unlock()

	// delete any existing cache before the interval
    cutoff := time.Now().Add(-c.interval)
    for key, entry := range c.store {
        if entry.createdAt.Before(cutoff) {
            delete(c.store, key)
        }
    }
}