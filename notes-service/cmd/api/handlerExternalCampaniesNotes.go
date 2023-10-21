package main

import (
	"errors"
	"fmt"
	"net/http"
	"notes-service/cmd/data"
)

type NewExternalCompanyNote struct {
	AuthorId        string `json:"author_id"`
	AuthorName      string `json:"author_name"`
	AuthorEmail     string `json:"author_email"`
	ExternalCompany string `json:"external_company"`
	Title           string `json:"title"`
	Note            string `json:"note"`
}

type UpdateExternalCompanyNotePayload struct {
	ID              string `json:"id"`
	AuthorId        string `json:"author_id"`
	AuthorName      string `json:"author_name"`
	AuthorEmail     string `json:"author_email"`
	ExternalCompany string `json:"external_company"`
	Title           string `json:"title"`
	Note            string `json:"note"`
}

type ReturnedIExternalCompanyNotes struct {
	Notes []data.ExternalCompanyNote `json:"notes"`
}

// -----------------------------------------------------
// ------- START OF CREATE EXTERNAL COMPANY NOTE  ------
// -----------------------------------------------------

func (app *Config) CreateExternalCompanyNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewExternalCompanyNote

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

	newNote := data.ExternalCompanyNote{
		AuthorId:        requestPayload.AuthorId,
		AuthorName:      requestPayload.AuthorName,
		AuthorEmail:     requestPayload.AuthorEmail,
		ExternalCompany: requestPayload.ExternalCompany,
		Title:           requestPayload.Title,
		Note:            requestPayload.Note,
	}

	noteId, err := app.Models.ExternalCompanyNote.InsertExternalCompanyNote(newNote)
	if err != nil {
		app.errorJSON(w, errors.New("could not create external company note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("created external company note %s", requestPayload.Title),
		Data:    noteId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ----------------------------------------------------------
// -- END OF CREATE EXTERNAL COMPANY NOTE  ------------------
// ----------------------------------------------------------

// ----------------------------------------------------------
// -- START OF GET ALL EXTERNAL COMPANY NOTES (userId) ------
// ----------------------------------------------------------

func (app *Config) GetAllExternalCompanyNotesByUserId(w http.ResponseWriter, r *http.Request) {

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

	notes, err := app.Models.ExternalCompanyNote.GetExternalCompanyNotesByAuthorId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.ExternalCompanyNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.ExternalCompanyNote{
			ID:              note.ID,
			AuthorId:        note.AuthorId,
			AuthorName:      note.AuthorName,
			AuthorEmail:     note.AuthorEmail,
			ExternalCompany: note.ExternalCompany,
			Title:           note.Title,
			Note:            note.Note,
			CreatedAt:       note.CreatedAt,
			UpdatedAt:       note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched external company notes by user id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// --------------------------------------------------------------------
// -- END OF GET ALL EXTERNAL COMPANY NOTES (userId) ------------------
// --------------------------------------------------------------------

// --------------------------------------------------------------------
// -- START OF GET ALL EXTERNAL COMPANY NOTES (externalCompanyId) -----
// --------------------------------------------------------------------

func (app *Config) GetAllExternalCompanyNotesByExternalCompanyId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

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

	notes, err := app.Models.ExternalCompanyNote.GetExternalCompanyNotesByExternalCompanyId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.ExternalCompanyNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.ExternalCompanyNote{
			ID:              note.ID,
			AuthorId:        note.AuthorId,
			AuthorName:      note.AuthorName,
			AuthorEmail:     note.AuthorEmail,
			ExternalCompany: note.ExternalCompany,
			Title:           note.Title,
			Note:            note.Note,
			CreatedAt:       note.CreatedAt,
			UpdatedAt:       note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched external company notes by external company id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------------------------
// -- END OF GET ALL EXTERNAL COMPANY NOTES (externalCompanyId) -----
// ------------------------------------------------------------------

// ------------------------------------------------------------------
// -- START OF GET EXTERNAL COMPANY NOTE BY ID ----------------------
// ------------------------------------------------------------------

func (app *Config) GetExternalCompanyNoteById(w http.ResponseWriter, r *http.Request) {
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

	note, err := app.Models.ExternalCompanyNote.GetExternalCompanyNoteById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get external company note by id"), http.StatusBadRequest)
		return
	}

	returnedNote := data.ExternalCompanyNote{
		ID:              note.ID,
		AuthorId:        note.AuthorId,
		AuthorName:      note.AuthorName,
		AuthorEmail:     note.AuthorEmail,
		ExternalCompany: note.ExternalCompany,
		Title:           note.Title,
		Note:            note.Note,
		CreatedAt:       note.CreatedAt,
		UpdatedAt:       note.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched external company note: %s", note.Title),
		Data:    returnedNote,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF GET EXTERNAL COMPANY NOTE BY ID ------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF UPDATE EXTERNAL COMPANY NOTE -------
// ------------------------------------------------

func (app *Config) UpdateExternalCompanyNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateExternalCompanyNotePayload

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

	returnedNote := data.ExternalCompanyNote{
		ID:              requestPayload.ID,
		AuthorId:        requestPayload.AuthorId,
		AuthorName:      requestPayload.AuthorName,
		AuthorEmail:     requestPayload.AuthorEmail,
		ExternalCompany: requestPayload.ExternalCompany,
		Title:           requestPayload.Title,
		Note:            requestPayload.Note,
	}

	err = returnedNote.UpdateExternalCompanyNote()
	if err != nil {
		app.errorJSON(w, errors.New("could not update external company note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated note with Id: %s", fmt.Sprint(returnedNote.ID)),
		Data:    returnedNote,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF UPDATE EXTERNAL COMPANY NOTE ---------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF DELETE EXTERNAL COMPANY NOTE -------
// ------------------------------------------------

func (app *Config) DeleteExternalCompanyNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

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

	err = data.DeleteExternalCompanyNote(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted external company note with Id: %s", fmt.Sprint(requestPayload.ID)),
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF DELETE EXTERNAL COMPANY NOTE ---------
// ------------------------------------------------
