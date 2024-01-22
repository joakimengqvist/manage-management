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

type NewInvoiceNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	InvoiceId   string `json:"invoice_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateInvoiceNotePayload struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	InvoiceId   string `json:"invoice_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type ReturnedInvoiceNotes struct {
	Notes []data.InvoiceNote `json:"notes"`
}

type UpdateInvoiceNote struct {
	NoteId    string `json:"note_id"`
	InvoiceId string `json:"invoice_id"`
}

type DeleteInvoiceNoteIdPayload struct {
	AuthorId  string `json:"author_id"`
	NoteId    string `json:"note_id"`
	InvoiceId string `json:"invoice_id"`
}

func (app *Config) CreateInvoiceNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewInvoiceNote

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_write")
	if err != nil {
		log.Println("authenticated - CreateInvoiceNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("authenticated - CreateInvoiceNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateInvoiceNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newNote := data.InvoiceNote{
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		InvoiceId:   requestPayload.InvoiceId,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	noteId, err := app.Models.InvoiceNote.InsertInvoiceNote(newNote)
	if err != nil {
		log.Println("postgres - InsertInvoiceNote", err)
		app.errorJSON(w, errors.New("could not create invoice note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("created note %s", requestPayload.Title),
		Data:    noteId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoiceNotesByInvoiceId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllInvoiceNotesByInvoiceId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetAllInvoiceNotesByInvoiceId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllInvoiceNotesByInvoiceId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	notes, err := app.Models.InvoiceNote.GetInvoiceNotesByInvoiceId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetInvoiceNotesByInvoiceId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.InvoiceNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.InvoiceNote{
			ID:          note.ID,
			AuthorId:    note.AuthorId,
			AuthorName:  note.AuthorName,
			AuthorEmail: note.AuthorEmail,
			InvoiceId:   note.InvoiceId,
			Title:       note.Title,
			Note:        note.Note,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all invoice notes by invoice id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoiceNotesByUserId(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetAllInvoiceNotesByUserId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllInvoiceNotesByUserId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllInvoiceNotesByUserId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	notes, err := app.Models.InvoiceNote.GetInvoiceNotesByAuthorId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetInvoiceNotesByAuthorId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.InvoiceNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.InvoiceNote{
			ID:          note.ID,
			AuthorId:    note.AuthorId,
			AuthorName:  note.AuthorName,
			AuthorEmail: note.AuthorEmail,
			InvoiceId:   note.InvoiceId,
			Title:       note.Title,
			Note:        note.Note,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all invoice notes by user id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetInvoiceNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetInvoiceNoteById", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetInvoiceNoteById")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetInvoiceNoteById", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.InvoiceNote.GetInvoiceNoteById(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetInvoiceNoteById", err)
		app.errorJSON(w, errors.New("failed to get invoice note by id"), http.StatusBadRequest)
		return
	}

	returnedNote := data.InvoiceNote{
		ID:          note.ID,
		AuthorId:    note.AuthorId,
		AuthorName:  note.AuthorName,
		AuthorEmail: note.AuthorEmail,
		InvoiceId:   note.InvoiceId,
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

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Get invoice note by id [/notes/get-invoice-note-by-id]", Name: "[notes-service] - Successfuly fetched invoice note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateInvoiceNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateInvoiceNotePayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - UpdateInvoiceNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - UpdateInvoiceNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateInvoiceNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	returnedNote := data.InvoiceNote{
		ID:          requestPayload.ID,
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		InvoiceId:   requestPayload.InvoiceId,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	err = returnedNote.UpdateInvoiceNote()
	if err != nil {
		log.Println("postgres - UpdateInvoiceNote", err)
		app.errorJSON(w, errors.New("could not update invoice note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated note with Id: %s", fmt.Sprint(returnedNote.ID)),
		Data:    returnedNote,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Invoice notes [/notes/update-invoice-note]", Name: "[notes-service] - Successful updated invoice-note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteInvoiceNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteInvoiceNoteIdPayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_sudo")
	if err != nil {
		log.Println("authenticated - DeleteInvoiceNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - DeleteInvoiceNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - DeleteInvoiceNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.InvoiceNote.GetInvoiceNoteById(requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - GetInvoiceNoteById", err)
		app.errorJSON(w, errors.New("failed to get invoice note by id"), http.StatusBadRequest)
		return
	}

	err = app.RemoveNoteFromInvoice(w, r, note.ID, requestPayload.InvoiceId)
	if err != nil {
		log.Println("RemoveNoteFromInvoice", err)
		app.errorJSON(w, err)
		return
	}

	err = data.DeleteInvoiceNote(requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - DeleteInvoiceNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted note with Id: %s", fmt.Sprint(requestPayload.NoteId)),
		Data:    nil,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete invoice note [/auth/delete-user-note]", Name: "[authentication-service] - Successful deleted user note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) RemoveNoteFromInvoice(w http.ResponseWriter, r *http.Request, noteId string, invoiceId string) error {
	deleteInvoiceNote := UpdateInvoiceNote{
		NoteId:    noteId,
		InvoiceId: invoiceId,
	}

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_sudo")
	if err != nil {
		log.Println("authenticated - RemoveNoteFromInvoice", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return err
	}

	if !authenticated {
		log.Println("!authenticated - RemoveNoteFromInvoice")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	jsonDataUser, _ := json.MarshalIndent(deleteInvoiceNote, "", "")

	endpoint := "http://" + os.Getenv("INVOICE_SERVICE_SERVICE_HOST") + "/invoice/delete-invoice-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonDataUser))

	if err != nil {
		log.Println("POST - RemoveNoteFromInvoice", err)
		app.errorJSON(w, err)
		return err
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - RemoveNoteFromInvoice", err)
		app.errorJSON(w, err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete note from invoice"))
		return err
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling invoice service - delete note from invoice"))
		return err
	}

	return nil
}
