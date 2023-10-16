package main

import (
	"errors"
	"net/http"
	"project-service/cmd/data"
)

type ProjectAndSubProject struct {
	ProjectId    string `json:"project_id"`
	SubProjectId string `json:"sub_project_id"`
}

// ----------------------------------------------------
// --------- START OF ADD PROJECT TO SUB PROJECT ------
// ----------------------------------------------------

func (app *Config) AddProjectSubProjectConnection(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectAndSubProject

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "sub_project_write")
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

	err = data.AppendProjectToSubProject(requestPayload.ProjectId, requestPayload.SubProjectId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to add project to subProject"), http.StatusBadRequest)
		return
	}

	err = data.AppendSubProjectToProject(requestPayload.ProjectId, requestPayload.SubProjectId)
	if err != nil {
		data.DeleteProjectFromSubProject(requestPayload.ProjectId, requestPayload.SubProjectId)
		app.errorJSON(w, errors.New("failed to add project to subProject"), http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusAccepted, requestPayload)
}

// -----------------------------------------------
// --------- END OF ADD SUB PROJECT NOTES --------
// -----------------------------------------------

// ----------------------------------------------------
// --------- START OF REMOVE PROJECT TO SUB PROJECT ---
// ----------------------------------------------------

func (app *Config) RemoveProjectSubProjectConnection(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProjectAndSubProject

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "sub_project_write")
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

	err = data.AppendSubProjectToProject(requestPayload.ProjectId, requestPayload.SubProjectId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to append project to subProject"), http.StatusBadRequest)
		return
	}

	err = data.AppendProjectToSubProject(requestPayload.ProjectId, requestPayload.SubProjectId)
	if err != nil {
		data.DeleteProjectFromSubProject(requestPayload.ProjectId, requestPayload.SubProjectId)
		app.errorJSON(w, errors.New("failed to append project to subProject"), http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusAccepted, requestPayload)
}

// ----------------------------------------------------
// --------- END OF REMOVE PROJECT TO SUB PROJECT -----
// ----------------------------------------------------
