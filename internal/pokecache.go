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
	data map[string]cacheEntry
	mut  sync.Mutex
}

func NewCache(d time.Duration) *cache {
	nCache := &cache{
		data: make(map[string]cacheEntry),
	}
	go nCache.reapLoop(d)
	return nCache
}

func (c *cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *cache) reapLoop(d time.Duration) {
	tk := time.NewTicker(d)
	for range tk.C {
		go func() {
			c.mut.Lock()
			defer c.mut.Unlock()
			for key, val := range c.data {
				if time.Since(val.createdAt) > d {
					delete(c.data, key)
				}
			}
		}()
	}
}
