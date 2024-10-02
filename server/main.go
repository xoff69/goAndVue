package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type PlayerHandler struct {
}

func (b PlayerHandler) listPlayers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(ListPlayers())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Mount("/players", PlayerRoutes())
	http.ListenAndServe(":3002", r)
}
func PlayerRoutes() chi.Router {
	r := chi.NewRouter()
	playerHandler := PlayerHandler{}
	r.Get("/", playerHandler.listPlayers)
	return r
}
