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

	mux.Post("/invoice/create-invoice", app.CreateInvoice)
	mux.Get("/invoice/get-all-invoices", app.GetAllInvoices)
	mux.Post("/invoice/update-invoice", app.UpdateInvoice)
	mux.Post("/invoice/get-invoice-by-id", app.GetInvoiceById)
	mux.Post("/invoice/get-all-invoices-by-project-id", app.GetAllInvoicesByProjectId)
	mux.Post("/invoice/get-all-invoices-by-sub-project-id", app.GetAllInvoicesBySubProjectId)
	mux.Post("/invoice/get-all-invoices-by-ids", app.GetAllInvoicesByIds)

	mux.Post("/invoice/create-invoice-item", app.CreateInvoiceItem)
	mux.Get("/invoice/get-all-invoice-items", app.GetAllInvoiceItems)
	mux.Post("/invoice/update-invoice-item", app.UpdateInvoiceItem)
	mux.Post("/invoice/get-invoice-item-by-id", app.GetInvoiceItemById)
	mux.Post("/invoice/get-all-invoice-items-by-ids", app.GetAllInvoiceItemsByIds)
	mux.Post("/invoice/get-all-invoice-items-by-product-id", app.GetAllInvoiceItemsByProductId)

	return mux
}
