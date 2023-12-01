package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"notes-service/cmd/data"
)

type NewSubProjectNote struct {
	AuthorId     string `json:"author_id"`
	AuthorName   string `json:"author_name"`
	AuthorEmail  string `json:"author_email"`
	SubProjectId string `json:"sub_project_id"`
	Title        string `json:"title"`
	Note         string `json:"note"`
}

type UpdateSubProjectNotePayload struct {
	ID           string `json:"id"`
	AuthorId     string `json:"author_id"`
	AuthorName   string `json:"author_name"`
	AuthorEmail  string `json:"author_email"`
	SubProjectId string `json:"sub_project_id"`
	Title        string `json:"title"`
	Note         string `json:"note"`
}

type ReturnedSubProjectNotes struct {
	Notes []data.SubProjectNote `json:"notes"`
}

type UpdateSubProjectNote struct {
	NoteId       string `json:"note_id"`
	SubProjectId string `json:"sub_project_id"`
}

type DeleteSubProjectNoteIdPayload struct {
	AuthorId     string `json:"author_id"`
	NoteId       string `json:"note_id"`
	SubProjectId string `json:"sub_project_id"`
}

func (app *Config) CreateSubProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewSubProjectNote

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

	newNote := data.SubProjectNote{
		AuthorId:     requestPayload.AuthorId,
		AuthorName:   requestPayload.AuthorName,
		AuthorEmail:  requestPayload.AuthorEmail,
		SubProjectId: requestPayload.SubProjectId,
		Title:        requestPayload.Title,
		Note:         requestPayload.Note,
	}

	fmt.Println("creating note", err)

	noteId, err := app.Models.SubProjectNote.InsertSubProjectNote(newNote)
	if err != nil {
		fmt.Println("error note id", err)
		app.errorJSON(w, errors.New("could not create subProject note: "+err.Error()), http.StatusBadRequest)
		return
	}

	updateSubProject := UpdateSubProjectNote{
		SubProjectId: requestPayload.SubProjectId,
		NoteId:       noteId,
	}

	jsonDataSubProject, _ := json.MarshalIndent(updateSubProject, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/add-sub-project-note", bytes.NewBuffer(jsonDataSubProject))

	if err != nil {
		fmt.Println("error note id", err)
		app.errorJSON(w, err)
		data.DeleteSubProjectNote(noteId)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	subProjectUpdateResponse, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		data.DeleteSubProjectNote(noteId)
		return
	}

	defer subProjectUpdateResponse.Body.Close()

	if subProjectUpdateResponse.StatusCode == http.StatusUnauthorized {
		data.DeleteSubProjectNote(noteId)
		return
	} else if subProjectUpdateResponse.StatusCode != http.StatusAccepted {
		data.DeleteSubProjectNote(noteId)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("created note %s", requestPayload.Title),
		Data:    noteId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllSubProjectNotesBySubProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Println("Readjson", err)
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

	notes, err := app.Models.SubProjectNote.GetSubProjectNotesBySubProjectId(requestPayload.ID)
	if err != nil {
		fmt.Print("GetSubProjectNotesBySubProjectId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.SubProjectNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.SubProjectNote{
			ID:           note.ID,
			AuthorId:     note.AuthorId,
			AuthorName:   note.AuthorName,
			AuthorEmail:  note.AuthorEmail,
			SubProjectId: note.SubProjectId,
			Title:        note.Title,
			Note:         note.Note,
			CreatedAt:    note.CreatedAt,
			UpdatedAt:    note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all subProject notes by subProject id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllSubProjectNotesByUserId(w http.ResponseWriter, r *http.Request) {

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

	notes, err := app.Models.SubProjectNote.GetSubProjectNotesByAuthorId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.SubProjectNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.SubProjectNote{
			ID:           note.ID,
			AuthorId:     note.AuthorId,
			AuthorName:   note.AuthorName,
			AuthorEmail:  note.AuthorEmail,
			SubProjectId: note.SubProjectId,
			Title:        note.Title,
			Note:         note.Note,
			CreatedAt:    note.CreatedAt,
			UpdatedAt:    note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all subProject notes by user id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetSubProjectNoteById(w http.ResponseWriter, r *http.Request) {
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

	note, err := app.Models.SubProjectNote.GetSubProjectNoteById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get subProject note by id"), http.StatusBadRequest)
		return
	}

	returnedNote := data.SubProjectNote{
		ID:           note.ID,
		AuthorId:     note.AuthorId,
		AuthorName:   note.AuthorName,
		AuthorEmail:  note.AuthorEmail,
		SubProjectId: note.SubProjectId,
		Title:        note.Title,
		Note:         note.Note,
		CreatedAt:    note.CreatedAt,
		UpdatedAt:    note.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched note: %s", note.Title),
		Data:    returnedNote,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateSubProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateSubProjectNotePayload

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

	returnedNote := data.SubProjectNote{
		ID:           requestPayload.ID,
		AuthorId:     requestPayload.AuthorId,
		AuthorName:   requestPayload.AuthorName,
		AuthorEmail:  requestPayload.AuthorEmail,
		SubProjectId: requestPayload.SubProjectId,
		Title:        requestPayload.Title,
		Note:         requestPayload.Note,
	}

	err = returnedNote.UpdateSubProjectNote()
	if err != nil {
		app.errorJSON(w, errors.New("could not update subProject note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated note with Id: %s", fmt.Sprint(returnedNote.ID)),
		Data:    returnedNote,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteSubProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteSubProjectNoteIdPayload

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

	note, err := app.Models.SubProjectNote.GetSubProjectNoteById(requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get subProject note by id"), http.StatusBadRequest)
		return
	}

	err = app.RemoveNoteFromSubProject(w, r, note.ID, requestPayload.SubProjectId)
	if err != nil {
		fmt.Println("error removing note from subProject: ", err)
		app.errorJSON(w, err)
		return
	}

	err = data.DeleteSubProjectNote(requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted note with Id: %s", fmt.Sprint(requestPayload.NoteId)),
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) RemoveNoteFromSubProject(w http.ResponseWriter, r *http.Request, noteId string, subProjectId string) error {
	deleteSubProjectNote := UpdateSubProjectNote{
		NoteId:       noteId,
		SubProjectId: subProjectId,
	}

	fmt.Println("deleteSubProjectNote: ", deleteSubProjectNote)

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

	jsonDataUser, _ := json.MarshalIndent(deleteSubProjectNote, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/delete-sub-project-note", bytes.NewBuffer(jsonDataUser))

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
		app.errorJSON(w, errors.New("status unauthorized - delete note from subProject"))
		return err
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling subProject service - delete note from subProject"))
		return err
	}

	return nil
}
