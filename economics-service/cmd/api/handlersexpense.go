package main

import (
	"economics-service/cmd/data"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *Config) CreateExpense(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "economics_write")
	if err != nil {
		log.Println("authenticated - CreateExpense", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - CreateExpense")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload data.NewExpense

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateExpense", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	response, err := data.InsertExpense(requestPayload, userId)
	if err != nil {
		log.Println("InsertExpense - CreateExpense", err)
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
		log.Println("authenticated - UpdateExpense", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - UpdateExpense")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var expense data.Expense

	err = app.readJSON(w, r, &expense)
	if err != nil {
		log.Println("readJSON - UpdateExpense", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = expense.UpdateExpense(userId)
	if err != nil {
		log.Println("postgres - UpdateExpense", err)
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
		log.Println("authenticated - GetAllExpenses", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("GetAllExpenses - !authenticated")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	expenses, err := data.GetAllExpenses()
	if err != nil {
		log.Println("postgres - GetAllProjectExpenses", err)
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

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllExpensesByProjectId(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "economics_read")
	if err != nil {
		log.Println("authenticated - GetAllExpensesByProjectId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllExpensesByProjectId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllExpensesByProjectId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	expenses, err := data.GetAllExpensesByProjectId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetAllExpensesByProjectId", err)
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
		log.Println("authenticated - GetExpenseById", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetExpenseById")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetExpenseById", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	expense, err := data.GetExpenseById(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetExpenseById", err)
		app.errorJSON(w, errors.New("failed to get expense by id"), http.StatusBadRequest)
		return
	}

	log.Println(expense)

	returnedExpense := data.Expense{
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
		Data:    returnedExpense,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
