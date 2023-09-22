package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-User-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/", app.Broker)

	mux.Post("/auth/authenticate", app.Authenticate)
	mux.Post("/auth/create-user", app.CreateUser)
	mux.Post("/auth/update-user", app.UpdateUser)
	mux.Post("/auth/delete-user", app.DeleteUser)
	mux.Post("/auth/get-user-by-id", app.GetUserById)
	mux.Get("/auth/get-all-users", app.GetAllUsers)

	mux.Post("/auth/create-privilege", app.CreatePrivilege)
	mux.Post("/auth/update-privilege", app.UpdatePrivilege)
	mux.Post("/auth/delete-privilege", app.DeletePrivilege)
	mux.Post("/auth/get-privilege-by-id", app.GetPrivilegeById)
	mux.Get("/auth/get-all-privileges", app.GetAllPrivileges)

	mux.Post("/project/create-project", app.CreateProject)
	mux.Post("/project/update-project", app.UpdateProject)
	mux.Post("/project/delete-project", app.DeleteProject)
	mux.Post("/project/get-project-by-id", app.GetProjectById)
	mux.Get("/project/get-all-projects", app.GetAllProjects)

	mux.Post("/email/send", app.SendEmail)

	return mux
}
