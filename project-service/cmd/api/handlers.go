package main

import (
	"errors"
	"fmt"
	"net/http"
	"project-service/cmd/data"
)

type ProjectIdPayload struct {
	ID string `json:"id"`
}

type UpdateProject struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Status string   `json:"status"`
	Notes  []string `json:"notes"`
}

type UpdateProjectNote struct {
	NoteId    string `json:"note_id"`
	ProjectId string `json:"project_id"`
}

// -------------------------------------------
// --------- START OF CREATE PROJECT  --------
// -------------------------------------------

func (app *Config) CreateProject(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "project_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload data.NewProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	Project := data.NewProject{
		Name:   requestPayload.Name,
		Status: requestPayload.Status,
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

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Create project [/project/create-project]", Name: "[project-service] - Successfully created new project"})
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
	authenticated, err := app.CheckPrivilege(w, userId, "project_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload UpdateProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	updatedProject := data.PostgresProject{
		ID:     requestPayload.ID,
		Name:   requestPayload.Name,
		Status: requestPayload.Status,
		Notes:  app.convertToPostgresArray(requestPayload.Notes),
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

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Update project [/project/update-project]", Name: "[project-service] - Successfully updated project"})
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
	authenticated, err := app.CheckPrivilege(w, userId, "project_sudo")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload ProjectIdPayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	project, err := app.Models.Project.GetProjectById(requestPayload.ID)
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

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete project [/project/delete-project]", Name: "[project-service] - Successful deleted project"})
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
	authenticated, err := app.CheckPrivilege(w, userId, "project_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload ProjectIdPayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.logItemViaRPC(w, nil, RPCLogData{Action: "Get project by id [/auth/get-project-by-id]", Name: "[project-service] - Failed to read JSON payload: " + err.Error()})
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	project, err := app.Models.Project.GetProjectById(requestPayload.ID)
	if err != nil {
		app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project by id [/auth/get-project-by-id]", Name: "[project-service] - Failed to get project by id: " + err.Error()})
		app.errorJSON(w, errors.New("failed to get project by id"), http.StatusBadRequest)
		return
	}

	returnedProject := data.Project{
		ID:     project.ID,
		Name:   project.Name,
		Status: project.Status,
		Notes:  app.parsePostgresArray(project.Notes),
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched project: %s", project.Name),
		Data:    project,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get project by id [/auth/get-project-by-id]", Name: "[project-service] - Successfuly fetched project"})
	app.writeJSON(w, http.StatusAccepted, returnedProject)
}

// -------------------------------------------
// --------- END OF GET PROJECT  -------------
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET ALL PROJECTS  ------
// -------------------------------------------

func (app *Config) GetAllProjects(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "project_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	projects, err := app.Models.Project.GetAll()
	if err != nil {
		app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all projects [/auth/get-all-projects]", Name: "[project-service] - Failed to read JSON payload" + err.Error()})
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var projectSlice []data.Project
	for _, projectPtr := range projects {
		project := *projectPtr

		returnProject := data.Project{
			ID:        project.ID,
			Name:      project.Name,
			Status:    project.Status,
			Notes:     app.parsePostgresArray(project.Notes),
			CreatedAt: project.CreatedAt,
			UpdatedAt: project.UpdatedAt,
		}

		projectSlice = append(projectSlice, returnProject)
	}

	app.logItemViaRPC(w, projectSlice, RPCLogData{Action: "Get all projects [/auth/get-all-projects]", Name: "[project-service] - Successfuly fetched all projects"})
	app.writeJSONFromSlice(w, http.StatusAccepted, projectSlice)
}

// -------------------------------------------
// --------- END OF GET ALL PROJECTS  --------
// -------------------------------------------

// -------------------------------------------
// --------- START OF ADD PROJECT NOTES ------
// -------------------------------------------

func (app *Config) AddProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProjectNote

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_write")
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

	err = data.AppendProjectNote(requestPayload.ProjectId, requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to append note to project"), http.StatusBadRequest)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Project [/project/create-project-note]", Name: "[project-service] - Successful added project note"})
	app.writeJSON(w, http.StatusAccepted, requestPayload)
}

// -------------------------------------------
// --------- END OF ADD PROJECT NOTES --------
// -------------------------------------------

// -------------------------------------------
// --------- START OF REMOVE PROJECT NOTES ---
// -------------------------------------------

func (app *Config) RemoveProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProjectNote

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_write")
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
	err = data.DeleteProjectNote(requestPayload.ProjectId, requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to delete note from project"), http.StatusBadRequest)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Project [/project/delete-project-note]", Name: "[project-service] - Successful deleted project note"})
	app.writeJSON(w, http.StatusAccepted, requestPayload)
}
