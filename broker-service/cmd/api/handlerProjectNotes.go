package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

type ProjectNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	ProjectId   string    `json:"project_id"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewProjectNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	ProjectId   string `json:"project_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateProjectNote struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	ProjectId   string `json:"project_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type DeleteProjectNotePayload struct {
	NoteId    string `json:"note_id"`
	AuthorId  string `json:"author_id"`
	ProjectId string `json:"project_id"`
}

func (app *Config) CreateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewProjectNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create project note [/notes/create-project-note]", Name: "[broker-service] - Create project note request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/create-project-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
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

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Create project note successfully [/notes/create-project-note]", Name: "[broker-service] - Successfully created project note"})

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) UpdateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProjectNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Update project note [/notes/update-project-note]", Name: "[broker-service] - Update project note request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/update-project-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
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

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Updated project note successfully [/notes/update-project-note]", Name: "[broker-service] - Successfully updated project note"})

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetProjectNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project note by id [/notes/get-project-note-by-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-project-note-by-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
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

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get project note by id successfully [/notes/get-project-note-by-id]", Name: "[broker-service] - Successfully fetched project note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllProjectNotesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project note by id [/notes/get-all-notes-by-project-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-all-project-notes-by-project-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
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

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get project note by id successfully [/notes/get-all-notes-by-project-id]", Name: "[broker-service] - Successfully fetched project note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllProjectNotesByUserId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get project note by id [/notes/get-all-notes-by-project-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-all-project-notes-by-user-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
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

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get project note by id successfully [/notes/get-all-notes-by-project-id]", Name: "[broker-service] - Successfully fetched project note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) DeleteProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteProjectNotePayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/delete-project-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
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

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get project note by id successfully [/notes/get-all-notes-by-project-id]", Name: "[broker-service] - Successfully fetched project note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}
