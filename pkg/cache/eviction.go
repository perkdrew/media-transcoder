package cache

import (
	"container/list"
)

// EvictionPolicy represents the cache eviction policy.
type EvictionPolicy interface {
	Evict(c *Cache) *CacheItem
}

// LRU (Least Recently Used) is an eviction policy based on the least recently used item.
type LRU struct{}

// Evict evicts the least recently used item from the cache.
func (lru *LRU) Evict(c *Cache) *CacheItem {
	// Access the underlying doubly linked list of items in the cache
	itemList := c.itemList

	// Get the least recently used item (front of the list)
	item := itemList.Front()

	// Remove the item from the list
	itemList.Remove(item)

	// Get the cache item associated with the list element
	cacheItem := item.Value.(*CacheItem)

	// Remove the item from the cache map
	delete(c.cacheMap, cacheItem.key)

	return cacheItem
}
