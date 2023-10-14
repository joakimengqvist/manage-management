package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Note struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	Project     string    `json:"project"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewProjectNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	Project     string `json:"project"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateNote struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	Project     string `json:"project"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type DeleteNotePayload struct {
	NoteId    string `json:"note_id"`
	AuthorId  string `json:"author_id"`
	ProjectId string `json:"project_id"`
}

// -------------------------------------------
// --------- START OF CREATE PROJECT NOTE  ---
// -------------------------------------------

func (app *Config) CreateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewProjectNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create project note [/notes/create-project-note]", Name: "[broker-service] - Create project note request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/create-project-note", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - create project note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - create project note"))
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
	payload.Message = "create project note successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Create project note successfully [/notes/create-project-note]", Name: "[broker-service] - Successfully created project note"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF CREATE PROJECT NOTE  -----
// -------------------------------------------

// -------------------------------------------
// --------- START OF UPDATE PROJECT NOTE  ---
// -------------------------------------------

func (app *Config) UpdateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Update project note [/notes/update-project-note]", Name: "[broker-service] - Update project note request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/update-project-note", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - update project note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - update project note"))
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
	payload.Message = "update project note successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Updated project note successfully [/notes/update-project-note]", Name: "[broker-service] - Successfully updated project note"})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF UPDATE PROJECT NOTE  -----
// -------------------------------------------

// -------------------------------------------
// --------- START OF GET PROJECT NOTE (ID) --
// -------------------------------------------

func (app *Config) GetProjectNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project note by id [/notes/get-project-note-by-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-project-note-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch project note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get project note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - get project note by id"))
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
	payload.Message = "get project note by id successful"
	payload.Data = jsonFromService.Data

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get project note by id successfully [/notes/get-project-note-by-id]", Name: "[broker-service] - Successfully fetched project note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --------- END OF GET PROJECT NOTE (ID) ----
// -------------------------------------------

// -------------------------------------------
// -- START OF GET PROJECT NOTES (projectId) -
// -------------------------------------------

func (app *Config) GetAllProjectNotesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	// userId := r.Header.Get("X-User-Id")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project note by id [/notes/get-all-notes-by-project-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-all-notes-by-project-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch project note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get project note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get project note by id"))
		return
	}

	var jsonFromService []Note

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project note by id successfully [/notes/get-all-notes-by-project-id]", Name: "[broker-service] - Successfully fetched project note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// --- END OF GET PROJECT NOTES (projectId) --
// -------------------------------------------

// -------------------------------------------
// --- START OF GET PROJECT NOTES (userId) -----
// -------------------------------------------

func (app *Config) GetAllProjectNotesByUserId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	// userId := r.Header.Get("X-User-Id")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project note by id [/notes/get-all-notes-by-project-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-all-notes-by-user-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch project note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get project note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get project note by id"))
		return
	}

	var jsonFromService []Note

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project note by id successfully [/notes/get-all-notes-by-project-id]", Name: "[broker-service] - Successfully fetched project note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// --- END OF GET PROJECT NOTES (userId) -----
// -------------------------------------------

// -------------------------------------------
// --- START OF DELETE PROJECT NOTE (id) -----
// -------------------------------------------

func (app *Config) DeleteProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteNotePayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/delete-project-note", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not delete project note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete project note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - delete project note by id"))
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "get project note by id successful"
	payload.Data = nil

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get project note by id successfully [/notes/get-all-notes-by-project-id]", Name: "[broker-service] - Successfully fetched project note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// -------------------------------------------
// --- END OF DELETE PROJECT NOTE (id) -------
// -------------------------------------------
