package pokecache

import (
    "fmt"
    "time"
)


// Cache is a struct that holds the cache data
type cacheEntry struct {
    createdAt time.Time
    val []bytes

