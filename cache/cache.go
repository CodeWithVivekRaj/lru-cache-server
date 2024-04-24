// cache/cache.go
package cache

import (
	"container/list"
	"sync"
	"time"
)

type CacheItem struct {
	key        string
	value      interface{}
	expiration time.Time
}

type LRUCache struct {
	capacity int
	items    map[string]*list.Element
	lruList  *list.List
	mutex    sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		lruList:  list.New(),
	}
}

// Get method to retrieve a value from the cache
func (c *LRUCache) Get(key string) interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.items[key]; ok {
		item := elem.Value.(*CacheItem)
		// Check if item has expired
		if time.Now().After(item.expiration) {
			c.Remove(key) // Corrected method call to Remove
			return nil
		}
		// Move item to front of the list (most recently used)
		c.lruList.MoveToFront(elem)
		return item.value
	}
	return nil
}

// Remove method to remove an item from the cache
func (c *LRUCache) Remove(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.items[key]; ok {
		delete(c.items, key)
		c.lruList.Remove(elem)
	}
}

// Implement other cache methods such as Set, Evict here
// Set method to set a value in the cache
func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Check if the key already exists in the cache
	if elem, ok := c.items[key]; ok {
		// Update the value and expiration time
		item := elem.Value.(*CacheItem)
		item.value = value
		item.expiration = time.Now().Add(expiration)
		// Move the item to the front of the list (most recently used)
		c.lruList.MoveToFront(elem)
	} else {
		// Create a new cache item
		item := &CacheItem{
			key:        key,
			value:      value,
			expiration: time.Now().Add(expiration),
		}
		// Add the item to the cache
		elem := c.lruList.PushFront(item)
		c.items[key] = elem
		// Evict the least recently used item if the cache is full
		if len(c.items) > c.capacity {
			c.evict()
		}
	}
}

// Evict method to remove the least recently used item from the cache
func (c *LRUCache) evict() {
	elem := c.lruList.Back()
	if elem != nil {
		item := elem.Value.(*CacheItem)
		delete(c.items, item.key)
		c.lruList.Remove(elem)
	}
}
