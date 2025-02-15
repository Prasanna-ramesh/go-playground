package user

import "github.com/go-chi/chi/v5"

func AddRoutes(router *chi.Mux) {
	router.Get("/users", getUsersHandler)
	router.Get("/users/{userId}", getUserHandler)
	router.Post("/users", createUserHandler)
	router.Put("/users/{userId}", modifyUserHandler)
	router.Delete("/users/{userId}", deleteUserHandler)
}
