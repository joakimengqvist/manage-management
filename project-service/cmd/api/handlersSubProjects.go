package main

import (
	"errors"
	"fmt"
	"net/http"
	"project-service/cmd/data"
	"time"
)

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

type UpdateSubProjectNote struct {
	NoteId       string `json:"note_id"`
	SubProjectId string `json:"project_id"`
}

// -----------------------------------------------
// --------- START OF CREATE SUB PROJECT  --------
// -----------------------------------------------

func (app *Config) CreateSubProject(w http.ResponseWriter, r *http.Request) {

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

	var requestPayload data.NewSubProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	SubProject := data.PostgresSubProject{
		Name:              requestPayload.Name,
		Description:       requestPayload.Description,
		Status:            requestPayload.Status,
		Priority:          requestPayload.Priority,
		StartDate:         requestPayload.StartDate,
		DueDate:           requestPayload.DueDate,
		EstimatedDuration: requestPayload.EstimatedDuration,
		ProjectID:         requestPayload.ProjectID,
		Notes:             app.convertToPostgresArray(requestPayload.Notes),
		Invoices:          app.convertToPostgresArray(requestPayload.Invoices),
		Incomes:           app.convertToPostgresArray(requestPayload.Incomes),
		Expenses:          app.convertToPostgresArray(requestPayload.Expenses),
	}

	response, err := app.Models.SubProject.InsertSubProject(SubProject, userId)
	if err != nil {
		app.errorJSON(w, errors.New("could not create subProject: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created subProject %s", requestPayload.Name),
		Data:    response,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Create subProject [/sub-project/create-subProject]", Name: "[subProject-service] - Successfully created new subProject"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -----------------------------------------------
// --------- END OF CREATE SUB PROJECT  ----------
// -----------------------------------------------

// -----------------------------------------------
// --------- START OF UPDATE SUB PROJECT  --------
// -----------------------------------------------

func (app *Config) UpdateSubProject(w http.ResponseWriter, r *http.Request) {

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

	var requestPayload UpdateSubProject

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	updatedSubProject := data.PostgresSubProject{
		ID:                requestPayload.ID,
		Name:              requestPayload.Name,
		Status:            requestPayload.Status,
		Priority:          requestPayload.Priority,
		StartDate:         requestPayload.StartDate,
		DueDate:           requestPayload.DueDate,
		EstimatedDuration: requestPayload.EstimatedDuration,
		ProjectID:         requestPayload.ProjectID,
		Notes:             app.convertToPostgresArray(requestPayload.Notes),
		Invoices:          app.convertToPostgresArray(requestPayload.Invoices),
		Incomes:           app.convertToPostgresArray(requestPayload.Incomes),
		Expenses:          app.convertToPostgresArray(requestPayload.Expenses),
	}

	err = updatedSubProject.UpdateSubProject(userId)
	if err != nil {
		app.errorJSON(w, errors.New("could not update subProject: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated subProject with Id: %s", fmt.Sprint(updatedSubProject.ID)),
		Data:    updatedSubProject,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Update subProject [/sub-project/update-subProject]", Name: "[subProject-service] - Successfully updated subProject"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -----------------------------------------------
// --------- END OF UPDATE SUB PROJECT  ----------
// -----------------------------------------------

// -----------------------------------------------
// --------- START OF DELETE SUB PROJECT  --------
// -----------------------------------------------

func (app *Config) DeleteSubProject(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "sub_project_sudo")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload IDpayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	subProject, err := app.Models.SubProject.GetSubProjectById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get subProject by id"), http.StatusBadRequest)
		return
	}

	err = subProject.DeleteSubProject()
	if err != nil {
		app.errorJSON(w, errors.New("could not delete subProject: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted subProject: %s", subProject.Name),
		Data:    nil,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete subProject [/sub-project/delete-subProject]", Name: "[subProject-service] - Successful deleted subProject"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -----------------------------------------------
// --------- END OF DELETE SUB PROJECT  ----------
// -----------------------------------------------

// -----------------------------------------------
// --------- START OF GET SUB PROJECT  -----------
// -----------------------------------------------

func (app *Config) GetSubProjectById(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "sub_project_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload IDpayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.logItemViaRPC(w, nil, RPCLogData{Action: "Get subProject by id [/auth/get-sub-project-by-id]", Name: "[subProject-service] - Failed to read JSON payload: " + err.Error()})
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	subProject, err := app.Models.SubProject.GetSubProjectById(requestPayload.ID)
	if err != nil {
		app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get subProject by id [/auth/get-sub-project-by-id]", Name: "[subProject-service] - Failed to get subProject by id: " + err.Error()})
		app.errorJSON(w, errors.New("failed to get subProject by id"), http.StatusBadRequest)
		return
	}

	returnedSubProject := data.SubProject{
		ID:                subProject.ID,
		Name:              subProject.Name,
		Status:            subProject.Status,
		Priority:          subProject.Priority,
		StartDate:         subProject.StartDate,
		DueDate:           subProject.DueDate,
		EstimatedDuration: subProject.EstimatedDuration,
		ProjectID:         subProject.ProjectID,
		CreatedAt:         subProject.CreatedAt,
		CreatedBy:         subProject.CreatedBy,
		UpdatedAt:         subProject.UpdatedAt,
		UpdatedBy:         subProject.UpdatedBy,
		Notes:             app.parsePostgresArray(subProject.Notes),
		Invoices:          app.parsePostgresArray(subProject.Invoices),
		Incomes:           app.parsePostgresArray(subProject.Incomes),
		Expenses:          app.parsePostgresArray(subProject.Expenses),
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched subProject: %s", subProject.Name),
		Data:    subProject,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get subProject by id [/auth/get-sub-project-by-id]", Name: "[subProject-service] - Successfuly fetched subProject"})
	app.writeJSON(w, http.StatusAccepted, returnedSubProject)
}

// -----------------------------------------------
// --------- END OF GET SUB PROJECT  -------------
// -----------------------------------------------

// -----------------------------------------------
// --------- START OF GET ALL SUB PROJECTS  ------
// -----------------------------------------------

func (app *Config) GetAllSubProjects(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "sub_project_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	projects, err := app.Models.SubProject.GetAllSubProjects()
	if err != nil {
		app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all projects [/auth/get-all-projects]", Name: "[subProject-service] - Failed to read JSON payload" + err.Error()})
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var subProjectSlice []data.SubProject
	for _, projectPtr := range projects {
		subProject := *projectPtr

		returnSubProject := data.SubProject{
			ID:        subProject.ID,
			Name:      subProject.Name,
			Status:    subProject.Status,
			Priority:  subProject.Priority,
			StartDate: subProject.StartDate,
			DueDate:   subProject.DueDate,
			ProjectID: subProject.ProjectID,
			CreatedAt: subProject.CreatedAt,
			CreatedBy: subProject.CreatedBy,
			UpdatedAt: subProject.UpdatedAt,
			UpdatedBy: subProject.UpdatedBy,
			Notes:     app.parsePostgresArray(subProject.Notes),
			Invoices:  app.parsePostgresArray(subProject.Invoices),
			Incomes:   app.parsePostgresArray(subProject.Incomes),
			Expenses:  app.parsePostgresArray(subProject.Expenses),
		}

		subProjectSlice = append(subProjectSlice, returnSubProject)
	}

	app.logItemViaRPC(w, subProjectSlice, RPCLogData{Action: "Get all projects [/auth/get-all-projects]", Name: "[subProject-service] - Successfuly fetched all projects"})
	app.writeSubProductJSONFromSlice(w, http.StatusAccepted, subProjectSlice)
}

// -----------------------------------------------
// --------- END OF GET ALL SUB PROJECTS  --------
// -----------------------------------------------

// -----------------------------------------------
// --------- START OF ADD SUB PROJECT NOTES ------
// -----------------------------------------------

func (app *Config) AddSubProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateSubProjectNote

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

	err = data.AppendSubProjectNote(requestPayload.SubProjectId, requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to append note to subProject"), http.StatusBadRequest)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "SubProject [/sub-project/create-sub-project-note]", Name: "[subProject-service] - Successful added subProject note"})
	app.writeJSON(w, http.StatusAccepted, requestPayload)
}

// -------------------------------------------
// --------- END OF ADD SUB PROJECT NOTES --------
// -------------------------------------------

// -----------------------------------------------
// --------- START OF REMOVE SUB PROJECT NOTES ---
// -----------------------------------------------

func (app *Config) RemoveSubProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateSubProjectNote

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
	err = data.DeleteSubProjectNote(requestPayload.SubProjectId, requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to delete note from subProject"), http.StatusBadRequest)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "SubProject [/sub-project/delete-sub-project-note]", Name: "[subProject-service] - Successful deleted subProject note"})
	app.writeJSON(w, http.StatusAccepted, requestPayload)
}
