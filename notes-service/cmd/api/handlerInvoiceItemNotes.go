package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"notes-service/cmd/data"
	"os"
)

type NewInvoiceItemNote struct {
	AuthorId      string `json:"author_id"`
	AuthorName    string `json:"author_name"`
	AuthorEmail   string `json:"author_email"`
	InvoiceItemId string `json:"invoice_item_id"`
	Title         string `json:"title"`
	Note          string `json:"note"`
}

type UpdateInvoiceItemNotePayload struct {
	ID            string `json:"id"`
	AuthorId      string `json:"author_id"`
	AuthorName    string `json:"author_name"`
	AuthorEmail   string `json:"author_email"`
	InvoiceItemId string `json:"invoice_item_id"`
	Title         string `json:"title"`
	Note          string `json:"note"`
}

type ReturnedInvoiceItemNotes struct {
	Notes []data.InvoiceItemNote `json:"notes"`
}

type UpdateInvoiceItemNote struct {
	NoteId        string `json:"note_id"`
	InvoiceItemId string `json:"invoice_item_id"`
}

type DeleteInvoiceItemNoteIdPayload struct {
	AuthorId      string `json:"author_id"`
	NoteId        string `json:"note_id"`
	InvoiceItemId string `json:"invoice_item_id"`
}

func (app *Config) CreateInvoiceItemNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewInvoiceItemNote

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newNote := data.InvoiceItemNote{
		AuthorId:      requestPayload.AuthorId,
		AuthorName:    requestPayload.AuthorName,
		AuthorEmail:   requestPayload.AuthorEmail,
		InvoiceItemId: requestPayload.InvoiceItemId,
		Title:         requestPayload.Title,
		Note:          requestPayload.Note,
	}

	noteId, err := app.Models.InvoiceItemNote.InsertInvoiceItemNote(newNote)
	if err != nil {
		app.errorJSON(w, errors.New("could not create invoice item note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("created note %s", requestPayload.Title),
		Data:    noteId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoiceItemNotesByInvoiceItemId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	fmt.Println("GetAllNotesByInvoiceItemId")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Println("Readjson", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	fmt.Print("requestPayload", requestPayload)

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	notes, err := app.Models.InvoiceItemNote.GetInvoiceItemNotesByInvoiceItemId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.InvoiceItemNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.InvoiceItemNote{
			ID:            note.ID,
			AuthorId:      note.AuthorId,
			AuthorName:    note.AuthorName,
			AuthorEmail:   note.AuthorEmail,
			InvoiceItemId: note.InvoiceItemId,
			Title:         note.Title,
			Note:          note.Note,
			CreatedAt:     note.CreatedAt,
			UpdatedAt:     note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all invoice item notes by invoice item id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoiceItemNotesByUserId(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	notes, err := app.Models.InvoiceItemNote.GetInvoiceItemNotesByAuthorId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.InvoiceItemNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.InvoiceItemNote{
			ID:            note.ID,
			AuthorId:      note.AuthorId,
			AuthorName:    note.AuthorName,
			AuthorEmail:   note.AuthorEmail,
			InvoiceItemId: note.InvoiceItemId,
			Title:         note.Title,
			Note:          note.Note,
			CreatedAt:     note.CreatedAt,
			UpdatedAt:     note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all invoice item notes by user id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetInvoiceItemNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.InvoiceItemNote.GetInvoiceItemNoteById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get invoice item note by id"), http.StatusBadRequest)
		return
	}

	returnedNote := data.InvoiceItemNote{
		ID:            note.ID,
		AuthorId:      note.AuthorId,
		AuthorName:    note.AuthorName,
		AuthorEmail:   note.AuthorEmail,
		InvoiceItemId: note.InvoiceItemId,
		Title:         note.Title,
		Note:          note.Note,
		CreatedAt:     note.CreatedAt,
		UpdatedAt:     note.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched note: %s", note.Title),
		Data:    returnedNote,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Get invoice item note by id [/notes/get-invoice-note-by-id]", Name: "[notes-service] - Successfuly fetched invoice item note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateInvoiceItemNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateInvoiceItemNotePayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	returnedNote := data.InvoiceItemNote{
		ID:            requestPayload.ID,
		AuthorId:      requestPayload.AuthorId,
		AuthorName:    requestPayload.AuthorName,
		AuthorEmail:   requestPayload.AuthorEmail,
		InvoiceItemId: requestPayload.InvoiceItemId,
		Title:         requestPayload.Title,
		Note:          requestPayload.Note,
	}

	err = returnedNote.UpdateInvoiceItemNote()
	if err != nil {
		app.errorJSON(w, errors.New("could not update invoice item note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated note with Id: %s", fmt.Sprint(returnedNote.ID)),
		Data:    returnedNote,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "InvoiceItem notes [/notes/update-invoice-note]", Name: "[notes-service] - Successful updated invoice-note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteInvoiceItemNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteInvoiceItemNoteIdPayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_sudo")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.InvoiceItemNote.GetInvoiceItemNoteById(requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get invoice item note by id"), http.StatusBadRequest)
		return
	}

	err = app.RemoveNoteFromInvoiceItem(w, r, note.ID, requestPayload.InvoiceItemId)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = data.DeleteInvoiceItemNote(requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted note with Id: %s", fmt.Sprint(requestPayload.NoteId)),
		Data:    nil,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete invoice item note [/auth/delete-user-note]", Name: "[authentication-service] - Successful deleted user note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) RemoveNoteFromInvoiceItem(w http.ResponseWriter, r *http.Request, noteId string, invoiceId string) error {
	deleteInvoiceItemNote := UpdateInvoiceItemNote{
		NoteId:        noteId,
		InvoiceItemId: invoiceId,
	}

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_sudo")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return err
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	jsonDataUser, _ := json.MarshalIndent(deleteInvoiceItemNote, "", "")

	endpoint := "http://" + os.Getenv("INVOICE_SERVICE_SERVICE_HOST") + "/invoice/delete-invoice-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonDataUser))

	if err != nil {
		app.errorJSON(w, err)
		return err
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - delete note from invoice"))
		return err
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling invoice item service - delete note from invoice"))
		return err
	}

	return nil
}
