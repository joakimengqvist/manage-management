package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
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
}

type NewIncome struct {
	ProjectID        string    `json:"project_id"`
	InvoiceID        string    `json:"invoice_id,omitempty"`
	IncomeDate       time.Time `json:"income_date"`
	IncomeCategory   string    `json:"income_category"`
	StatisticsIncome bool      `json:"statistics_income"`
	Vendor           string    `json:"vendor"`
	Description      string    `json:"description"`
	Amount           float64   `json:"amount"`
	Tax              float64   `json:"tax"`
	Status           string    `json:"status"`
	Currency         string    `json:"currency"`
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
	ID               string    `json:"id"`
	InvoiceID        string    `json:"invoice_id,omitempty"`
	ProjectID        string    `json:"project_id"`
	IncomeDate       time.Time `json:"income_date"`
	IncomeCategory   string    `json:"income_category"`
	StatisticsIncome bool      `json:"statistics_income"`
	Vendor           string    `json:"vendor"`
	Description      string    `json:"description"`
	Amount           float64   `json:"amount"`
	Tax              float64   `json:"tax"`
	Status           string    `json:"status"`
	Currency         string    `json:"currency"`
	CreatedBy        string    `json:"created_by"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedBy        string    `json:"updated_by"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (app *Config) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewExpense
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateExpense", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/create-expense"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - CreateExpense", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - CreateExpense", err)
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
		log.Println("json.NewDecoder - CreateExpense", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) CreateIncome(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewIncome
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateIncome", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/create-income"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - CreateIncome", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - CreateIncome", err)
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
		log.Println("json.NewDecoder - CreateIncome", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllExpensesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllExpensesByProjectId", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/get-all-expenses-by-project-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetAllExpensesByProjectId", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetAllExpensesByProjectId", err)
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
		log.Println("json.NewDecoder - GetAllExpensesByProjectId", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllIncomesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllIncomesByProjectId", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/get-all-incomes-by-project-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetAllIncomesByProjectId", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetAllIncomesByProjectId", err)
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
		log.Println("json.NewDecoder - GetAllIncomesByProjectId", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllExpenses(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/get-all-expenses"

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println("GET - GetAllExpenses", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetAllExpenses", err)
		app.errorJSON(w, errors.New("could not fetch expenses"))
		return
	}

	defer response.Body.Close()

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - GetAllExpenses", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllIncomes(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/get-all-incomes"

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println("GET - GetAllIncomes", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetAllIncomes", err)
		app.errorJSON(w, errors.New("could not fetch incomes"))
		return
	}

	defer response.Body.Close()

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - GetAllIncomes", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	var requestPayload Expense
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateExpense", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/update-expense"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - UpdateExpense", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - UpdateExpense", err)
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
		log.Println("json.NewDecoder - UpdateExpense", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) UpdateIncome(w http.ResponseWriter, r *http.Request) {
	var requestPayload Income
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateIncome", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/update-income"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - UpdateIncome", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - UpdateIncome", err)
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
		log.Println("json.NewDecoder - UpdateIncome", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetExpenseById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetExpenseById", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/get-expense-by-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetExpenseById", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetExpenseById", err)
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
		log.Println("json.NewDecoder - GetExpenseById", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetIncomeById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetIncomeById", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("ECONOMICS_SERVICE_SERVICE_HOST") + "/economics/get-income-by-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetIncomeById", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetIncomeById", err)
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
		log.Println("json.NewDecoder - GetIncomeById", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}
