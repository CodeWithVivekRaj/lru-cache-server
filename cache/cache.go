package cache

import (
	"container/list"
	"fmt"
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
	Items    map[string]*list.Element // Change to exported field
	lruList  *list.List
	Mutex    sync.Mutex // Change to exported field
}

// NewLRUCache initializes a new LRUCache instance
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		Items:    make(map[string]*list.Element),
		lruList:  list.New(),
	}
}

// Get retrieves a value from the cache
func (c *LRUCache) Get(key string) interface{} {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if elem, ok := c.Items[key]; ok {
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

// Remove removes an item from the cache
func (c *LRUCache) Remove(key string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if elem, ok := c.Items[key]; ok {
		delete(c.Items, key)
		c.lruList.Remove(elem)
	}
}

// Set sets a value in the cache
func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// Check if the key already exists in the cache
	if elem, ok := c.Items[key]; ok {
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
		c.Items[key] = elem
		// Evict the least recently used item if the cache is full
		if len(c.Items) > c.capacity {
			c.evict()
		}
	}
}

// evict removes the least recently used item from the cache
func (c *LRUCache) evict() {
	elem := c.lruList.Back()
	if elem != nil {
		item := elem.Value.(*CacheItem)
		delete(c.Items, item.key)
		c.lruList.Remove(elem)
	}
}

// GetDataArray method to retrieve stored data from the cache as an array of key-value pairs
func (c *LRUCache) GetDataArray() []interface{} {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	data := make([]interface{}, 0, len(c.Items))

	// Iterate over the items in the cache and extract the key-value pairs
	for key, elem := range c.Items {
		item := elem.Value.(*CacheItem)
		// Check if item has expired
		if time.Now().After(item.expiration) {
			// c.Remove(key)
			fmt.Println(key)
		} else {
			// Create a map for each cache item
			itemData := map[string]interface{}{
				"key":      item.key,
				"value":    item.value,
				"duration": time.Until(item.expiration).String(),
			}
			// Append the map to the data slice
			data = append(data, itemData)
		}
	}

	return data
}

// Cache is an interface for cache operations
type Cache interface {
	Get(key string) interface{}
	Set(key string, value interface{}, expiration time.Duration)
	Remove(key string)
	GetDataArray() []interface{}
}
