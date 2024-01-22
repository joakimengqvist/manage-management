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
		log.Println("readJSON - CreateInvoiceItemNote", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/create-invoice-item-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("http.NewRequest - CreateInvoiceItemNote", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - CreateInvoiceItemNote", err)
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
		log.Println("json.NewDecoder - CreateInvoiceItemNote", err)
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
		log.Println("readJSON - UpdateInvoiceItemNote", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/update-invoice-item-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("http.NewRequest - UpdateInvoiceItemNote", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - UpdateInvoiceItemNote", err)
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
		log.Println("json.NewDecoder - UpdateInvoiceItemNote", err)
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
		log.Println("readJSON - GetInvoiceItemNoteById", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-invoice-item-note-by-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetInvoiceItemNoteById", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetInvoiceItemNoteById", err)
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
		log.Println("json.NewDecoder - GetInvoiceItemNoteById", err)
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
		log.Println("readJSON - GetAllInvoiceItemNotesByInvoiceItemId", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-all-invoice-item-notes-by-invoice-item-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetAllInvoiceItemNotesByInvoiceItemId", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetAllInvoiceItemNotesByInvoiceItemId", err)
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
		log.Println("json.NewDecoder - GetAllInvoiceItemNotesByInvoiceItemId", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllInvoiceItemNotesByUserId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllInvoiceItemNotesByUserId", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-all-invoice-item-notes-by-user-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - GetAllInvoiceItemNotesByUserId", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - GetAllInvoiceItemNotesByUserId", err)
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
		log.Println("json.NewDecoder - GetAllInvoiceItemNotesByUserId", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) DeleteInvoiceItemNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - DeleteInvoiceItemNote", err)
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/delete-invoice-item-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("POST - DeleteInvoiceItemNote", err)
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - DeleteInvoiceItemNote", err)
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
		log.Println("json.NewDecoder - DeleteInvoiceItemNote", err)
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}
