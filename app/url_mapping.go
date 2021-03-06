package app

import (
	"bookstore_users-api/controllers/ping"
	"bookstore_users-api/controllers/users"

	"github.com/go-chi/chi"
)

func mapUrls(router *chi.Mux) {
	router.Get("/ping", ping.Ping())

	router.Post("/users", users.CreateUser())
	router.Get("/users/{user_id}", users.GetUser())
	router.Put("/users/{user_id}", users.UpdateUser())
	router.Patch("/users/{user_id}", users.UpdateUser())
	router.Delete("/users/{user_id}", users.DeleteUser())
	router.Get("/internal/users/search", users.Search())
}
