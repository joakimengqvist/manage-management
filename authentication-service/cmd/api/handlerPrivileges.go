package main

import (
	"authentication/cmd/data"
	"errors"
	"fmt"
	"net/http"
)

type NewPrivilege struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PrivilegeIdPayload struct {
	ID string `json:"id"`
}

type UpdatePrivilege struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (app *Config) GetAllPrivileges(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "privilege_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	privileges, err := app.Models.Privilege.GetAllPrivileges()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var privilegesSlice []data.Privilege
	for _, privilegePtr := range privileges {
		privilege := *privilegePtr

		returnedSlice := data.Privilege{
			ID:          privilege.ID,
			Name:        privilege.Name,
			Description: privilege.Description,
			CreatedAt:   privilege.CreatedAt,
			UpdatedAt:   privilege.UpdatedAt,
		}

		privilegesSlice = append(privilegesSlice, returnedSlice)
	}

	app.logItemViaRPC(w, privilegesSlice, RPCLogData{Action: "Get all privileges [/auth/get-all-privileges]", Name: "[authentication-service] - Successfuly fetched all privileges"})
	app.writePrivilegesJSONFromSlice(w, http.StatusAccepted, privilegesSlice)
}

func (app *Config) CreatePrivilege(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewPrivilege

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "privilege_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newPrivilege := data.Privilege{
		Name:        requestPayload.Name,
		Description: requestPayload.Description,
	}

	response, err := app.Models.Privilege.InsertPrivilege(newPrivilege)
	if err != nil {
		app.errorJSON(w, errors.New("could not create privilege: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created privilege: %s", requestPayload.Name),
		Data:    response,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Create privilege [/auth/create-privilege]", Name: "[authentication-service] - Successfuly created new privilege"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetPrivilegeById(w http.ResponseWriter, r *http.Request) {
	var requestPayload PrivilegeIdPayload

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "privilege_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	privilege, err := app.Models.Privilege.GetPrivilegeById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get privilege by id"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched privilege: %s", privilege.Name),
		Data:    privilege,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get privilege by id [/auth/get-privilege-by-id]", Name: "[authentication-service] - Successfuly fetched privilege"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdatePrivilege(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdatePrivilege

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "privilege_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	updatedPrivilege := data.Privilege{
		ID:          requestPayload.ID,
		Name:        requestPayload.Name,
		Description: requestPayload.Description,
	}

	err = updatedPrivilege.UpdatePrivilege()
	if err != nil {
		app.errorJSON(w, errors.New("could not update privilege: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated privilege with Id: %s", fmt.Sprint(updatedPrivilege.ID)),
		Data:    updatedPrivilege,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Update privilege [/auth/update-privilege]", Name: "[authentication-service] - Successful updated privilege"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeletePrivilege(w http.ResponseWriter, r *http.Request) {
	var requestPayload PrivilegeIdPayload

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "privilege_sudo")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	privilege, err := app.Models.Privilege.GetPrivilegeById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get privilege by id"), http.StatusBadRequest)
		return
	}

	err = privilege.DeletePrivilege()
	if err != nil {
		app.errorJSON(w, errors.New("could not delete privilege: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted privilege: %s", privilege.Name),
		Data:    nil,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete privilege [/auth/delete-privilege]", Name: "[authentication-service] - Successful deleted privilege"})
	app.writeJSON(w, http.StatusAccepted, payload)
}
