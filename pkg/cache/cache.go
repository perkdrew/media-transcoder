package cache

import (
	"github.com/perkdrew/pkg/cache"
)

// Cache represents a cache with a specific capacity and eviction policy.
type Cache struct {
	capacity        int
	cacheMap        map[string]interface{}
	evictionPolicy eviction.EvictionPolicy
}

// NewCache creates a new cache with the specified capacity and eviction policy.
func NewCache(capacity int, evictionPolicy EvictionPolicy) *Cache {
	return &Cache{
		capacity:        capacity,
		cacheMap:        make(map[string]interface{}),
		evictionPolicy: evictionPolicy,
	}
}

// Set adds or updates a key-value pair in the cache.
func (c *Cache) Set(key string, value interface{}) {
	c.cacheMap[key] = value

	// If the cache has exceeded its capacity, apply the eviction policy
	if len(c.cacheMap) > c.capacity {
		c.evictionPolicy.Evict(c)
	}
}

// Get retrieves the value associated with the specified key from the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	value, exists := c.cacheMap[key]
	return value, exists
}
