package main

import (
	"cache_server/cache"
	"cache_server/handlers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	// Create a new LRUCache with capacity 10
	cacheInstance := cache.NewLRUCache(10)

	// Register cache API endpoints
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// console.log("i am hiy")
		handleWebSocket(cacheInstance, w, r)
	})
	http.HandleFunc("/get", handlers.GetHandler(cacheInstance))
	http.HandleFunc("/set", handlers.SetHandler(cacheInstance))
	http.HandleFunc("/delete", handlers.DeleteHandler(cacheInstance))

	// WebSocket endpoint

	// CORS middleware with custom options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                       // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow all HTTP methods
		AllowedHeaders:   []string{"*"},                                       // Allow all headers
		AllowCredentials: true,                                                // Allow credentials
		MaxAge:           3600,                                                // Cache preflight requests for 1 hour
	})

	// Apply CORS middleware to each handler
	handler := c.Handler(http.DefaultServeMux)

	// Create server with timeouts
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler, // Use the CORS-wrapped handler
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Cache server is running on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}

func handleWebSocket(cacheInstance cache.Cache, w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Send initial cache data to the client
	sendCacheUpdates(conn, cacheInstance)

	// Create a ticker to trigger cache updates at regular intervals (every second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sendCacheUpdates(conn, cacheInstance)
		case <-r.Context().Done(): // Handle client disconnect
			return
		}
	}
}

func sendCacheUpdates(conn *websocket.Conn, cacheInstance cache.Cache) {
	// Retrieve current cache data
	cacheData := cacheInstance.GetDataArray()

	// Marshal cache data to JSON
	jsonData, err := json.Marshal(cacheData)
	fmt.Println(jsonData)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return
	}

	// Write JSON data to WebSocket connection
	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("Error writing message to WebSocket:", err)
		return
	}
}
