package main

import (
	"economics-service/cmd/data"
	"errors"
	"fmt"
	"net/http"
)

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
			Status:         income.Status,
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
// -- START OF GET ALL PROJECT INCOMES (project ID) ---
// ----------------------------------------------------

func (app *Config) GetAllProjectIncomesByProjectId(w http.ResponseWriter, r *http.Request) {
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

	incomes, err := data.GetAllProjectIncomesByProjectId(requestPayload.ID)
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
			Status:         income.Status,
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
// -- START OF GET PROJECT INCOME (ID) ----------------
// ----------------------------------------------------

func (app *Config) GetProjectIncomeById(w http.ResponseWriter, r *http.Request) {
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

	income, err := data.GetIncomeById(requestPayload.ID)
	if err != nil {
		fmt.Println("income err", err)
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
		Status:         income.Status,
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
