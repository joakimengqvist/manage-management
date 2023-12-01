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
	mux.Post("/notes/get-all-project-notes-by-project-id", app.GetAllProjectNotesByProjectId)
	mux.Post("/notes/get-all-project-notes-by-user-id", app.GetAllProjectNotesByUserId)

	mux.Post("/notes/create-sub-project-note", app.CreateSubProjectNote)
	mux.Get("/notes/get-sub-project-note-by-id", app.GetSubProjectNoteById)
	mux.Post("/notes/update-sub-project-note", app.UpdateSubProjectNote)
	mux.Post("/notes/delete-sub-project-note", app.DeleteSubProjectNote)
	mux.Post("/notes/get-all-sub-project-notes-by-sub-project-id", app.GetAllSubProjectNotesBySubProjectId)
	mux.Post("/notes/get-all-sub-project-notes-by-user-id", app.GetAllSubProjectNotesByUserId)

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

	mux.Post("/notes/create-product-note", app.CreateProductNote)
	mux.Get("/notes/get-product-note-by-id", app.GetProductNoteById)
	mux.Post("/notes/update-product-note", app.UpdateProductNote)
	mux.Post("/notes/get-all-product-notes-by-product-id", app.GetAllProductNotesByProductId)
	mux.Post("/notes/get-all-product-notes-by-user-id", app.GetAllProductNotesByUserId)
	mux.Post("/notes/delete-product-note", app.DeleteProductNote)

	mux.Post("/notes/create-invoice-note", app.CreateInvoiceNote)
	mux.Get("/notes/get-invoice-note-by-id", app.GetInvoiceNoteById)
	mux.Post("/notes/update-invoice-note", app.UpdateInvoiceNote)
	mux.Post("/notes/get-all-invoice-notes-by-invoice-id", app.GetAllInvoiceNotesByInvoiceId)
	mux.Post("/notes/get-all-invoice-notes-by-user-id", app.GetAllInvoiceNotesByUserId)
	mux.Post("/notes/delete-invoice-note", app.DeleteInvoiceNote)

	mux.Post("/notes/create-invoice-item-note", app.CreateInvoiceItemNote)
	mux.Get("/notes/get-invoice-item-note-by-id", app.GetInvoiceItemNoteById)
	mux.Post("/notes/update-invoice-item-note", app.UpdateInvoiceItemNote)
	mux.Post("/notes/get-all-invoice-item-notes-by-invoice-item-id", app.GetAllInvoiceItemNotesByInvoiceItemId)
	mux.Post("/notes/get-all-invoice-item-notes-by-user-id", app.GetAllInvoiceItemNotesByUserId)
	mux.Post("/notes/delete-invoice-item-note", app.DeleteInvoiceItemNote)

	return mux
}
