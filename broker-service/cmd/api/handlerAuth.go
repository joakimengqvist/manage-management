package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	Privileges []string  `json:"privileges"`
	Projects   []int     `json:"projects"`
	Password   string    `json:"-"`
	Active     int       `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserIdPayload struct {
	Id int `json:"id"`
}

type NewUser struct {
	Email      string   `json:"email"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Privileges []string `json:"privileges"`
	Projects   []int    `json:"projects"`
	Password   string   `json:"-"`
}

type UserUpdatePayload struct {
	ID         int      `json:"id"`
	Email      string   `json:"email"`
	FirstName  string   `json:"first_name"`
	Privileges []string `json:"privileges"`
	Projects   []int    `json:"projects"`
	LastName   string   `json:"last_name"`
}

// -------------------------------------------
// ------- START OF AUTHENTICATE  ------------
// -------------------------------------------

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload AuthPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	logpayload := requestPayload
	logpayload.Password = ""

	app.logItemViaRPC(w, logpayload, RPCLogData{Action: "Authenticate [/auth/authenticate]", Name: "[broker-service] - Authenticate call recieved"})
	app.AuthenticateCall(w, requestPayload)
}

func (app *Config) AuthenticateCall(w http.ResponseWriter, authPayload AuthPayload) {

	jsonData, _ := json.MarshalIndent(authPayload, "", "")

	request, err := http.NewRequest("POST", "http://authentication-service/auth/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - authenticate"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, authPayload, RPCLogData{Action: "Authenticate success [/auth/authenticate]", Name: "[broker-service]"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// ------- END OF AUTHENTICATE  --------------
// -------------------------------------------

// -------------------------------------------
// --------- START OF CREATE USER  -----------
// -------------------------------------------

func (app *Config) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewUser
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create user [/auth/create-user]", Name: "[broker-service] - Create user request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://authentication-service/auth/create-user", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - create user"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - create user"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "create user successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create user successfully [/auth/create-user]", Name: "[broker-service] - Successfully authenticated"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF CREATE USER  -------------
// -------------------------------------------

// -------------------------------------------
// --------- START OF UPDATE USER  -----------
// -------------------------------------------

func (app *Config) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload UserUpdatePayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Update user [/auth/update-user]", Name: "[broker-service] - Update user request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://authentication-service/auth/update-user", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - update user"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - update user"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "update user successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Updated user successfully [/auth/update-user]", Name: "[broker-service] - Successfully updated user"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF UPDATE USER  -------------
// -------------------------------------------

// -------------------------------------------
// --------- START OF DELETE USERS  ----------
// -------------------------------------------

func (app *Config) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload UserIdPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Delete user [/auth/delete-user]", Name: "[broker-service] - Delete user request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://authentication-service/auth/delete-user", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete user"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - delete user"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "delete user successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Delete user successfully [/auth/delete-user]", Name: "[broker-service] - Successfully deleteted user"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF DELETE USER  -------------
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET USER (ID)  ---------
// -------------------------------------------

func (app *Config) GetUserById(w http.ResponseWriter, r *http.Request) {
	var requestPayload UserIdPayload

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get user by id [/auth/get-user-by-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://authentication-service/auth/get-user-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch user"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get user by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - get user by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "get user by id successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get user by id successfully [/auth/get-user-by-id]", Name: "[broker-service] - Successfully fetched user"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF GET USER (ID)  -----------
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET USERS  -------------
// -------------------------------------------

func (app *Config) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all users [/auth/get-all-users]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	fmt.Println("------------  userId -----------", userId)

	request, err := http.NewRequest("GET", "http://authentication-service/auth/get-all-users", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch users"))
		return
	}

	defer response.Body.Close()

	var jsonFromService []User

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all users success [/auth/get-all-users]", Name: "[broker-service]"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// --------- END OF GET USERS  ---------------
// -------------------------------------------
