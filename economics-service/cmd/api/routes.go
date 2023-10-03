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

	mux.Post("/economics/create-project-expense", app.CreateProjectExpense)
	mux.Get("/economics/get-all-project-expenses", app.GetAllProjectExpenses)
	mux.Post("/economics/get-all-project-expenses-by-project-id", app.GetAllProjectExpensesByProjectId)
	mux.Post("/economics/get-project-expense-by-id", app.GetProjectExpenseById)

	mux.Post("/economics/create-project-income", app.CreateProjectIncome)
	mux.Get("/economics/get-all-project-incomes", app.GetAllProjectIncomes)
	mux.Post("/economics/get-project-income-by-id", app.GetProjectIncomeById)
	mux.Post("/economics/get-all-project-incomes-by-project-id", app.GetAllProjectIncomesByProjectId)

	return mux
}
