package handlers

import (
	"cache_server/cache"
	"fmt"
	"net/http"
	"time"
)

func SetHandler(cache *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		expiration, _ := time.ParseDuration(r.URL.Query().Get("expiration"))

		cache.Set(key, value, expiration)
		fmt.Fprintf(w, "Set %s=%s with expiration %s", key, value, expiration)
	}
}
