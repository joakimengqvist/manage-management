package main

import (
	"errors"
	"fmt"
	"net/http"
	"notes-service/cmd/data"
)

type NewIncomeNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	Income      string `json:"income"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateIncomeNotePayload struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	Income      string `json:"income"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type ReturnedIncomeNotes struct {
	Notes []data.IncomeNote `json:"notes"`
}

// -------------------------------------------
// ------- START OF CREATE INCOME NOTE  ------
// -------------------------------------------

func (app *Config) CreateIncomeNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewIncomeNote

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

	newNote := data.IncomeNote{
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		Income:      requestPayload.Income,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	noteId, err := app.Models.IncomeNote.InsertIncomeNote(newNote)
	if err != nil {
		app.errorJSON(w, errors.New("could not create income note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("created income note %s", requestPayload.Title),
		Data:    noteId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF CREATE INCOME NOTE  ------------------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF GET ALL INCOME NOTES (userId) ------
// ------------------------------------------------

func (app *Config) GetAllIncomeNotesByUserId(w http.ResponseWriter, r *http.Request) {

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

	notes, err := app.Models.IncomeNote.GetIncomeNotesByAuthorId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.IncomeNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.IncomeNote{
			ID:          note.ID,
			AuthorId:    note.AuthorId,
			AuthorName:  note.AuthorName,
			AuthorEmail: note.AuthorEmail,
			Income:      note.Income,
			Title:       note.Title,
			Note:        note.Note,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched income by user id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF GET ALL INCOME NOTES (userId) --------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF GET ALL INCOME NOTES (incomeId) ----
// ------------------------------------------------

func (app *Config) GetAllIncomeNotesByIncomeId(w http.ResponseWriter, r *http.Request) {
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

	notes, err := app.Models.IncomeNote.GetIncomeNotesByIncomeId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.IncomeNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.IncomeNote{
			ID:          note.ID,
			AuthorId:    note.AuthorId,
			AuthorName:  note.AuthorName,
			AuthorEmail: note.AuthorEmail,
			Income:      note.Income,
			Title:       note.Title,
			Note:        note.Note,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched income by income id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF GET ALL EXPENSE NOTES (incomeId) -----
// ------------------------------------------------

// ------------------------------------------------
// -- START OF GET INCOME NOTE BY ID --------------
// ------------------------------------------------

func (app *Config) GetIncomeNoteById(w http.ResponseWriter, r *http.Request) {
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

	note, err := app.Models.IncomeNote.GetIncomeNoteById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get income note by id"), http.StatusBadRequest)
		return
	}

	returnedNote := data.IncomeNote{
		ID:          note.ID,
		AuthorId:    note.AuthorId,
		AuthorName:  note.AuthorName,
		AuthorEmail: note.AuthorEmail,
		Income:      note.Income,
		Title:       note.Title,
		Note:        note.Note,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched income note: %s", note.Title),
		Data:    returnedNote,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF GET Income NOTE BY ID ---------------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF UPDATE Income NOTE ----------------
// ------------------------------------------------

func (app *Config) UpdateIncomeNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateIncomeNotePayload

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

	returnedNote := data.IncomeNote{
		ID:          requestPayload.ID,
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		Income:      requestPayload.Income,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	err = returnedNote.UpdateIncomeNote()
	if err != nil {
		app.errorJSON(w, errors.New("could not update Income note: "+err.Error()), http.StatusBadRequest)
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
// -- END OF UPDATE INCOME NOTE ------------------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF DELETE INCOME NOTE ----------------
// ------------------------------------------------

func (app *Config) DeleteIncomeNote(w http.ResponseWriter, r *http.Request) {
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

	err = data.DeleteIncomeNote(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted income note with Id: %s", fmt.Sprint(requestPayload.ID)),
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF DELETE INCOME NOTE ------------------
// ------------------------------------------------
