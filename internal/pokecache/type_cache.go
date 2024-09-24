package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mux     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: map[string]cacheEntry{},
		mux:     &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}
