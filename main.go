package main

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"GoogleProject/nearby"
	"GoogleProject/auth"

)

func main() {
	serv := Init()
	go setupRoutes(serv)
	go authRoutes()
	serv.Post("/api/v1/NearbyPlaces", nearby.HandleNearbyPlaces)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", serv)
}
func authRoutes(serv *chi.Mux) {
	serv.Post("/api/v1/Login", auth.HandleLogin)
	serv.Post("/api/v1/Register", auth.HandleRegister)
}
func setupRoutes(serv *chi.Mux) {
	serv.Get("/api/v1/HealthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Everything works correctly!!"))
		})
}

func Init() *chi.Mux {
	serv := chi.NewRouter()
	return serv
}
