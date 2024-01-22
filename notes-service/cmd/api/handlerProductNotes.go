package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"notes-service/cmd/data"
	"os"
)

type NewProductNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	ProductId   string `json:"product_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateProductNotePayload struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	ProductId   string `json:"product_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type ReturnedProductNotes struct {
	Notes []data.ProductNote `json:"notes"`
}

type UpdateProductNote struct {
	NoteId    string `json:"note_id"`
	ProductId string `json:"product_id"`
}

type DeleteProductNoteIdPayload struct {
	AuthorId  string `json:"author_id"`
	NoteId    string `json:"note_id"`
	ProductId string `json:"product_id"`
}

func (app *Config) CreateProductNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewProductNote

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_write")
	if err != nil {
		log.Println("authenticated - CreateProductNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - CreateProductNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateProductNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newNote := data.ProductNote{
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		ProductId:   requestPayload.ProductId,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	noteId, err := app.Models.ProductNote.InsertProductNote(newNote)
	if err != nil {
		log.Println("postgres - InsertProductNote", err)
		app.errorJSON(w, errors.New("could not create product note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("created note %s", requestPayload.Title),
		Data:    noteId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllProductNotesByProductId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllProductNotesByProductId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetAllProductNotesByProductId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllProductNotesByProductId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	notes, err := app.Models.ProductNote.GetProductNotesByProductId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetProductNotesByProductId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.ProductNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.ProductNote{
			ID:          note.ID,
			AuthorId:    note.AuthorId,
			AuthorName:  note.AuthorName,
			AuthorEmail: note.AuthorEmail,
			ProductId:   note.ProductId,
			Title:       note.Title,
			Note:        note.Note,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all product notes by product id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllProductNotesByUserId(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetAllProductNotesByUserId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllProductNotesByUserId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllProductNotesByUserId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	notes, err := app.Models.ProductNote.GetProductNotesByAuthorId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetProductNotesByAuthorId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.ProductNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.ProductNote{
			ID:          note.ID,
			AuthorId:    note.AuthorId,
			AuthorName:  note.AuthorName,
			AuthorEmail: note.AuthorEmail,
			ProductId:   note.ProductId,
			Title:       note.Title,
			Note:        note.Note,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all product notes by user id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetProductNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetProductNoteById", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetProductNoteById")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetProductNoteById", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.ProductNote.GetProductNoteById(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetProductNoteById", err)
		app.errorJSON(w, errors.New("failed to get product note by id"), http.StatusBadRequest)
		return
	}

	returnedNote := data.ProductNote{
		ID:          note.ID,
		AuthorId:    note.AuthorId,
		AuthorName:  note.AuthorName,
		AuthorEmail: note.AuthorEmail,
		ProductId:   note.ProductId,
		Title:       note.Title,
		Note:        note.Note,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched note: %s", note.Title),
		Data:    returnedNote,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Get product note by id [/notes/get-product-note-by-id]", Name: "[notes-service] - Successfuly fetched product note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateProductNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProductNotePayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - UpdateProductNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - UpdateProductNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateProductNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	returnedNote := data.ProductNote{
		ID:          requestPayload.ID,
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		ProductId:   requestPayload.ProductId,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	err = returnedNote.UpdateProductNote()
	if err != nil {
		log.Println("postgres - UpdateProductNote", err)
		app.errorJSON(w, errors.New("could not update product note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated note with Id: %s", fmt.Sprint(returnedNote.ID)),
		Data:    returnedNote,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Product notes [/notes/update-product-note]", Name: "[notes-service] - Successful updated product-note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteProductNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteProductNoteIdPayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_sudo")
	if err != nil {
		log.Println("authenticated - DeleteProductNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - DeleteProductNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - DeleteProductNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.ProductNote.GetProductNoteById(requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - GetProductNoteById", err)
		app.errorJSON(w, errors.New("failed to get product note by id"), http.StatusBadRequest)
		return
	}

	err = app.RemoveNoteFromProduct(w, r, note.ID, requestPayload.ProductId)
	if err != nil {
		log.Println("RemoveNoteFromProduct", err)
		app.errorJSON(w, err)
		return
	}

	err = data.DeleteProductNote(requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - DeleteProductNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted note with Id: %s", fmt.Sprint(requestPayload.NoteId)),
		Data:    nil,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete product note [/auth/delete-user-note]", Name: "[authentication-service] - Successful deleted user note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) RemoveNoteFromProduct(w http.ResponseWriter, r *http.Request, noteId string, productId string) error {
	deleteProductNote := UpdateProductNote{
		NoteId:    noteId,
		ProductId: productId,
	}

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_sudo")
	if err != nil {
		log.Println("authenticated - RemoveNoteFromProduct", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return err
	}

	if !authenticated {
		log.Println("!authenticated - RemoveNoteFromProduct")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	jsonDataUser, _ := json.MarshalIndent(deleteProductNote, "", "")

	endpoint := "http://" + os.Getenv("PRODUCT_SERVICE_SERVICE_HOST") + "/product/delete-product-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonDataUser))
	if err != nil {
		log.Println("POST - RemoveNoteFromProduct", err)
		app.errorJSON(w, err)
		return err
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - RemoveNoteFromProduct", err)
		app.errorJSON(w, err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete note from product"))
		return err
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling product service - delete note from product"))
		return err
	}

	return nil
}
