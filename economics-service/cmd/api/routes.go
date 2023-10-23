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

	mux.Post("/economics/create-expense", app.CreateExpense)
	mux.Get("/economics/get-all-expenses", app.GetAllExpenses)
	mux.Post("/economics/update-expense", app.UpdateExpense)
	mux.Post("/economics/get-all-expenses-by-project-id", app.GetAllExpensesByProjectId)
	mux.Post("/economics/get-expense-by-id", app.GetExpenseById)

	mux.Post("/economics/create-income", app.CreateIncome)
	mux.Get("/economics/get-all-incomes", app.GetAllIncomes)
	mux.Post("/economics/update-income", app.UpdateIncome)
	mux.Post("/economics/get-all-incomes-by-project-id", app.GetAllProjectIncomesByProjectId)
	mux.Post("/economics/get-income-by-id", app.GetIncomeById)

	return mux
}
