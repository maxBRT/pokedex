package pokecache

import (
	"sync"
	"time"
)

// cacheEntry represents a single entry in the cache with its creation time and value.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// cache is a thread-safe in-memory cache with a map for storing data and a mutex for synchronization.
type cache struct {
	data map[string]cacheEntry
	mut  sync.Mutex
}

// NewCache creates a new cache instance and starts a background goroutine to remove expired entries.
func NewCache(d time.Duration) *cache {
	nCache := &cache{
		data: make(map[string]cacheEntry),
	}
	go nCache.reapLoop(d) // Start the cleanup loop.
	return nCache
}

// Add inserts a new key-value pair into the cache.
func (c *cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves a value from the cache by its key. Returns false if the key does not exist.
func (c *cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

// reapLoop periodically removes expired entries from the cache.
func (c *cache) reapLoop(d time.Duration) {
	tk := time.NewTicker(d) // Create a ticker that triggers at the specified duration.
	for range tk.C {
		go func() {
			c.mut.Lock()
			defer c.mut.Unlock()
			// Iterate through the cache and delete entries older than the specified duration.
			for key, val := range c.data {
				if time.Since(val.createdAt) > d {
					delete(c.data, key)
				}
			}
		}()
	}
}
