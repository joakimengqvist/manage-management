package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

// Project is the structure which holds one project from the database.

type ProjectIdPayload struct {
	Id int `json:"id"`
}

type Project struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewProject struct {
	Name string `json:"name"`
}

// -------------------------------------------
// --------- START OF CREATE PROJECT  --------
// -------------------------------------------

func (app *Config) CreateProject(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewProject
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create project [/project/create-project]", Name: "[broker-service] - Create project request recieved"})
	app.CreateProjectCall(w, requestPayload)
}

func (app *Config) CreateProjectCall(w http.ResponseWriter, requestPayload NewProject) {
	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/create-project", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - create project"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling project service - create project"))
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
	payload.Message = "create project successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create project successfully [/project/create-project]", Name: "[broker-service]"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF CREATE PROJECT  ----------
// -------------------------------------------

// -------------------------------------------
// --------- START OF UPDATE PROJECT  --------
// -------------------------------------------

type ProjectPayload struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (app *Config) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Update project [/project/update-project]", Name: "[broker-service] - Update project request recieved"})
	app.UpdateProjectCall(w, requestPayload)
}

func (app *Config) UpdateProjectCall(w http.ResponseWriter, requestPayload ProjectPayload) {
	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/update-project", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - update project"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - update project"))
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
	payload.Message = "update project successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Updated project successfully [/project/update-project]", Name: "[broker-service] - Successfully updated project"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF UPDATE PROJECT  -------------
// -------------------------------------------

// -------------------------------------------
// --------- START OF DELETE PROJECTS  ----------
// -------------------------------------------

func (app *Config) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectIdPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Delete project [/project/delete-project]", Name: "[broker-service] - Delete project request recieved"})
	app.DeleteProjectCall(w, requestPayload)
}

func (app *Config) DeleteProjectCall(w http.ResponseWriter, requestPayload ProjectIdPayload) {
	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/delete-project", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete project"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - delete project"))
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
	payload.Message = "delete project successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Delete project successfully [/project/delete-project]", Name: "[broker-service] - Successfully deleteted project"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF DELETE PROJECT  ----------
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET PROJECTS (ID)  -----
// -------------------------------------------

func (app *Config) GetProjectById(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectIdPayload

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project by id [/project/get-project-by-id]", Name: "[broker-service]"})

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/get-project-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch project"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get project by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - get project by id"))
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
	payload.Message = "get project by id successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project by id successfully [/project/get-project-by-id]", Name: "[broker-service] - Successfully fetched project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF GET PROJECTS (ID)  -------
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET PROJECTS  ----------
// -------------------------------------------

func (app *Config) GetAllProjects(w http.ResponseWriter, r *http.Request) {

	app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all projects [/project/get-all-projects]", Name: "[broker-service]"})

	request, err := http.NewRequest("GET", "http://project-service/project/get-all-projects", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch projects"))
		return
	}

	log.Println("response", response)

	defer response.Body.Close()

	var jsonFromService []Project

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all projects success [/project/get-all-projects]", Name: "[broker-service]"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// --------- END OF GET PROJECTS  ------------
// -------------------------------------------
