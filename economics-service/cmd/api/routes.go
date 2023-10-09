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
	mux.Post("/economics/get-all-project-incomes-by-project-id", app.GetAllProjectIncomesByProjectId)
	mux.Post("/economics/get-project-income-by-id", app.GetProjectIncomeById)

	return mux
}

/*

func (app *Config) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateUser

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "user_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	updatedUser := data.User{
		ID:         requestPayload.ID,
		Email:      requestPayload.Email,
		FirstName:  requestPayload.FirstName,
		Privileges: app.convertToPostgresArray(requestPayload.Privileges),
		Projects:   app.convertToPostgresArray(requestPayload.Projects),
		LastName:   requestPayload.LastName,
	}

	err = updatedUser.UpdateUser()
	if err != nil {
		app.errorJSON(w, errors.New("could not update user: "+err.Error()), http.StatusBadRequest)
		return
	}

	returnedData := data.ReturnedUser{
		ID:         requestPayload.ID,
		Email:      requestPayload.Email,
		FirstName:  requestPayload.FirstName,
		Privileges: requestPayload.Privileges,
		Projects:   requestPayload.Projects,
		LastName:   requestPayload.LastName,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated user with Id: %s", fmt.Sprint(updatedUser.ID)),
		Data:    returnedData,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Authenticate [/auth/update-user]", Name: "[authentication-service] - Successful updated user"})
	app.writeJSON(w, http.StatusAccepted, payload)
}


*/
