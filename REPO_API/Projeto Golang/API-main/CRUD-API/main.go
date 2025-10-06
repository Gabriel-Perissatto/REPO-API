package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustavolubacheski/API/CRUD-API/routes"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Listening on port 8080")

	r := routes.NewRouter()

	// Wrap the router with CORS middleware
	handler := enableCORS(r)

	log.Fatal(http.ListenAndServe(":8080", handler)) // Single blocking call with your router
}
