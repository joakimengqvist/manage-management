package main

import (
	"economics-service/cmd/data"
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) CreateIncome(w http.ResponseWriter, r *http.Request) {

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

	var requestPayload data.NewIncome

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

func (app *Config) UpdateIncome(w http.ResponseWriter, r *http.Request) {

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

	var income data.Income

	err = app.readJSON(w, r, &income)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = income.UpdateIncome(userId)
	if err != nil {
		app.errorJSON(w, errors.New("could not update income: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated income with Id: %s", fmt.Sprint(income.ID)),
		Data:    income,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllIncomes(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "economics_read")
	if err != nil {
		fmt.Println("GetAllExpenses - authenticated error", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		fmt.Println("GetAllProjectExpenses - not authenticated")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	incomes, err := data.GetAllIncomes()
	if err != nil {
		fmt.Println("GetAllExpenses - incomes error", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var incomesSlice []data.Income
	for _, incomePtr := range incomes {
		income := *incomePtr

		returnedSlice := data.Income{
			ID:             income.ID,
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
			UpdatedBy:      income.UpdatedBy,
			UpdatedAt:      income.UpdatedAt,
		}

		incomesSlice = append(incomesSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched incomes successfull",
		Data:    incomesSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
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

	var incomesSlice []data.Income
	for _, incomesPtr := range incomes {
		income := *incomesPtr

		returnedSlice := data.Income{
			ID:             income.ID,
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
			UpdatedBy:      income.UpdatedBy,
			UpdatedAt:      income.UpdatedAt,
		}

		incomesSlice = append(incomesSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched incomes by project id successfull: %s", requestPayload.ID),
		Data:    incomesSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// --- END OF GET ALL PROJECT INCOMES (project ID) ----
// ----------------------------------------------------

// ----------------------------------------------------
// -- START OF GET INCOME (ID) ------------------------
// ----------------------------------------------------

func (app *Config) GetIncomeById(w http.ResponseWriter, r *http.Request) {
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

	returnedUser := data.Income{
		ID:             income.ID,
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
		UpdatedBy:      income.UpdatedBy,
		UpdatedAt:      income.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched income by id successfull: %s", income.ID),
		Data:    returnedUser,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// --------------------------------------------
// -- END OF GET INCOME (ID) ------------------
// --------------------------------------------
