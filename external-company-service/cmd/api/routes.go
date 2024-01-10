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

	mux.Post("/create-external-company", app.CreateExternalCompany)
	mux.Get("/get-all-external-companies", app.GetAllExternalCompanies)
	mux.Post("/get-external-company-by-id", app.GetExternalCompanyById)
	mux.Post("/update-external-company", app.UpdateExternalCompany)

	mux.Post("/append-invoice", app.AddInvoiceToCompany)
	mux.Post("/remove-invoice", app.RemoveInvoiceToCompany)

	return mux
}
