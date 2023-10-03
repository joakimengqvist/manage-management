package main

import (
	"economics-service/cmd/data"
	"errors"
	"fmt"
	"net/http"
)

type ProjectId struct {
	ProjectId string `json:"project_id"`
}

type ExpenseId struct {
	ExpenseId string `json:"expense_id"`
}

type IncomeId struct {
	IncomeId string `json:"income_id"`
}

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

// ----------------------------------------------------
// --------- END OF CREATE PROJECT EXPENSE  -----------
// ----------------------------------------------------

// ----------------------------------------------------
// --------- START OF CREATE PROJECT INCOME  ----------
// ----------------------------------------------------

func (app *Config) CreateProjectIncome(w http.ResponseWriter, r *http.Request) {

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

	var requestPayload data.NewProjectIncome

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	response, err := data.InsertIncome(requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("could not create income: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created income %s", requestPayload.Description),
		Data:    response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// ---- END OF CREATE PROJECT INCOME  -----------------
// ----------------------------------------------------

// ----------------------------------------------------
// ---- START OF GET ALL PROJECT EXPENSES -------------
// ----------------------------------------------------

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
			Currency:        expense.Currency,
			PaymentMethod:   expense.PaymentMethod,
			CreatedBy:       expense.CreatedBy,
			CreatedAt:       expense.CreatedAt,
			ModifiedBy:      expense.ModifiedBy,
			ModifiedAt:      expense.ModifiedAt,
		}

		expensesSlice = append(expensesSlice, returnedSlice)
	}

	app.logItemViaRPC(w, expensesSlice, RPCLogData{Action: "Get all privileges [/auth/get-all-privileges]", Name: "[authentication-service] - Successfuly fetched all privileges"})
	app.writeExpensesJSONFromSlice(w, http.StatusAccepted, expensesSlice)
}

// ----------------------------------------------------
// ---- END OF GET ALL PROJECT EXPENSES ---------------
// ----------------------------------------------------

// ----------------------------------------------------
// ---- START OF GET ALL PROJECT INCOMES --------------
// ----------------------------------------------------

