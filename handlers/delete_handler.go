package handlers

import (
	"cache_server/cache"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteHandler(cacheInstance *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		cacheInstance.Remove(key)

		// Create a map to hold the response data
		response := map[string]interface{}{
			"message": fmt.Sprintf("Deleted %s", key),
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
