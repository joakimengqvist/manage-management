package main

import (
	"authentication/cmd/data"
	"errors"
	"fmt"
	"net/http"
)

type NewUser struct {
	Email      string   `json:"email"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Password   string   `json:"-"`
	Privileges []string `json:"privileges"`
	Projects   []string `json:"projects"`
}

type Authenticate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Id         string   `json:"id"`
	Email      string   `json:"email"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Privileges []string `json:"privileges"`
	Projects   []string `json:"projects"`
}

type UserIdPayload struct {
	Id string `json:"id"`
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload Authenticate

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	/*
		valid, err := user.PasswordMatches(requestPayload.Password)
		if err != nil || !valid {
			app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
			return
		}

		if err != nil {
			app.errorJSON(w, err)
		}
	*/

	returnedUser := data.ReturnedUser{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		Privileges: app.parsePostgresArray(user.Privileges),
		Projects:   app.parsePostgresArray(user.Projects),
		LastName:   user.LastName,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    returnedUser,
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Authenticate [/auth/authenticate]", Name: "[authentication-service] - Successful authentication"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewUser

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "user_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newUser := data.User{
		Email:      requestPayload.Email,
		FirstName:  requestPayload.FirstName,
		LastName:   requestPayload.LastName,
		Privileges: app.convertToPostgresArray(requestPayload.Privileges),
		Projects:   app.convertToPostgresArray(requestPayload.Projects),
		Password:   requestPayload.Password,
	}

	response, err := app.Models.User.InsertUser(newUser)
	if err != nil {
		app.errorJSON(w, errors.New("could not create user: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", requestPayload.FirstName),
		Data:    response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateUser

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "user_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	updatedUser := data.User{
		ID:         requestPayload.Id,
		Email:      requestPayload.Email,
		FirstName:  requestPayload.FirstName,
		Privileges: app.convertToPostgresArray(requestPayload.Privileges),
		Projects:   app.convertToPostgresArray(requestPayload.Projects),
		LastName:   requestPayload.LastName,
	}

	err = updatedUser.UpdateUser()
	if err != nil {
		app.errorJSON(w, errors.New("could not update user: "+err.Error()), http.StatusBadRequest)
		return
	}

	returnedData := data.ReturnedUser{
		ID:         requestPayload.Id,
		Email:      requestPayload.Email,
		FirstName:  requestPayload.FirstName,
		Privileges: requestPayload.Privileges,
		Projects:   requestPayload.Projects,
		LastName:   requestPayload.LastName,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated user with Id: %s", fmt.Sprint(updatedUser.ID)),
		Data:    returnedData,
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Authenticate [/auth/update-user]", Name: "[authentication-service] - Successful updated user"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteUser(w http.ResponseWriter, r *http.Request) {

	var requestPayload UserIdPayload

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "user_sudo")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetUserById(requestPayload.Id)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get user by id"), http.StatusBadRequest)
		return
	}

	err = user.DeleteUser()
	if err != nil {
		app.errorJSON(w, errors.New("could not delete user: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted user: %s %s", user.FirstName, user.LastName),
		Data:    nil,
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Authenticate [/auth/delete-user]", Name: "[authentication-service] - Successful deleted user"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetUserById(w http.ResponseWriter, r *http.Request) {
	var requestPayload UserIdPayload

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "user_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetUserById(requestPayload.Id)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get user by id"), http.StatusBadRequest)
		return
	}

	returnedUser := data.ReturnedUser{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Privileges: app.parsePostgresArray(user.Privileges),
		Projects:   app.parsePostgresArray(user.Projects),
		Password:   "",
		Active:     user.Active,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched user: %s %s", user.FirstName, user.LastName),
		Data:    returnedUser,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get user by id [/auth/get-user-by-id]", Name: "[authentication-service] - Successfuly fetched user"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	err := app.CheckUserPrivilege(w, userId, "user_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	users, err := app.Models.User.GetAllUsers()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var userSlice []data.ReturnedUser
	for _, userPtr := range users {
		user := *userPtr

		returnedUser := data.ReturnedUser{
			ID:         user.ID,
			Email:      user.Email,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Privileges: app.parsePostgresArray(user.Privileges),
			Projects:   app.parsePostgresArray(user.Projects),
			Password:   user.Password,
			Active:     user.Active,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		}

		userSlice = append(userSlice, returnedUser)
	}

	app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all users [/auth/get-all-users]", Name: "[authentication-service] - Successfuly fetched all users"})
	app.writeUsersJSONFromSlice(w, http.StatusAccepted, userSlice)
}

func (app *Config) CheckUserPrivilege(w http.ResponseWriter, userId string, action string) error {

	user, err := app.Models.User.GetUserById(userId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get user by id"), http.StatusBadRequest)
		return err
	}

	privileges := app.parsePostgresArray(user.Privileges)
	isAuthenticated := app.containsString(privileges, action)

	if !isAuthenticated {
		app.errorJSON(w, errors.New("Unauthorized"), http.StatusUnauthorized)
		return errors.New("Unauthorized")
	}

	return nil
}

type CheckPrivilegePayload struct {
	UserId string `json:"userId"`
	Action string `json:"action"`
}

type CheckPrivilegeResponse struct {
	Authenticated bool   `json:"authenticated"`
	Message       string `json:"message"`
}

func (app *Config) CheckPrivilege(w http.ResponseWriter, r *http.Request) {
	var requestPayload CheckPrivilegePayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetUserById(requestPayload.UserId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get user by id"), http.StatusBadRequest)
		return
	}

	privileges := app.parsePostgresArray(user.Privileges)
	isAuthenticated := app.containsString(privileges, requestPayload.Action)

	if !isAuthenticated {
		app.errorJSON(w, errors.New("Unauthorized"), http.StatusUnauthorized)
		return
	}

	payload := CheckPrivilegeResponse{
		Authenticated: true,
		Message:       "Authenticated",
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Authenticate [/auth/delete-user]", Name: "[authentication-service] - Successful deleted user"})
	app.writeJSON(w, http.StatusAccepted, payload)
}
