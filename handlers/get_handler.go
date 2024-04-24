package handlers

import (
	"cache_server/cache"
	"encoding/json"
	"net/http"
)

func GetHandler(cacheInstance *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := cacheInstance.Get(key)

		// Create a map to hold the response data
		response := map[string]interface{}{
			"key":     key,
			"value":   value,
			"message": "key value fethed",
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