func (app *Config) GetAllProjectIncomes(w http.ResponseWriter, r *http.Request) {

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

	incomes, err := data.GetAllProjectIncomes()
	if err != nil {
		fmt.Println("GetAllProjectExpenses - expenses error", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var incomesSlice []data.ProjectIncome
	for _, incomePtr := range incomes {
		income := *incomePtr

		returnedSlice := data.ProjectIncome{
			IncomeID:       income.IncomeID,
			ProjectID:      income.ProjectID,
			IncomeDate:     income.IncomeDate,
			IncomeCategory: income.IncomeCategory,
			Vendor:         income.Vendor,
			Description:    income.Description,
			Amount:         income.Amount,
			Tax:            income.Tax,
			Currency:       income.Currency,
			PaymentMethod:  income.PaymentMethod,
			CreatedBy:      income.CreatedBy,
			CreatedAt:      income.CreatedAt,
			ModifiedBy:     income.ModifiedBy,
			ModifiedAt:     income.ModifiedAt,
		}

		incomesSlice = append(incomesSlice, returnedSlice)
	}

	app.writeIncomesJSONFromSlice(w, http.StatusAccepted, incomesSlice)
}

// ----------------------------------------------------
// ---- END OF GET ALL PROJECT INCOMES ----------------
// ----------------------------------------------------

// ----------------------------------------------------
// -- START OF GET ALL PROJECT EXPENSES (project ID) --
// ----------------------------------------------------

func (app *Config) GetAllProjectExpensesByProjectId(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "economics_read")
	if err != nil {
		fmt.Println("GetAllProjectExpensesByProjectId - authenticated error", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		fmt.Println("GetAllProjectExpensesByProjectId - not authenticated")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload ProjectId

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Println("GetAllProjectExpensesByProjectId - readJSON error", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	expenses, err := data.GetAllProjectExpensesByProjectId(requestPayload.ProjectId)
	if err != nil {
		fmt.Println("GetAllProjectExpensesByProjectId - expenses error", err)
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
			Currency:        expense.Currency,
			PaymentMethod:   expense.PaymentMethod,
			CreatedBy:       expense.CreatedBy,
			CreatedAt:       expense.CreatedAt,
			ModifiedBy:      expense.ModifiedBy,
			ModifiedAt:      expense.ModifiedAt,
		}

		expensesSlice = append(expensesSlice, returnedSlice)
	}

	app.writeExpensesJSONFromSlice(w, http.StatusAccepted, expensesSlice)
}

// ----------------------------------------------------
// --- END OF GET ALL PROJECT EXPENSES (project ID) ---
// ----------------------------------------------------

// ----------------------------------------------------
// -- START OF GET ALL PROJECT INCOMES (project ID) ---
// ----------------------------------------------------

func (app *Config) GetAllProjectIncomesByProjectId(w http.ResponseWriter, r *http.Request) {

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

	var requestPayload ProjectId

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	incomes, err := data.GetAllProjectIncomesByProjectId(requestPayload.ProjectId)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var incomesSlice []data.ProjectIncome
	for _, incomesPtr := range incomes {
		income := *incomesPtr

		returnedSlice := data.ProjectIncome{
			IncomeID:       income.IncomeID,
			ProjectID:      income.ProjectID,
			IncomeDate:     income.IncomeDate,
			IncomeCategory: income.IncomeCategory,
			Vendor:         income.Vendor,
			Description:    income.Description,
			Amount:         income.Amount,
			Tax:            income.Tax,
			Currency:       income.Currency,
			PaymentMethod:  income.PaymentMethod,
			CreatedBy:      income.CreatedBy,
			CreatedAt:      income.CreatedAt,
			ModifiedBy:     income.ModifiedBy,
			ModifiedAt:     income.ModifiedAt,
		}

		incomesSlice = append(incomesSlice, returnedSlice)
	}

	app.writeIncomesJSONFromSlice(w, http.StatusAccepted, incomesSlice)
}

// ----------------------------------------------------
// --- END OF GET ALL PROJECT INCOMES (project ID) ----
// ----------------------------------------------------

// ----------------------------------------------------
// -- START OF GET PROJECT EXPENSE (ID) ---------------
// ----------------------------------------------------

func (app *Config) GetProjectExpenseById(w http.ResponseWriter, r *http.Request) {

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

	var requestPayload ExpenseId

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	expense, err := data.GetExpenseById(requestPayload.ExpenseId)
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
		Currency:        expense.Currency,
		PaymentMethod:   expense.PaymentMethod,
		CreatedBy:       expense.CreatedBy,
		CreatedAt:       expense.CreatedAt,
		ModifiedBy:      expense.ModifiedBy,
		ModifiedAt:      expense.ModifiedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched expense successfull - ID: %s", expense.ExpenseID),
		Data:    returnedUser,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// -- END OF GET PROJECT EXPENSE (ID) -----------------
// ----------------------------------------------------

// ----------------------------------------------------
// -- START OF GET PROJECT INCOME (ID) ----------------
// ----------------------------------------------------

func (app *Config) GetProjectIncomeById(w http.ResponseWriter, r *http.Request) {

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

	var requestPayload IncomeId

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	income, err := data.GetIncomeById(requestPayload.IncomeId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get income by id"), http.StatusBadRequest)
		return
	}

	returnedUser := data.ProjectIncome{
		IncomeID:       income.IncomeID,
		ProjectID:      income.ProjectID,
		IncomeDate:     income.IncomeDate,
		IncomeCategory: income.IncomeCategory,
		Vendor:         income.Vendor,
		Description:    income.Description,
		Amount:         income.Amount,
		Tax:            income.Tax,
		Currency:       income.Currency,
		PaymentMethod:  income.PaymentMethod,
		CreatedBy:      income.CreatedBy,
		CreatedAt:      income.CreatedAt,
		ModifiedBy:     income.ModifiedBy,
		ModifiedAt:     income.ModifiedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched income successfull - ID: %s", income.IncomeID),
		Data:    returnedUser,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// -- END OF GET PROJECT INCOME (ID) ------------------
// ----------------------------------------------------
