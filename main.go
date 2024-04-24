package main

import (
	"cache_server/cache"
	"cache_server/handlers"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
)

func main() {
	// Create a new LRUCache with capacity 10
	cacheInstance := cache.NewLRUCache(10)

	// Register cache API endpoints
	http.HandleFunc("/get", handlers.GetHandler(cacheInstance))
	http.HandleFunc("/set", handlers.SetHandler(cacheInstance))
	http.HandleFunc("/delete", handlers.DeleteHandler(cacheInstance))

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
