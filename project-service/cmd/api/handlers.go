package main

import (
	"errors"
	"fmt"
	"net/http"
	"project-service/cmd/data"
)

type NewProject struct {
	Name string `json:"name"`
}

type ProjectIdPayload struct {
	Id string `json:"id"`
}

type UpdateProject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// -------------------------------------------
// --------- START OF CREATE PROJECT  --------
// -------------------------------------------

func (app *Config) CreateProject(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	privilegePayload := Privilege{
		UserId: userId,
		Action: "project_write",
	}

	authorized, err := app.CheckPrivilege(w, privilegePayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	if !authorized {
		app.errorJSON(w, errors.New("could not create projects: Unauthorized"), http.StatusUnauthorized)
		return
	}

	var requestPayload NewProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	Project := data.Project{
		Name: requestPayload.Name,
	}

	response, err := app.Models.Project.Insert(Project)
	if err != nil {
		app.errorJSON(w, errors.New("could not create project: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created project %s", requestPayload.Name),
		Data:    response,
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create project [/project/create-project]", Name: "[project-service] - Successfully created new project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF CREATE PROJECT  ----------
// -------------------------------------------

// -------------------------------------------
// --------- START OF UPDATE PROJECT  --------
// -------------------------------------------

func (app *Config) UpdateProject(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	privilegePayload := Privilege{
		UserId: userId,
		Action: "project_write",
	}

	authorized, err := app.CheckPrivilege(w, privilegePayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	if !authorized {
		app.errorJSON(w, errors.New("could not update project: Unauthorized"), http.StatusUnauthorized)
		return
	}

	var requestPayload UpdateProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	updatedProject := data.Project{
		ID:   requestPayload.Id,
		Name: requestPayload.Name,
	}

	err = updatedProject.Update()
	if err != nil {
		app.errorJSON(w, errors.New("could not update project: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated project with Id: %s", fmt.Sprint(updatedProject.ID)),
		Data:    updatedProject,
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Update project [/project/update-project]", Name: "[project-service] - Successfully updated project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF UPDATE PROJECT  ----------
// -------------------------------------------

// -------------------------------------------
// --------- START OF DELETE PROJECT  --------
// -------------------------------------------

func (app *Config) DeleteProject(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	privilegePayload := Privilege{
		UserId: userId,
		Action: "project_sudo",
	}

	authorized, err := app.CheckPrivilege(w, privilegePayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	if !authorized {
		app.errorJSON(w, errors.New("could not delete project: Unauthorized"), http.StatusUnauthorized)
		return
	}

	var requestPayload ProjectIdPayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	project, err := app.Models.Project.GetProjectById(requestPayload.Id)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get project by id"), http.StatusBadRequest)
		return
	}

	err = project.Delete()
	if err != nil {
		app.errorJSON(w, errors.New("could not delete project: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted project: %s", project.Name),
		Data:    nil,
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Delete project [/project/delete-project]", Name: "[project-service] - Successful deleted project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF DELETE PROJECT  ----------
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET PROJECT  -----------
// -------------------------------------------

func (app *Config) GetProjectById(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	privilegePayload := Privilege{
		UserId: userId,
		Action: "project_read",
	}

	authorized, err := app.CheckPrivilege(w, privilegePayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	if !authorized {
		app.errorJSON(w, errors.New("could not get project: Unauthorized"), http.StatusUnauthorized)
		return
	}

	var requestPayload ProjectIdPayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project by id [/auth/get-project-by-id]", Name: "[project-service] - Failed to read JSON payload" + err.Error()})
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	project, err := app.Models.Project.GetProjectById(requestPayload.Id)
	if err != nil {
		app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project by id [/auth/get-project-by-id]", Name: "[project-service] - Failed to get project by id" + err.Error()})
		app.errorJSON(w, errors.New("failed to get project by id"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched project: %s", project.Name),
		Data:    project,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get project by id [/auth/get-project-by-id]", Name: "[project-service] - Successfuly fetched project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF GET PROJECT  -------------
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET ALL PROJECTS  ------
// -------------------------------------------

func (app *Config) GetAllProjects(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	privilegePayload := Privilege{
		UserId: userId,
		Action: "project_read",
	}

	authorized, err := app.CheckPrivilege(w, privilegePayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	if !authorized {
		app.errorJSON(w, errors.New("could not get all projects: Unauthorized"), http.StatusUnauthorized)
		return
	}

	projects, err := app.Models.Project.GetAll()
	if err != nil {
		app.logItemViaRPC(w, err, RPCLogData{Action: "Get all projects [/auth/get-all-projects]", Name: "[project-service] - Failed to read JSON payload" + err.Error()})
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var projectSlice []data.Project
	for _, projectPtr := range projects {
		projectSlice = append(projectSlice, *projectPtr)
	}

	app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all projects [/auth/get-all-projects]", Name: "[project-service] - Successfuly fetched all projects"})
	app.writeJSONFromSlice(w, http.StatusAccepted, projectSlice)
}

// -------------------------------------------
// --------- END OF GET ALL PROJECTS  --------
// -------------------------------------------
