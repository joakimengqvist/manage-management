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

	mux.Post("/notes/create-project-note", app.CreateProjectNote)
	mux.Get("/notes/get-project-note-by-id", app.GetProjectNoteById)
	mux.Post("/notes/update-project-note", app.UpdateProjectNote)
	mux.Post("/notes/get-all-notes-by-project-id", app.GetAllNotesByProductId)
	mux.Post("/notes/get-all-notes-by-user-id", app.GetAllNotesByUserId)
	mux.Post("/notes/delete-project-note-by-id", app.DeleteProjectNoteById)

	mux.Post("/economics/create-project-expense", app.CreateProjectExpense)
	mux.Get("/economics/get-all-project-expenses", app.GetAllProjectExpenses)
	mux.Get("/economics/update-project-expense", app.UpdateProjectExpense)
	mux.Post("/economics/get-project-expense-by-id", app.GetProjectExpenseById)
	mux.Post("/economics/get-all-project-expenses-by-project-id", app.GetAllProjectExpensesByProjectId)

	mux.Post("/economics/create-project-income", app.CreateProjectIncome)
	mux.Get("/economics/get-all-project-incomes", app.GetAllProjectIncomes)
	mux.Post("/economics/get-project-income-by-id", app.GetProjectIncomeById)
	mux.Post("/economics/get-all-project-incomes-by-project-id", app.GetAllProjectIncomesByProjectId)

	mux.Post("/external-company/create-external-company", app.CreateExternalCompany)
	mux.Get("/external-company/get-all-external-companies", app.GetAllExternalCompanies)
	mux.Post("/external-company/get-external-company-by-id", app.GetExternalCompanyById)

	mux.Post("/email/send", app.SendEmail)

	return mux
}
