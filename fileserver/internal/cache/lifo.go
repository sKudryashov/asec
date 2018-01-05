package cache

import (
	"sync"

	model "github.com/sKudryashov/asec/fileserver/internal/platform"
)

const (
	cacheLength = 10
)

// Cache performs simple LIFO cache
type Cache struct {
	storage []model.FileInfo
	mu      sync.Mutex
}

// New returns new LIFO Cache
func New() *Cache {
	c := new(Cache)
	c.storage = make([]model.FileInfo, 0, cacheLength+5)
	return c
}

// GetAll returns last items
func (c *Cache) GetAll() []model.FileInfo {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.storage
}

// Set adds data to cache and evicts excessive data
func (c *Cache) Set(value model.FileInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage = append(c.storage, value)
	c.evict()
}

func (c *Cache) evict() {
	if len(c.storage) == cacheLength+1 {
		storageShifted := c.storage[1 : cacheLength+1]
		c.storage = storageShifted
	}
}
