package handlers

import (
	"cache_server/cache"
	"fmt"
	"net/http"
)

func GetHandler(cacheInstance *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := cacheInstance.Get(key) // Update the variable name to cacheInstance
		fmt.Fprintf(w, "%v", value)
	}
}
