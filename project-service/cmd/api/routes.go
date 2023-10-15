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

	mux.Get("/project/get-all-projects", app.GetAllProjects)
	mux.Post("/project/get-project-by-id", app.GetProjectById)
	mux.Post("/project/create-project", app.CreateProject)
	mux.Post("/project/update-project", app.UpdateProject)
	mux.Post("/project/delete-project", app.DeleteProject)

	mux.Post("/project/add-project-note", app.AddProjectNote)
	mux.Post("/project/delete-project-note", app.RemoveProjectNote)

	mux.Get("/project/get-all-sub-projects", app.GetAllSubProjects)
	mux.Post("/project/get-sub-project-by-id", app.GetSubProjectById)
	mux.Post("/project/create-sub-project", app.CreateSubProject)
	mux.Post("/project/update-sub-project", app.UpdateSubProject)
	mux.Post("/project/delete-sub-project", app.DeleteSubProject)

	return mux
}
