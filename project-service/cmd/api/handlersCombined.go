package main

import (
	"errors"
	"log"
	"net/http"
	"project-service/cmd/data"
)

type ProjectsAndSubProject struct {
	ProjectIds   []string `json:"project_ids"`
	SubProjectId string   `json:"sub_project_id"`
}

type SubProjectsAndProject struct {
	ProjectId     string   `json:"project_id"`
	SubProjectIds []string `json:"sub_project_ids"`
}

func (app *Config) AddProjectsSubProjectConnection(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "sub_project_write")
	if err != nil {
		log.Println("authenticated - AddProjectsSubProjectConnection", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - AddProjectsSubProjectConnection")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload ProjectsAndSubProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - AddProjectsSubProjectConnection", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = data.AppendProjectsToSubProject(requestPayload.ProjectIds, requestPayload.SubProjectId)
	if err != nil {
		log.Println("postgres - AddProjectsSubProjectConnection", err)
		app.errorJSON(w, errors.New("failed to add project to subProject"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Added projects to sub project",
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) RemoveProjectsSubProjectConnection(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectsAndSubProject

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "sub_project_write")
	if err != nil {
		log.Println("authenticated - RemoveProjectsSubProjectConnection", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - RemoveProjectsSubProjectConnection")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - RemoveProjectsSubProjectConnection", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = data.DeleteProjectsFromSubProject(requestPayload.ProjectIds, requestPayload.SubProjectId)
	if err != nil {
		log.Println("postgres - RemoveProjectsSubProjectConnection", err)
		app.errorJSON(w, errors.New("failed to remove project from sub project"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Removed projects to sub project",
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) AddSubProjectsProjectConnection(w http.ResponseWriter, r *http.Request) {
	var requestPayload SubProjectsAndProject

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "sub_project_write")
	if err != nil {
		log.Println("authenticated - AddSubProjectsProjectConnection", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - AddSubProjectsProjectConnection")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - AddSubProjectsProjectConnection", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = data.AppendSubProjectsToProject(requestPayload.ProjectId, requestPayload.SubProjectIds)
	if err != nil {
		log.Println("postgres - AddSubProjectsProjectConnection", err)
		app.errorJSON(w, errors.New("failed to add sub project to project"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Added sub projects to project",
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) RemoveSubProjectsProjectConnection(w http.ResponseWriter, r *http.Request) {
	var requestPayload SubProjectsAndProject

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "sub_project_write")
	if err != nil {
		log.Println("authenticated - RemoveSubProjectsProjectConnection", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - RemoveSubProjectsProjectConnection")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - RemoveSubProjectsProjectConnection", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = data.DeleteSubProjectsFromProject(requestPayload.ProjectId, requestPayload.SubProjectIds)
	if err != nil {
		log.Println("postgres - RemoveSubProjectsProjectConnection", err)
		app.errorJSON(w, errors.New("failed to add sub project to project"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Removed sub projects to project",
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
