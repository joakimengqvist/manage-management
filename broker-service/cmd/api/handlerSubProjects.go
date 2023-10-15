package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type SubProject struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Priority          int       `json:"priority"`
	StartDate         time.Time `json:"start_date"`
	DueDate           time.Time `json:"due_date"`
	EstimatedDuration int       `json:"estimated_duration"`
	Notes             []string  `json:"notes"`
	ProjectID         string    `json:"project_id"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedBy         string    `json:"updated_by"`
	Invoices          []string  `json:"invoices"`
	Incomes           []string  `json:"incomes"`
	Expenses          []string  `json:"expenses"`
}

type UpdateSubProject struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Priority          int       `json:"priority"`
	StartDate         time.Time `json:"start_date"`
	DueDate           time.Time `json:"due_date"`
	EstimatedDuration int       `json:"estimated_duration"`
	Notes             []string  `json:"notes"`
	ProjectID         string    `json:"project_id"`
	Invoices          []string  `json:"invoices"`
	Incomes           []string  `json:"incomes"`
	Expenses          []string  `json:"expenses"`
}

type NewSubProject struct {
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Priority          int       `json:"priority"`
	StartDate         time.Time `json:"start_date"`
	DueDate           time.Time `json:"due_date"`
	EstimatedDuration int       `json:"estimated_duration"`
	Notes             []string  `json:"notes"`
	ProjectID         string    `json:"project_id"`
	CreatedBy         string    `json:"created_by"`
	UpdatedBy         string    `json:"updated_by"`
	Invoices          []string  `json:"invoices"`
	Incomes           []string  `json:"incomes"`
	Expenses          []string  `json:"expenses"`
}

// -------------------------------------------
// --------- START OF CREATE PROJECT  --------
// -------------------------------------------

func (app *Config) CreateSubProject(w http.ResponseWriter, r *http.Request) {

	var requestPayload NewSubProject
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create subProject [/project/create-sub-project]", Name: "[broker-service] - Create subProject request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/create-sub-project", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - create subProject"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling subProject service - create subProject"))
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
	payload.Message = "create subProject successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Create subProject successfully [/project/create-sub-project]", Name: "[broker-service]"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF CREATE PROJECT  ----------
// -------------------------------------------

// -------------------------------------------
// --------- START OF UPDATE PROJECT  --------
// -------------------------------------------

func (app *Config) UpdateSubProject(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateSubProject
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Update subProject [/project/update-sub-project]", Name: "[broker-service] - Update subProject request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/update-sub-project", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - update subProject"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - update subProject"))
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
	payload.Message = "update subProject successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Updated subProject successfully [/project/update-sub-project]", Name: "[broker-service] - Successfully updated subProject"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF UPDATE PROJECT  ----------
// -------------------------------------------

// -------------------------------------------
// --------- START OF DELETE PROJECTS  -------
// -------------------------------------------

func (app *Config) DeleteSubProject(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Delete subProject [/project/delete-sub-project]", Name: "[broker-service] - Delete subProject request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/delete-sub-project", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - delete subProject"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - delete subProject"))
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
	payload.Message = "delete subProject successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete subProject successfully [/project/delete-sub-project]", Name: "[broker-service] - Successfully deleteted subProject"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF DELETE PROJECT  ----------
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET PROJECTS (ID)  -----
// -------------------------------------------

func (app *Config) GetSubProjectById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get subProject by id [/project/get-sub-project-by-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/get-sub-project-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch subProject"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get subProject by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - get subProject by id"))
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
	payload.Message = "get subProject by id successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get subProject by id successfully [/project/get-sub-project-by-id]", Name: "[broker-service] - Successfully fetched subProject"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF GET PROJECTS (ID)  -------
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET PROJECTS  ----------
// -------------------------------------------

func (app *Config) GetAllSubProjects(w http.ResponseWriter, r *http.Request) {

	app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all subProjects [/project/get-all-sub-projects]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	request, err := http.NewRequest("GET", "http://project-service/project/get-all-sub-projects", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch subProjects"))
		return
	}

	defer response.Body.Close()

	var jsonFromService []SubProject

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get all subProjects success [/project/get-all-sub-projects]", Name: "[broker-service]"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// --------- END OF GET PROJECTS  ------------
// -------------------------------------------
