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
	mux.Post("/project/get-projects-by-ids", app.GetProjectsByIds)
	mux.Get("/project/get-all-projects", app.GetAllProjects)

	mux.Post("/project/create-sub-project", app.CreateSubProject)
	mux.Post("/project/update-sub-project", app.UpdateSubProject)
	mux.Post("/project/delete-sub-project", app.DeleteSubProject)
	mux.Post("/project/get-sub-project-by-id", app.GetSubProjectById)
	mux.Post("/project/get-sub-projects-by-ids", app.GetSubProjectsByIds)
	mux.Get("/project/get-all-sub-projects", app.GetAllSubProjects)

	mux.Post("/project/add-projects-sub-project-connection", app.AddProjectsSubProjectConnection)
	mux.Post("/project/delete-projects-sub-project-connection", app.RemoveProjectsSubProjectConnection)
	mux.Post("/project/add-sub-projects-project-connection", app.AddSubProjectsSubProjectConnection)
	mux.Post("/project/delete-sub-projects-project-connection", app.RemoveSubProjectsProjectConnection)

	mux.Post("/notes/create-project-note", app.CreateProjectNote)
	mux.Get("/notes/get-project-note-by-id", app.GetProjectNoteById)
	mux.Post("/notes/update-project-note", app.UpdateProjectNote)
	mux.Post("/notes/get-all-project-notes-by-project-id", app.GetAllProjectNotesByProjectId)
	mux.Post("/notes/get-all-project-notes-by-user-id", app.GetAllProjectNotesByUserId)
	mux.Post("/notes/delete-project-note", app.DeleteProjectNote)

	mux.Post("/notes/create-sub-project-note", app.CreateSubProjectNote)
	mux.Get("/notes/get-sub-project-note-by-id", app.GetSubProjectNoteById)
	mux.Post("/notes/update-sub-project-note", app.UpdateSubProjectNote)
	mux.Post("/notes/delete-sub-project-note", app.DeleteSubProjectNote)
	mux.Post("/notes/get-all-sub-project-notes-by-sub-project-id", app.GetAllSubProjectNotesBySubProjectId)
	mux.Post("/notes/get-all-sub-project-notes-by-user-id", app.GetAllSubProjectNotesByUserId)

	mux.Post("/notes/create-income-note", app.CreateIncomeNote)
	mux.Get("/notes/get-income-note-by-id", app.GetIncomeNoteById)
	mux.Post("/notes/update-income-note", app.UpdateIncomeNote)
	mux.Post("/notes/get-all-income-notes-by-income-id", app.GetAllIncomeNotesByIncomeId)
	mux.Post("/notes/get-all-income-notes-by-user-id", app.GetAllIncomeNotesByUserId)
	mux.Post("/notes/delete-income-note", app.DeleteIncomeNote)

	mux.Post("/notes/create-expense-note", app.CreateExpenseNote)
	mux.Get("/notes/get-expense-note-by-id", app.GetExpenseNoteById)
	mux.Post("/notes/update-expense-note", app.UpdateExpenseNote)
	mux.Post("/notes/get-all-expense-notes-by-expense-id", app.GetAllExpenseNotesByExpenseId)
	mux.Post("/notes/get-all-expense-notes-by-user-id", app.GetAllExpenseNotesByUserId)
	mux.Post("/notes/delete-expense-note", app.DeleteExpenseNote)

	mux.Post("/notes/create-external-company-note", app.CreateExternalCompanyNote)
	mux.Get("/notes/get-external-company-note-by-id", app.GetExternalCompanyNoteById)
	mux.Post("/notes/update-external-company-note", app.UpdateExternalCompanyNote)
	mux.Post("/notes/delete-external-company-note", app.DeleteExternalCompanyNote)
	mux.Post("/notes/get-all-external-company-notes-by-external-company-id", app.GetAllExternalCompanyNotesByExternalCompanyId)
	mux.Post("/notes/get-all-external-company-notes-by-user-id", app.GetAllExternalCompanyNotesByUserId)

	mux.Post("/economics/create-expense", app.CreateExpense)
	mux.Get("/economics/get-all-expenses", app.GetAllExpenses)
	mux.Post("/economics/update-expense", app.UpdateExpense)
	mux.Post("/economics/get-expense-by-id", app.GetExpenseById)
	mux.Post("/economics/get-all-expenses-by-project-id", app.GetAllExpensesByProjectId)

	mux.Post("/economics/create-income", app.CreateIncome)
	mux.Get("/economics/get-all-incomes", app.GetAllIncomes)
	mux.Post("/economics/update-income", app.UpdateIncome)
	mux.Post("/economics/get-income-by-id", app.GetIncomeById)
	mux.Post("/economics/get-all-incomes-by-project-id", app.GetAllIncomesByProjectId)

	mux.Post("/external-company/create-external-company", app.CreateExternalCompany)
	mux.Get("/external-company/get-all-external-companies", app.GetAllExternalCompanies)
	mux.Post("/external-company/get-external-company-by-id", app.GetExternalCompanyById)
	mux.Post("/external-company/update-external-company", app.UpdateExternalCompany)

	mux.Post("/email/send", app.SendEmail)

	return mux
}
