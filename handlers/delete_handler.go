// handlers/delete_handler.go
package handlers

import (
	"cache_server/cache"
	"fmt"
	"net/http"
)

func DeleteHandler(cacheInstance *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		cacheInstance.Remove(key)
		fmt.Fprintf(w, "Deleted %s", key)
	}
}
