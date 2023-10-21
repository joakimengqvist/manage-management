package main

import (
	"economics-service/cmd/data"
	"errors"
	"fmt"
	"net/http"
)

// ----------------------------------------------------
// --------- START OF CREATE PROJECT EXPENSE  ---------
// ----------------------------------------------------

func (app *Config) CreateProjectExpense(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "economics_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload data.NewProjectExpense

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	response, err := data.InsertExpense(requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("could not create expense: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created expense %s", requestPayload.Description),
		Data:    response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllProjectExpenses(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "economics_read")
	if err != nil {
		fmt.Println("GetAllProjectExpenses - authenticated error", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		fmt.Println("GetAllProjectExpenses - not authenticated")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	expenses, err := data.GetAllProjectExpenses()
	if err != nil {
		fmt.Println("GetAllProjectExpenses - expenses error", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var expensesSlice []data.ProjectExpense
	for _, expensesPtr := range expenses {
		expense := *expensesPtr

		returnedSlice := data.ProjectExpense{
			ExpenseID:       expense.ExpenseID,
			ProjectID:       expense.ProjectID,
			ExpenseDate:     expense.ExpenseDate,
			ExpenseCategory: expense.ExpenseCategory,
			Vendor:          expense.Vendor,
			Description:     expense.Description,
			Amount:          expense.Amount,
			Tax:             expense.Tax,
			Status:          expense.Status,
			Currency:        expense.Currency,
			PaymentMethod:   expense.PaymentMethod,
			CreatedBy:       expense.CreatedBy,
			CreatedAt:       expense.CreatedAt,
			ModifiedBy:      expense.ModifiedBy,
			ModifiedAt:      expense.ModifiedAt,
		}

		expensesSlice = append(expensesSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "fetched all exepnses",
		Data:    expensesSlice,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get all privileges [/auth/get-all-privileges]", Name: "[authentication-service] - Successfuly fetched all privileges"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// ---- END OF GET ALL PROJECT EXPENSES ---------------
// ----------------------------------------------------

// ----------------------------------------------------
// -- START OF GET ALL PROJECT EXPENSES (project ID) --
// ----------------------------------------------------

func (app *Config) GetAllProjectExpensesByProjectId(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "economics_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	expenses, err := data.GetAllProjectExpensesByProjectId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var expensesSlice []data.ProjectExpense
	for _, expensesPtr := range expenses {
		expense := *expensesPtr

		returnedSlice := data.ProjectExpense{
			ExpenseID:       expense.ExpenseID,
			ProjectID:       expense.ProjectID,
			ExpenseDate:     expense.ExpenseDate,
			ExpenseCategory: expense.ExpenseCategory,
			Vendor:          expense.Vendor,
			Description:     expense.Description,
			Amount:          expense.Amount,
			Tax:             expense.Tax,
			Status:          expense.Status,
			Currency:        expense.Currency,
			PaymentMethod:   expense.PaymentMethod,
			CreatedBy:       expense.CreatedBy,
			CreatedAt:       expense.CreatedAt,
			ModifiedBy:      expense.ModifiedBy,
			ModifiedAt:      expense.ModifiedAt,
		}

		expensesSlice = append(expensesSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all expenses by project id",
		Data:    expensesSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// --- END OF GET ALL PROJECT EXPENSES (project ID) ---
// ----------------------------------------------------

// ----------------------------------------------------
// -- START OF GET PROJECT EXPENSE (ID) ---------------
// ----------------------------------------------------

func (app *Config) GetProjectExpenseById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "economics_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	expense, err := data.GetExpenseById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get expense by id"), http.StatusBadRequest)
		return
	}

	returnedUser := data.ProjectExpense{
		ExpenseID:       expense.ExpenseID,
		ProjectID:       expense.ProjectID,
		ExpenseDate:     expense.ExpenseDate,
		ExpenseCategory: expense.ExpenseCategory,
		Vendor:          expense.Vendor,
		Description:     expense.Description,
		Amount:          expense.Amount,
		Tax:             expense.Tax,
		Status:          expense.Status,
		Currency:        expense.Currency,
		PaymentMethod:   expense.PaymentMethod,
		CreatedBy:       expense.CreatedBy,
		CreatedAt:       expense.CreatedAt,
		ModifiedBy:      expense.ModifiedBy,
		ModifiedAt:      expense.ModifiedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched expense by id successfull: %s", expense.ExpenseID),
		Data:    returnedUser,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// -- END OF GET PROJECT EXPENSE (ID) -----------------
// ----------------------------------------------------
