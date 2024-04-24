package handlers

import (
	"cache_server/cache"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SetHandler(cache *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		expirationStr := r.URL.Query().Get("expiration")
		expiration, _ := time.ParseDuration(expirationStr)

		cache.Set(key, value, expiration)

		// Create a map to hold the response data
		response := map[string]interface{}{
			"message": fmt.Sprintf("Set %s=%s with expiration %s", key, value, expiration),
		}

		// Encode the response data into JSON format
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the response writer
		w.Write(jsonResponse)
	}
}
