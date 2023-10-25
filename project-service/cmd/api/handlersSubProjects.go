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
	Invoices          []string  `json:"invoices"`
	Incomes           []string  `json:"incomes"`
	Expenses          []string  `json:"expenses"`
}

type UpdateSubProjectNote struct {
	NoteId       string `json:"note_id"`
	SubProjectId string `json:"sub_project_id"`
}

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
		Notes:             app.convertToPostgresArray(requestPayload.Notes),
		Invoices:          app.convertToPostgresArray(requestPayload.Invoices),
		Incomes:           app.convertToPostgresArray(requestPayload.Incomes),
		Expenses:          app.convertToPostgresArray(requestPayload.Expenses),
	}

	response, err := app.Models.SubProject.InsertSubProject(SubProject, userId)
	if err != nil {
		fmt.Println("ERROR InsertSubProject", err)
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
		Description:       requestPayload.Description,
		Status:            requestPayload.Status,
		Priority:          requestPayload.Priority,
		StartDate:         requestPayload.StartDate,
		DueDate:           requestPayload.DueDate,
		EstimatedDuration: requestPayload.EstimatedDuration,
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
		Description:       subProject.Description,
		Priority:          subProject.Priority,
		StartDate:         subProject.StartDate,
		DueDate:           subProject.DueDate,
		EstimatedDuration: subProject.EstimatedDuration,
		CreatedAt:         subProject.CreatedAt,
		CreatedBy:         subProject.CreatedBy,
		UpdatedAt:         subProject.UpdatedAt,
		UpdatedBy:         subProject.UpdatedBy,
		Projects:          app.parsePostgresArray(subProject.Projects),
		Notes:             app.parsePostgresArray(subProject.Notes),
		Invoices:          app.parsePostgresArray(subProject.Invoices),
		Incomes:           app.parsePostgresArray(subProject.Incomes),
		Expenses:          app.parsePostgresArray(subProject.Expenses),
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched subProject: %s", subProject.Name),
		Data:    returnedSubProject,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get subProject by id [/auth/get-sub-project-by-id]", Name: "[subProject-service] - Successfuly fetched subProject"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

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
			CreatedAt: subProject.CreatedAt,
			CreatedBy: subProject.CreatedBy,
			UpdatedAt: subProject.UpdatedAt,
			UpdatedBy: subProject.UpdatedBy,
			Projects:  app.parsePostgresArray(subProject.Projects),
			Notes:     app.parsePostgresArray(subProject.Notes),
			Invoices:  app.parsePostgresArray(subProject.Invoices),
			Incomes:   app.parsePostgresArray(subProject.Incomes),
			Expenses:  app.parsePostgresArray(subProject.Expenses),
		}

		subProjectSlice = append(subProjectSlice, returnSubProject)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all subProjects",
		Data:    subProjectSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetSubProjectsByIds(w http.ResponseWriter, r *http.Request) {

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

	var requestPayload IdsPayload

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	subProjects, err := data.GetSubProjectsByIds(app.convertToPostgresArray(requestPayload.Ids))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var subProjectSlice []data.SubProject
	for _, subProjectPtr := range subProjects {
		subProject := *subProjectPtr

		returnedSubProject := data.SubProject{
			ID:        subProject.ID,
			Name:      subProject.Name,
			Status:    subProject.Status,
			Priority:  subProject.Priority,
			StartDate: subProject.StartDate,
			DueDate:   subProject.DueDate,
			CreatedAt: subProject.CreatedAt,
			CreatedBy: subProject.CreatedBy,
			UpdatedAt: subProject.UpdatedAt,
			UpdatedBy: subProject.UpdatedBy,
			Projects:  app.parsePostgresArray(subProject.Projects),
			Notes:     app.parsePostgresArray(subProject.Notes),
			Invoices:  app.parsePostgresArray(subProject.Invoices),
			Incomes:   app.parsePostgresArray(subProject.Incomes),
			Expenses:  app.parsePostgresArray(subProject.Expenses),
		}

		subProjectSlice = append(subProjectSlice, returnedSubProject)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all subProjects by ids",
		Data:    subProjectSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

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

	payload := jsonResponse{
		Error:   false,
		Message: "appended sub project note",
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

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
		fmt.Println("DeleteSubProjectNote", err)
		app.errorJSON(w, errors.New("failed to delete note from subProject"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "deleted sub project note",
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
