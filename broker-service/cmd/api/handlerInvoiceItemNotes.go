package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type InvoiceItemNote struct {
	ID            string    `json:"id"`
	AuthorId      string    `json:"author_id"`
	AuthorName    string    `json:"author_name"`
	AuthorEmail   string    `json:"author_email"`
	InvoiceItemId string    `json:"invoice_item_id"`
	Title         string    `json:"title"`
	Note          string    `json:"note"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type NewInvoiceItemNote struct {
	AuthorId      string `json:"author_id"`
	AuthorName    string `json:"author_name"`
	AuthorEmail   string `json:"author_email"`
	InvoiceItemId string `json:"invoice_item_id"`
	Title         string `json:"title"`
	Note          string `json:"note"`
}

type UpdateInvoiceItemNote struct {
	ID            string `json:"id"`
	AuthorId      string `json:"author_id"`
	AuthorName    string `json:"author_name"`
	AuthorEmail   string `json:"author_email"`
	InvoiceItemId string `json:"invoice_item_id"`
	Title         string `json:"title"`
	Note          string `json:"note"`
}

type DeleteInvoiceItemNotePayload struct {
	NoteId        string `json:"note_id"`
	AuthorId      string `json:"author_id"`
	InvoiceItemId string `json:"invoice_item_id"`
}

func (app *Config) CreateInvoiceItemNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewInvoiceItemNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/create-invoice-item-note", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - create invoice item note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - create invoice item note"))
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

func (app *Config) UpdateInvoiceItemNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateInvoiceItemNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/update-invoice-item-note", bytes.NewBuffer(jsonData))
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
		app.errorJSON(w, errors.New("status unauthorized - update invoice item note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - update invoice item note"))
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

func (app *Config) GetInvoiceItemNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-invoice-item-note-by-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch invoice item note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get invoice item note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - get invoice item note by id"))
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

func (app *Config) GetAllInvoiceItemNotesByInvoiceItemId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-all-invoice-item-notes-by-invoice-item-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch invoice item note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get invoice item note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get invoice item note by id"))
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

func (app *Config) GetAllInvoiceItemNotesByUserId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/get-all-invoice-item-notes-by-user-id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch invoice item note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get invoice item note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get invoice item note by id"))
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

func (app *Config) DeleteInvoiceItemNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	request, err := http.NewRequest("POST", "http://notes-service/notes/delete-invoice-item-note", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not delete invoice item note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete invoice item note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - delete invoice item note by id"))
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
