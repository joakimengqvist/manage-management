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

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	SubProjects []string  `json:"sub_projects"`
	Notes       []string  `json:"notes"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
}

type UpdateProject struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

type NewProject struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ProjectsToSubProject struct {
	ProjectIds   []string `json:"project_ids"`
	SubProjectId string   `json:"sub_project_id"`
}

func (app *Config) CreateProject(w http.ResponseWriter, r *http.Request) {

	var requestPayload NewProject
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateProject", err)
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create project [/project/create-project]", Name: "[broker-service] - Create project request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/create-project"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - CreateProject", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - CreateProject", err)
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
		log.Println("json.NewDecoder - CreateProject", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Create project successfully [/project/create-project]", Name: "[broker-service]"})

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProject
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateProject", err)
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Update project [/project/update-project]", Name: "[broker-service] - Update project request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/update-project"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - UpdateProject", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - UpdateProject", err)
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
		log.Println("json.NewDecoder - UpdateProject", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Updated project successfully [/project/update-project]", Name: "[broker-service] - Successfully updated project"})

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - DeleteProject", err)
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Delete project [/project/delete-project]", Name: "[broker-service] - Delete project request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/delete-project"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - DeleteProject", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - DeleteProject", err)
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
		log.Println("json.NewDecoder - DeleteProject", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Delete project successfully [/project/delete-project]", Name: "[broker-service] - Successfully deleteted project"})

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetProjectById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetProjectById", err)
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project by id [/project/get-project-by-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/get-project-by-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetProjectById", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetProjectById", err)
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
		log.Println("json.NewDecoder - GetProjectById", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get project by id successfully [/project/get-project-by-id]", Name: "[broker-service] - Successfully fetched project"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllProjects(w http.ResponseWriter, r *http.Request) {

	// app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all projects [/project/get-all-projects]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/get-all-projects"

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println("GET - GetAllProjects", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetAllProjects", err)
		app.errorJSON(w, errors.New("could not fetch projects"))
		return
	}

	defer response.Body.Close()

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - GetAllProjects", err)
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get all projects success [/project/get-all-projects]", Name: "[broker-service]"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetProjectsByIds(w http.ResponseWriter, r *http.Request) {

	// app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all projects [/project/get-all-projects-by-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetProjectsByIds", err)
		app.errorJSON(w, err)
		return
	}

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/get-projects-by-ids"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetProjectsByIds", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetProjectsByIds", err)
		app.errorJSON(w, errors.New("could not fetch projects by ids"))
		return
	}

	defer response.Body.Close()

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - GetProjectsByIds", err)
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get all projects by ids success [/project/get-all-projects-by-id]", Name: "[broker-service]"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) AddProjectsSubProjectConnection(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectsToSubProject

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - AddProjectsSubProjectConnection", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/add-projects-sub-project-connection"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - AddProjectsSubProjectConnection", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - AddProjectsSubProjectConnection", err)
		app.errorJSON(w, errors.New("could not update sub project"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - update sub project"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling project service - update sub project"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - AddProjectsSubProjectConnection", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) AddSubProjectsSubProjectConnection(w http.ResponseWriter, r *http.Request) {
	var requestPayload SubProjectsToProject

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - AddSubProjectsSubProjectConnection", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/add-sub-projects-project-connection"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - AddSubProjectsSubProjectConnection", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - AddSubProjectsSubProjectConnection", err)
		app.errorJSON(w, errors.New("could not update sub project"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - update sub project"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling project service - update sub project"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - AddSubProjectsSubProjectConnection", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) RemoveProjectsSubProjectConnection(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectsToSubProject

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - RemoveProjectsSubProjectConnection", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/delete-projects-sub-project-connection"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - RemoveProjectsSubProjectConnection", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - RemoveProjectsSubProjectConnection", err)
		app.errorJSON(w, errors.New("could not update sub project"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - update sub project"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - update sub project"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - RemoveProjectsSubProjectConnection", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}
