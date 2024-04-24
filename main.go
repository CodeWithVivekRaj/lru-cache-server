package main

import (
	"cache_server/cache"
	"cache_server/handlers"
	"fmt"
	"net/http"
)

func main() {
	// Create a new LRUCache with capacity 10

	cacheInstance := cache.NewLRUCache(10)

	// Register cache API endpoints
	http.HandleFunc("/get", handlers.GetHandler(cacheInstance))
	http.HandleFunc("/set", handlers.SetHandler(cacheInstance))
	http.HandleFunc("/delete", handlers.DeleteHandler(cacheInstance))

	fmt.Println("Cache server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
