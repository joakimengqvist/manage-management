package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"project-service/cmd/data"
)

type UpdateProject struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type UpdateProjectNote struct {
	NoteId    string `json:"note_id"`
	ProjectId string `json:"project_id"`
}

func (app *Config) CreateProject(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "project_write")
	if err != nil {
		log.Println("authenticated - CreateProject", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - CreateProject")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload data.NewProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateProject", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	Project := data.NewProject{
		Name:   requestPayload.Name,
		Status: requestPayload.Status,
	}

	response, err := app.Models.Project.InsertProject(Project, userId)
	if err != nil {
		log.Println("postgres - CreateProject", err)
		app.errorJSON(w, errors.New("could not create project: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created project %s", requestPayload.Name),
		Data:    response,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Create project [/project/create-project]", Name: "[project-service] - Successfully created new project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateProject(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "project_write")
	if err != nil {
		log.Println("authenticated - UpdateProject", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - UpdateProject")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload UpdateProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateProject", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	updatedProject := data.PostgresProject{
		ID:     requestPayload.ID,
		Name:   requestPayload.Name,
		Status: requestPayload.Status,
	}

	err = updatedProject.UpdateProject(userId)
	if err != nil {
		log.Println("postgres - UpdateProject", err)
		app.errorJSON(w, errors.New("could not update project: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated project with Id: %s", fmt.Sprint(updatedProject.ID)),
		Data:    updatedProject,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Update project [/project/update-project]", Name: "[project-service] - Successfully updated project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteProject(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "project_sudo")
	if err != nil {
		log.Println("authenticated - DeleteProject", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - DeleteProject")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload IDpayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - DeleteProject", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	project, err := app.Models.Project.GetProjectById(requestPayload.ID)
	if err != nil {
		log.Println("postgres - DeleteProject", err)
		app.errorJSON(w, errors.New("failed to get project by id"), http.StatusBadRequest)
		return
	}

	err = project.DeleteProject()
	if err != nil {
		log.Println("DeleteProject - DeleteProject", err)
		app.errorJSON(w, errors.New("could not delete project: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted project: %s", project.Name),
		Data:    nil,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete project [/project/delete-project]", Name: "[project-service] - Successful deleted project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetProjectById(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "project_read")
	if err != nil {
		log.Println("authenticated - GetProjectById", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetProjectById")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload IDpayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetProjectById", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	project, err := app.Models.Project.GetProjectById(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetProjectById", err)
		app.errorJSON(w, errors.New("failed to get project by id"), http.StatusBadRequest)
		return
	}

	returnedProject := data.Project{
		ID:          project.ID,
		Name:        project.Name,
		Status:      project.Status,
		Notes:       app.parsePostgresArray(project.Notes),
		SubProjects: app.parsePostgresArray(project.SubProjects),
		CreatedAt:   project.CreatedAt,
		CreatedBy:   project.CreatedBy,
		UpdatedAt:   project.UpdatedAt,
		UpdatedBy:   project.UpdatedBy,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched project: %s", project.Name),
		Data:    returnedProject,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Get project by id [/auth/get-project-by-id]", Name: "[project-service] - Successfuly fetched project"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllProjects(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "project_read")
	if err != nil {
		log.Println("authenticated - GetAllProjects", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllProjects")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	projects, err := app.Models.Project.GetAllProjects()
	if err != nil {
		log.Println("postgres - GetAllProjects", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var projectSlice []data.Project
	for _, projectPtr := range projects {
		project := *projectPtr

		returnProject := data.Project{
			ID:          project.ID,
			Name:        project.Name,
			Status:      project.Status,
			Notes:       app.parsePostgresArray(project.Notes),
			SubProjects: app.parsePostgresArray(project.SubProjects),
			CreatedAt:   project.CreatedAt,
			CreatedBy:   project.CreatedBy,
			UpdatedAt:   project.UpdatedAt,
			UpdatedBy:   project.UpdatedBy,
		}

		projectSlice = append(projectSlice, returnProject)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all projects",
		Data:    projectSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetProjectsByIds(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "project_read")
	if err != nil {
		log.Println("authenticated - GetProjectsByIds", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetProjectsByIds")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload IdsPayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetProjectsByIds", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	projects, err := data.GetProjectsByIds(app.convertToPostgresArray(requestPayload.Ids))
	if err != nil {
		log.Println("postgres - GetProjectsByIds", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var projectSlice []data.Project
	for _, projectPtr := range projects {
		project := *projectPtr

		returnedProject := data.Project{
			ID:          project.ID,
			Name:        project.Name,
			Status:      project.Status,
			Notes:       app.parsePostgresArray(project.Notes),
			SubProjects: app.parsePostgresArray(project.SubProjects),
			CreatedAt:   project.CreatedAt,
			CreatedBy:   project.CreatedBy,
			UpdatedAt:   project.UpdatedAt,
			UpdatedBy:   project.UpdatedBy,
		}

		projectSlice = append(projectSlice, returnedProject)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all projects by ids",
		Data:    projectSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) AddProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProjectNote

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_write")
	if err != nil {
		log.Println("authenticated - AddProjectNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - AddProjectNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - AddProjectNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = data.AppendProjectNote(requestPayload.ProjectId, requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - AddProjectNote", err)
		app.errorJSON(w, errors.New("failed to append note to project"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "added note to project project",
		Data:    nil,
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Project [/project/create-project-note]", Name: "[project-service] - Successful added project note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) RemoveProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProjectNote

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_write")
	if err != nil {
		log.Println("authenticated - RemoveProjectNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - RemoveProjectNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - RemoveProjectNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	err = data.DeleteProjectNote(requestPayload.ProjectId, requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - RemoveProjectNote", err)
		app.errorJSON(w, errors.New("failed to delete note from project"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "deletet note from project project",
		Data:    nil,
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Project [/project/delete-project-note]", Name: "[project-service] - Successful deleted project note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}
