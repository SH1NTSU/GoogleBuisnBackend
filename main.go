package main

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"GoogleProject/nearby"
)

func main() {
	serv := chi.NewRouter()
	go setupRoutes(serv)
	serv.Post("/api/v1/NearbyPlaces", nearby.HandleNearbyPlaces)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", serv)
}

func setupRoutes(serv *chi.Mux) {
	serv.Get("/api/v1/HealthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Everything works correctly!!"))
		})
}
