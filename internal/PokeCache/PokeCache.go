package PokeCache

import (
	"sync"
	"time"
)

type Cache struct {
	Data map[string]CacheEntry
	rwMu sync.RWMutex
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	res, ok := c.Data[url]
	return res.Val, ok
}

func (c *Cache) Add(url string, val []byte) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.Data[url] = CacheEntry{time.Now(), val}
}

func (c *Cache) reapLoop(duration time.Duration) {
	for {
		time.Sleep(duration)
		c.rwMu.Lock()
		for k, v := range c.Data {
			if time.Since(v.CreatedAt) > duration {
				delete(c.Data, k)
			}
		}
		c.rwMu.Unlock()
	}
}
func NewCache(duration time.Duration) *Cache {
	newCache := &Cache{
		Data: make(map[string]CacheEntry),
	}
	go newCache.reapLoop(duration)
	return newCache
}
