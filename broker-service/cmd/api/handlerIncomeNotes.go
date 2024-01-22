package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

type IncomeNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	IncomeId    string    `json:"income_id"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewIncomeNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	IncomeId    string `json:"income_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateIncomeNote struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	IncomeId    string `json:"income_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type DeleteIncomeNotePayload struct {
	NoteId   string `json:"note_id"`
	AuthorId string `json:"author_id"`
	IncomeId string `json:"income_id"`
}

func (app *Config) CreateIncomeNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewIncomeNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/create-income-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - CreateIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - CreateIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - create income note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - create income note"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - CreateIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) UpdateIncomeNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateIncomeNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/update-income-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - UpdateIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - UpdateIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - update income note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - update income note"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - UpdateIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetIncomeNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetIncomeNoteById", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-income-note-by-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetIncomeNoteById", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetIncomeNoteById", err)
		app.errorJSON(w, errors.New("could not fetch income note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get income note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - get income note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - GetIncomeNoteById", err)
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllIncomeNotesByIncomeId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllIncomeNotesByIncomeId", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-all-income-notes-by-income-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetAllIncomeNotesByIncomeId", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetAllIncomeNotesByIncomeId", err)
		app.errorJSON(w, errors.New("could not fetch income note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get income note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get income note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - GetAllIncomeNotesByIncomeId", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllIncomeNotesByUserId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllIncomeNotesByUserId", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-all-income-notes-by-user-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetAllIncomeNotesByUserId", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetAllIncomeNotesByUserId", err)
		app.errorJSON(w, errors.New("could not fetch income note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get income note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get income note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - GetAllIncomeNotesByUserId", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) DeleteIncomeNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - DeleteIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/delete-income-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - DeleteIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - DeleteIncomeNote", err)
		app.errorJSON(w, errors.New("could not delete income note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete income note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - delete income note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		log.Println("json.NewDecoder - DeleteIncomeNote", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}
