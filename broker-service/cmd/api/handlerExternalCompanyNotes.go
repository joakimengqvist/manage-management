package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type ExternalCompanyNote struct {
	ID              string    `json:"id"`
	AuthorId        string    `json:"author_id"`
	AuthorName      string    `json:"author_name"`
	AuthorEmail     string    `json:"author_email"`
	ExternalCompany string    `json:"external_company"`
	Title           string    `json:"title"`
	Note            string    `json:"note"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type NewExternalCompanyNote struct {
	AuthorId        string `json:"author_id"`
	AuthorName      string `json:"author_name"`
	AuthorEmail     string `json:"author_email"`
	ExternalComapny string `json:"external_company"`
	Title           string `json:"title"`
	Note            string `json:"note"`
}

type UpdateExternalCompanyNote struct {
	ID              string `json:"id"`
	AuthorId        string `json:"author_id"`
	AuthorName      string `json:"author_name"`
	AuthorEmail     string `json:"author_email"`
	ExternalCompany string `json:"external_company"`
	Title           string `json:"title"`
	Note            string `json:"note"`
}

// ----------------------------------------------------
// --------- START OF CREATE EXTERNAL COMPANY NOTE  ---
// ----------------------------------------------------

func (app *Config) CreateExternalCompanyNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewExternalCompanyNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/create-external-company-note", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - create external company note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - create external company note"))
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
	payload.Message = "create external company note successful"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// --------- END OF CREATE EXTERNAL COMPANY NOTE  -----
// ----------------------------------------------------

// -------------------------------------------------
// --------- START OF UPDATE EXTERNAL NOTE NOTE  ---
// -------------------------------------------------

func (app *Config) UpdateExternalCompanyNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateExternalCompanyNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/update-external-company-note", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - update external company note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - update external company note"))
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
	payload.Message = "update external company note successful"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// --------- END OF UPDATE EXTERNAL COMPANY NOTE  -----
// ----------------------------------------------------

// ----------------------------------------------------
// --------- START OF GET EXTERNAL COMPANY NOTE (ID) --
// ----------------------------------------------------

func (app *Config) GetExternalCompanyNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-external-company-note-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch external company note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get external company note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get external company note by id"))
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
	payload.Message = "get external company note by id successful"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}

// --------------------------------------------------------------
// ---- END OF GET EXTERNAL COMPANY NOTE (ID) -------------------
// --------------------------------------------------------------

// --------------------------------------------------------------
// -- START OF GET EXTERNAL COMPANY NOTES (externalCompantId) ---
// --------------------------------------------------------------

func (app *Config) GetAllExternalCompanyNotesByExternalCompanyId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-all-external-company-notes-by-external-company-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch external company note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get external company note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get external company note by id"))
		return
	}

	var jsonFromService []ExternalCompanyNote

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------------------------
// --- END OF GET EXTERNAL COMAPNY NOTES (externalCompanyId) ---
// -------------------------------------------------------------

// -------------------------------------------------------------
// --- START OF GET EXTERNAL COMAPNY NOTES (userId) ------------
// -------------------------------------------------------------

func (app *Config) GetAllExternalCompanyNotesByUserId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-all-external-company-notes-by-user-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch external company note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get external company note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get external company note by id"))
		return
	}

	var jsonFromService []ExternalCompanyNote

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// ----------------------------------------------------
// --- END OF GET EXTERNAL COMPANY NOTES (userId) -----
// ----------------------------------------------------

// ----------------------------------------------------
// --- START OF DELETE EXTERNAL COMPANY NOTE (id) -----
// ----------------------------------------------------

func (app *Config) DeleteExternalCompanyNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/delete-external-company-note", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not delete external company note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete external company note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - delete external company note by id"))
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "get external company note by id successful"
	payload.Data = nil

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------
// --- END OF DELETE EXTERNAL COMPANY NOTE (id) -------
// ----------------------------------------------------
