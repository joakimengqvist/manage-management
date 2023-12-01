package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type InvoiceNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	InvoiceId   string    `json:"invoice_id"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewInvoiceNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	InvoiceId   string `json:"invoice_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateInvoiceNote struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	InvoiceId   string `json:"invoice_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type DeleteInvoiceNotePayload struct {
	NoteId    string `json:"note_id"`
	AuthorId  string `json:"author_id"`
	InvoiceId string `json:"invoice_id"`
}

func (app *Config) CreateInvoiceNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewInvoiceNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/create-invoice-note", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - create invoice note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - create invoice note"))
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

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) UpdateInvoiceNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateInvoiceNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/update-invoice-note", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - update invoice note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - update invoice note"))
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

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetInvoiceNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-invoice-note-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch invoice note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get invoice note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - get invoice note by id"))
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

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllInvoiceNotesByInvoiceId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-all-invoice-notes-by-invoice-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch invoice note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get invoice note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get invoice note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllInvoiceNotesByUserId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-all-invoice-notes-by-user-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch invoice note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get invoice note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get invoice note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) DeleteInvoiceNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/delete-invoice-note", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not delete invoice note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete invoice note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - delete invoice note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}
