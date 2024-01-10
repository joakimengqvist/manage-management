package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

type ProductNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	ProductId   string    `json:"product_id"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewProductNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	ProductId   string `json:"product_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateProductNote struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	ProductId   string `json:"product_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type DeleteProductNotePayload struct {
	NoteId    string `json:"note_id"`
	AuthorId  string `json:"author_id"`
	ProductId string `json:"product_id"`
}

func (app *Config) CreateProductNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewProductNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Create product note [/notes/create-product-note]", Name: "[broker-service] - Create product note request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/create-product-note"

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
		app.errorJSON(w, errors.New("status unauthorized - create product note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - create product note"))
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

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Create product note successfully [/notes/create-product-note]", Name: "[broker-service] - Successfully created product note"})

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) UpdateProductNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProductNote
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Update product note [/notes/update-product-note]", Name: "[broker-service] - Update product note request recieved"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/update-product-note"

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
		app.errorJSON(w, errors.New("status unauthorized - update product note"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - update product note"))
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

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Updated product note successfully [/notes/update-product-note]", Name: "[broker-service] - Successfully updated product note"})

	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetProductNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get product note by id [/notes/get-product-note-by-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-product-note-by-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch product note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get product note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - get product note by id"))
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

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get product note by id successfully [/notes/get-product-note-by-id]", Name: "[broker-service] - Successfully fetched product note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllProductNotesByProductId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get product note by id [/notes/get-all-notes-by-product-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-all-product-notes-by-product-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch product note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get product note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get product note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get product note by id successfully [/notes/get-all-notes-by-product-id]", Name: "[broker-service] - Successfully fetched product note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) GetAllProductNotesByUserId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Get product note by id [/notes/get-all-notes-by-product-id]", Name: "[broker-service]"})

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/get-all-product-notes-by-user-id"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch product note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - get product note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - get product note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get product note by id successfully [/notes/get-all-notes-by-product-id]", Name: "[broker-service] - Successfully fetched product note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

func (app *Config) DeleteProductNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteProductNotePayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	userId := r.Header.Get("X-User-Id")

	jsonData, _ := json.MarshalIndent(requestPayload, "", "")

	endpoint := "http://" + os.Getenv("NOTES_SERVICE_SERVICE_HOST") + "/notes/delete-product-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("could not delete product note"))
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete product note by id"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling notes service - delete product note by id"))
		return
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Get product note by id successfully [/notes/get-all-notes-by-product-id]", Name: "[broker-service] - Successfully fetched product note"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}
