package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type ExternalCompany struct {
	ID                        string    `json:"id"`
	CompanyName               string    `json:"company_name" sql:"not null"`
	CompanyRegistrationNumber string    `json:"company_registration_number"`
	ContactPerson             string    `json:"contact_person"`
	ContactEmail              string    `json:"contact_email"`
	ContactPhone              string    `json:"contact_phone"`
	Address                   string    `json:"address"`
	City                      string    `json:"city"`
	StateProvince             string    `json:"state_province"`
	Country                   string    `json:"country"`
	PostalCode                string    `json:"postal_code"`
	PaymentTerms              string    `json:"payment_terms"`
	BillingCurrency           string    `json:"billing_currency"`
	BankAccountInfo           string    `json:"bank_account_info"`
	TaxIdentificationNumber   string    `json:"tax_identification_number"`
	CreatedAt                 time.Time `json:"created_at,omitempty"`
	CreatedBy                 string    `json:"created_by,omitempty"`
	UpdatedAt                 time.Time `json:"updated_at,omitempty"`
	UpdatedBy                 string    `json:"updated_by,omitempty"`
	Status                    string    `json:"status"`
	AssignedProjects          []string  `json:"assigned_projects"`
	InvoicePending            []string  `json:"invoice_pending"`
	InvoiceHistory            []string  `json:"invoice_history"`
	ContractualAgreements     []string  `json:"contractual_agreements"`
}

func (app *Config) CreateExternalCompany(w http.ResponseWriter, r *http.Request) {
	var requestPayload ExternalCompany
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://external-company-service/create-external-company", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - create external company"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - create external company"))
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

func (app *Config) UpdateExternalCompany(w http.ResponseWriter, r *http.Request) {
	var requestPayload ExternalCompany
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://external-company-service/update-external-company", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - update external company"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - update external company"))
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

func (app *Config) GetAllExternalCompanies(w http.ResponseWriter, r *http.Request) {

	app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all external companies [external-company/get-all-external-companies]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	request, err := http.NewRequest("GET", "http://external-company-service/get-all-external-companies", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch external companies"))
		return
	}

	defer response.Body.Close()

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get all external companies success [external-company/get-all-external-companies]", Name: "[broker-service]"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetExternalCompanyById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://external-company-service/get-external-company-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch project external company"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get project external company by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - get project external company by id"))
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
