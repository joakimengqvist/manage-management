package main

import (
	"errors"
	"fmt"
	"net/http"
	"project-service/cmd/data"
)

type ProjectIdPayload struct {
	Id string `json:"id"`
}

type UpdateProject struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Status string   `json:"status"`
	Notes  []string `json:"notes"`
}

type UpdateProjectNote struct {
	NoteId    string `json:"noteId"`
	ProjectId string `json:"projectId"`
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
		fmt.Println("------- authorized, err", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	if !authorized {
		fmt.Println("------- !authorized")
		app.errorJSON(w, errors.New("could not create projects: Unauthorized"), http.StatusUnauthorized)
		return
	}

	var requestPayload data.NewProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Println("-------  app.readJSON", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	Project := data.NewProject{
		Name:   requestPayload.Name,
		Status: requestPayload.Status,
	}

	response, err := app.Models.Project.Insert(Project)
	if err != nil {
		fmt.Println("-------  response", err)
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

	updatedProject := data.PostgresProject{
		ID:     requestPayload.Id,
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
		app.logItemViaRPC(w, nil, RPCLogData{Action: "Get project by id [/auth/get-project-by-id]", Name: "[project-service] - Failed to read JSON payload: " + err.Error()})
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	project, err := app.Models.Project.GetProjectById(requestPayload.Id)
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
// --------- START OF UPDATE PROJECT NOTES ---
// -------------------------------------------

func (app *Config) UpdateProjectNotes(w http.ResponseWriter, r *http.Request) {

	/*
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
			app.errorJSON(w, errors.New("could not delete project: Unauthorized"), http.StatusUnauthorized)
			return
		}

	*/

	var requestPayload UpdateProjectNote

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	fmt.Println("------- requestPayload", requestPayload)

	project, err := app.Models.Project.GetProjectById(requestPayload.ProjectId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get project by id"), http.StatusBadRequest)
		return
	}

	fmt.Println("------- project", project)

	notesSlice := app.parsePostgresArray(project.Notes)
	fmt.Println("------- notesSlice", notesSlice)

	updatedNotes := append(notesSlice, requestPayload.NoteId)
	fmt.Println("------- updatedNotes", updatedNotes)

	updatedProject := data.PostgresProject{
		ID:     requestPayload.ProjectId,
		Name:   project.Name,
		Status: project.Status,
		Notes:  app.convertToPostgresArray(updatedNotes),
	}

	fmt.Println("------- updatedProject", updatedProject)

	err = updatedProject.Update()
	if err != nil {
		fmt.Println("------- err", err)
		app.errorJSON(w, errors.New("could not update project: "+err.Error()), http.StatusBadRequest)
		return
	}

	returnedProject := data.Project{
		ID:     requestPayload.ProjectId,
		Name:   project.Name,
		Status: project.Status,
		Notes:  updatedNotes,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated project with Id: %s", fmt.Sprint(updatedProject.ID)),
		Data:    returnedProject,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Update project [/project/update-project]", Name: "[project-service] - Successfully updated project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF DELETE PROJECT  ----------
// -------------------------------------------
