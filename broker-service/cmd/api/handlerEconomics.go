package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type NewProjectExpense struct {
	ProjectID       string    `json:"project_id"`
	ExpenseDate     time.Time `json:"expense_date"`
	ExpenseCategory string    `json:"expense_category"`
	Vendor          string    `json:"vendor"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Tax             float64   `json:"tax"`
	Currency        string    `json:"currency"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedBy       string    `json:"created_by"`
	ModifiedBy      string    `json:"modified_by"`
}

type NewProjectIncome struct {
	ProjectID      string    `json:"project_id"`
	IncomeDate     time.Time `json:"income_date"`
	IncomeCategory string    `json:"income_category"`
	Vendor         string    `json:"vendor"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Tax            float64   `json:"tax"`
	Currency       string    `json:"currency"`
	PaymentMethod  string    `json:"payment_method"`
	CreatedBy      string    `json:"created_by"`
	ModifiedBy     string    `json:"modified_by"`
}

type ProjectExpense struct {
	ExpenseID       string    `json:"expense_id"`
	ProjectID       string    `json:"project_id"`
	ExpenseDate     time.Time `json:"expense_date"`
	ExpenseCategory string    `json:"expense_category"`
	Vendor          string    `json:"vendor"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Tax             float64   `json:"tax"`
	Currency        string    `json:"currency"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	ModifiedBy      string    `json:"modified_by"`
	ModifiedAt      time.Time `json:"modified_at"`
}

type ProjectIncome struct {
	ExpenseID      string    `json:"income_id"`
	ProjectID      string    `json:"project_id"`
	IncomeDate     time.Time `json:"income_date"`
	IncomeCategory string    `json:"income_category"`
	Vendor         string    `json:"vendor"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Tax            float64   `json:"tax"`
	Currency       string    `json:"currency"`
	PaymentMethod  string    `json:"payment_method"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	ModifiedBy     string    `json:"modified_by"`
	ModifiedAt     time.Time `json:"modified_at"`
}

type ProjectId struct {
	ProjectId string `json:"project_id"`
}

type ExpenseId struct {
	ExpenseId string `json:"expense_id"`
}

type IncomeId struct {
	IncomeId string `json:"income_id"`
}

// -------------------------------------------
// ------ START OF CREATE PROJECT EXPENSE  ---
// -------------------------------------------

func (app *Config) CreateProjectExpense(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewProjectExpense
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/create-project-expense", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - create project expense"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economic service - create project expense"))
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

	var payload jsonResponse
	payload.Error = false
	payload.Message = "create project expense successful"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// ------ END OF CREATE PROJECT EXPENSE  -----
// -------------------------------------------

// -------------------------------------------
// ------ START OF CREATE PROJECT INCOME  ----
// -------------------------------------------

func (app *Config) CreateProjectIncome(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewProjectIncome
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/create-project-income", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - create project income"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - create project income"))
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

	var payload jsonResponse
	payload.Error = false
	payload.Message = "create project income successful"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// ------ END OF CREATE PROJECT INCOME  ------
// -------------------------------------------

// --------------------------------------------------------
// ------ START OF GET PROJECT EXPENSES BY PROJECT ID  ----
// --------------------------------------------------------

func (app *Config) GetAllProjectExpensesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectId

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/get-all-project-expenses-by-project-id", bytes.NewBuffer(jsonData))
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

	var jsonFromService []ProjectExpense

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "get expense by project id successful"
	payload.Data = jsonFromService

	app.writeJSON(w, http.StatusAccepted, payload)
}

// --------------------------------------------------------
// ------ END OF GET PROJECT EXPENSES BY PROJECT ID  ------
// --------------------------------------------------------

// --------------------------------------------------------
// ------ START OF GET PROJECT EXPENSES BY PROJECT ID  ----
// --------------------------------------------------------

func (app *Config) GetAllProjectIncomesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectId

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/get-all-project-incomes-by-project-id", bytes.NewBuffer(jsonData))
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

	var jsonFromService []ProjectIncome

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "get incomes by project id successful"
	payload.Data = jsonFromService

	app.writeJSON(w, http.StatusAccepted, payload)
}

// --------------------------------------------------------
// ------ END OF GET PROJECT EXPENSES BY PROJECT ID  ------
// --------------------------------------------------------

// -------------------------------------------------
// ------ START OF GET ALL PROJECT EXPENSES  -------
// -------------------------------------------------

func (app *Config) GetAllProjectExpenses(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	request, err := http.NewRequest("GET", "http://economics-service/economics/get-all-project-expenses", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch project expenses"))
		return
	}

	defer response.Body.Close()

	var jsonFromService []ProjectExpense

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------------
// ------ END OF GET ALL PROJECT EXPENSES  ---------
// -------------------------------------------------

// -------------------------------------------------
// ------ START OF GET ALL PROJECT INCOMES  -------
// -------------------------------------------------

func (app *Config) GetAllProjectIncomes(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	request, err := http.NewRequest("GET", "http://economics-service/economics/get-all-project-incomes", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch project incomes"))
		return
	}

	defer response.Body.Close()

	var jsonFromService []ProjectIncome

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
// ------ START OF GET PROJECT EXPENSE (ID)  -------
// -------------------------------------------------

func (app *Config) GetProjectExpenseById(w http.ResponseWriter, r *http.Request) {
	var requestPayload ExpenseId

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/get-project-expense-by-id", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - get project expense by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - get project expense by id"))
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

	var payload jsonResponse
	payload.Error = false
	payload.Message = "get project expsense by id successful"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------------
// ------ END OF GET PROJECT EXPENSE (ID)  ---------
// -------------------------------------------------

// -------------------------------------------------
// ------ START OF GET PROJECT INCOME (ID)  -------
// -------------------------------------------------

func (app *Config) GetProjectIncomeById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IncomeId

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://economics-service/economics/get-project-income-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch project income"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get project income by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - get project income by id"))
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

	var payload jsonResponse
	payload.Error = false
	payload.Message = "get project income by id successful"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------------
// ------ END OF GET PROJECT INCOME (ID)  ---------
// -------------------------------------------------
