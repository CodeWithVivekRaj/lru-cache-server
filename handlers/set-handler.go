package handlers

import (
	"cache_server/cache"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SetRequest struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration string `json:"expiration"`
}

func SetHandler(cache *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode JSON request body into SetRequest struct
		var requestData SetRequest
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Parse expiration duration
		expiration, err := time.ParseDuration(requestData.Expiration)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Set the key-value pair in the cache
		cache.Set(requestData.Key, requestData.Value, expiration)

		// Create a map to hold the response data
		response := map[string]interface{}{
			"message": fmt.Sprintf("Set %s=%s with expiration %s", requestData.Key, requestData.Value, requestData.Expiration),
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
