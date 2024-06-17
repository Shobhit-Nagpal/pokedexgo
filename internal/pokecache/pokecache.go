package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	mtx := &sync.Mutex{}
	cache := Cache{
    cache: make(map[string]cacheEntry),
		mu: mtx,
	}

	cache.readLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	data, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}

	return data.val, true
}

func (c *Cache) readLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		for {
			time.Sleep(interval)
			done <- true
		}
	}()
	go func() {
		for {

			select {
			case <-done:
				for k, v := range c.cache {
					if time.Now().Sub(v.createdAt) >= interval {
						delete(c.cache, k)
					}
				}
			}
		}
	}()
}
