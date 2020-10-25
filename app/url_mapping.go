package app

import (
	"bookstore_users-api/controllers/ping"
	"bookstore_users-api/controllers/users"

	"github.com/go-chi/chi"
)

func mapUrls(router *chi.Mux) {
	router.Get("/ping", ping.Ping())
	router.Get("/users/{user_id}", users.GetUser())
	router.Get("/users/search", users.SearchUser())
	router.Post("/users", users.CreateUser())
}
