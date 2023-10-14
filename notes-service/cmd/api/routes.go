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

	mux.Post("/notes/create-project-note", app.CreateProjectNote)
	mux.Get("/notes/get-project-note-by-id", app.GetProjectNoteById)
	mux.Post("/notes/update-project-note", app.UpdateProjectNote)
	mux.Post("/notes/delete-project-note", app.DeleteProjectNote)
	mux.Post("/notes/get-all-notes-by-project-id", app.GetAllProjectNotesByProjectId)
	mux.Post("/notes/get-all-notes-by-user-id", app.GetAllProjectNotesByUserId)

	mux.Post("/notes/create-income-note", app.CreateIncomeNote)
	mux.Get("/notes/get-income-note-by-id", app.GetIncomeNoteById)
	mux.Post("/notes/update-income-note", app.UpdateIncomeNote)
	mux.Post("/notes/delete-income-note", app.DeleteIncomeNote)
	mux.Post("/notes/get-all-income-notes-by-income-id", app.GetAllIncomeNotesByIncomeId)
	mux.Post("/notes/get-all-income-notes-by-user-id", app.GetAllIncomeNotesByUserId)

	mux.Post("/notes/create-expense-note", app.CreateExpenseNote)
	mux.Get("/notes/get-expense-note-by-id", app.GetExpenseNoteById)
	mux.Post("/notes/update-expense-note", app.UpdateExpenseNote)
	mux.Post("/notes/delete-expense-note", app.DeleteExpenseNote)
	mux.Post("/notes/get-all-expense-notes-by-expense-id", app.GetAllExpenseNotesByExpenseId)
	mux.Post("/notes/get-all-expense-notes-by-user-id", app.GetAllExpenseNotesByUserId)

	mux.Post("/notes/create-external-company-note", app.CreateExternalCompanyNote)
	mux.Get("/notes/get-external-company-note-by-id", app.GetExternalCompanyNoteById)
	mux.Post("/notes/update-external-company-note", app.UpdateExternalCompanyNote)
	mux.Post("/notes/delete-external-company-note", app.DeleteExternalCompanyNote)
	mux.Post("/notes/get-all-external-company-notes-by-external-company-id", app.GetAllExternalCompanyNotesByExternalCompanyId)
	mux.Post("/notes/get-all-external-company-notes-by-user-id", app.GetAllExternalCompanyNotesByUserId)

	return mux
}
