package handlers

import (
	"cache_server/cache"
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func DeleteHandler(cacheInstance *cache.LRUCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		key := r.URL.Query().Get("key")
		cacheInstance.Remove(key)

		// Create the response JSON object
		response := Response{Message: "Deleted " + key}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the content type header and send the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
