package main

import (
	"economics-service/cmd/data"
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) CreateExpense(w http.ResponseWriter, r *http.Request) {

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

	var requestPayload data.NewExpense

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	response, err := data.InsertExpense(requestPayload, userId)
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

func (app *Config) UpdateExpense(w http.ResponseWriter, r *http.Request) {

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

	var expense data.Expense

	err = app.readJSON(w, r, &expense)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = expense.UpdateExpense(userId)
	if err != nil {
		app.errorJSON(w, errors.New("could not update expense: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated expense with Id: %s", fmt.Sprint(expense.ID)),
		Data:    expense,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllExpenses(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "economics_read")
	if err != nil {
		fmt.Println("GetAllExpenses - authenticated error", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		fmt.Println("GetAllExpenses - not authenticated")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	expenses, err := data.GetAllExpenses()
	if err != nil {
		fmt.Println("GetAllProjectExpenses - expenses error", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var expensesSlice []data.Expense
	for _, expensesPtr := range expenses {
		expense := *expensesPtr

		returnedSlice := data.Expense{
			ID:              expense.ID,
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
			UpdatedBy:       expense.UpdatedBy,
			UpdatedAt:       expense.UpdatedAt,
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

func (app *Config) GetAllExpensesByProjectId(w http.ResponseWriter, r *http.Request) {

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

	expenses, err := data.GetAllExpensesByProjectId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var expensesSlice []data.Expense
	for _, expensesPtr := range expenses {
		expense := *expensesPtr

		returnedSlice := data.Expense{
			ID:              expense.ID,
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
			UpdatedBy:       expense.UpdatedBy,
			UpdatedAt:       expense.UpdatedAt,
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

func (app *Config) GetExpenseById(w http.ResponseWriter, r *http.Request) {
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

	returnedUser := data.Expense{
		ID:              expense.ID,
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
		UpdatedBy:       expense.UpdatedBy,
		UpdatedAt:       expense.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched expense by id successfull: %s", expense.ID),
		Data:    returnedUser,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
