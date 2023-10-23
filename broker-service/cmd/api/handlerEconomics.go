package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type NewExpense struct {
	ProjectID       string    `json:"project_id"`
	ExpenseDate     time.Time `json:"expense_date"`
	ExpenseCategory string    `json:"expense_category"`
	Vendor          string    `json:"vendor"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Tax             float64   `json:"tax"`
	Status          string    `json:"status"`
	Currency        string    `json:"currency"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedBy       string    `json:"created_by"`
	UpdatedBy       string    `json:"updated_by"`
}

type NewIncome struct {
	ProjectID      string    `json:"project_id"`
	IncomeDate     time.Time `json:"income_date"`
	IncomeCategory string    `json:"income_category"`
	Vendor         string    `json:"vendor"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Tax            float64   `json:"tax"`
	Status         string    `json:"status"`
	Currency       string    `json:"currency"`
	PaymentMethod  string    `json:"payment_method"`
	CreatedBy      string    `json:"created_by"`
	UpdatedBy      string    `json:"updated_by"`
}

type Expense struct {
	ID              string    `json:"id"`
	ProjectID       string    `json:"project_id"`
	ExpenseDate     time.Time `json:"expense_date"`
	ExpenseCategory string    `json:"expense_category"`
	Vendor          string    `json:"vendor"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Tax             float64   `json:"tax"`
	Status          string    `json:"status"`
	Currency        string    `json:"currency"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedBy       string    `json:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Income struct {
	ID             string    `json:"id"`
	ProjectID      string    `json:"project_id"`
	IncomeDate     time.Time `json:"income_date"`
	IncomeCategory string    `json:"income_category"`
	Vendor         string    `json:"vendor"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Tax            float64   `json:"tax"`
	Status         string    `json:"status"`
	Currency       string    `json:"currency"`
	PaymentMethod  string    `json:"payment_method"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// -------------------------------------------
// ------ START OF CREATE EXPENSE  -----------
// -------------------------------------------

func (app *Config) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewExpense
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/create-expense", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - create expense"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economic service - create expense"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// ------ END OF CREATE EXPENSE  -------------
// -------------------------------------------

// -------------------------------------------
// ------ START OF CREATE INCOME  ------------
// -------------------------------------------

func (app *Config) CreateIncome(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewIncome
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/create-income", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - create income"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - create income"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// ------ END OF CREATE INCOME  --------------
// -------------------------------------------

// --------------------------------------------------------
// ------ START OF GET EXPENSES BY PROJECT ID  ------------
// --------------------------------------------------------

func (app *Config) GetAllExpensesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/get-all-expenses-by-project-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch expense"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get expense by project id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - get expense by project id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// --------------------------------------------------------
// ------ END OF GET EXPENSES BY PROJECT ID  --------------
// --------------------------------------------------------

// --------------------------------------------------------
// ------ START OF GET EXPENSES BY PROJECT ID  ------------
// --------------------------------------------------------

func (app *Config) GetAllIncomesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/get-all-incomes-by-project-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch incomes"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get incomes by project id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - get incomes by project id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// --------------------------------------------------------
// ------ END OF GET EXPENSES BY PROJECT ID  --------------
// --------------------------------------------------------

// -------------------------------------------------
// ------ START OF GET ALL EXPENSES  ---------------
// -------------------------------------------------

func (app *Config) GetAllExpenses(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	request, err := http.NewRequest("GET", "http://economics-service/economics/get-all-expenses", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch expenses"))
		return
	}

	defer response.Body.Close()

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------------
// ------ END OF GET ALL EXPENSES  -----------------
// -------------------------------------------------

// -------------------------------------------------
// ------ START OF GET ALL INCOMES  ----------------
// -------------------------------------------------

func (app *Config) GetAllIncomes(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	request, err := http.NewRequest("GET", "http://economics-service/economics/get-all-incomes", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch incomes"))
		return
	}

	defer response.Body.Close()

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------------
// ------ END OF GET ALL PROJECT INCOMES  ----------
// -------------------------------------------------

// -------------------------------------------------
// ------ START OF UPDATE PROJECT EXPENSE  ---------
// -------------------------------------------------

func (app *Config) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	var requestPayload Expense
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/update-expense", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - update expense"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - update expense"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// --------- END OF UPDATE EXPENSE  ----------
// -------------------------------------------

// -------------------------------------------
// ------ START OF UPDATE INCOME  ------------
// -------------------------------------------

func (app *Config) UpdateIncome(w http.ResponseWriter, r *http.Request) {
	var requestPayload Income
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/update-income", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - update income"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - update income"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// --------- END OF UPDATE INCOME  -----------
// -------------------------------------------

// -------------------------------------------------
// ------ START OF GET EXPENSE (ID)  ---------------
// -------------------------------------------------

func (app *Config) GetExpenseById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/get-expense-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch project expense"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get expense by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - get expense by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------------
// ------ END OF GET EXPENSE (ID)  -----------------
// -------------------------------------------------

// -------------------------------------------------
// ------ START OF GET INCOME (ID)  ----------------
// -------------------------------------------------

func (app *Config) GetIncomeById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/get-income-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch income"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get income by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - get income by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// ----------------------------------------
// ------ END OF GET INCOME (ID)  ---------
// ----------------------------------------
