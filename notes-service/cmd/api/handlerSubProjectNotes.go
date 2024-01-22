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
		log.Println("authenticated - CreateSubProjectNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("authenticated - CreateSubProjectNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateSubProjectNote", err)
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

	noteId, err := app.Models.SubProjectNote.InsertSubProjectNote(newNote)
	if err != nil {
		log.Println("postgres - CreateSubProjectNote", err)
		app.errorJSON(w, errors.New("could not create subProject note: "+err.Error()), http.StatusBadRequest)
		return
	}

	updateSubProject := UpdateSubProjectNote{
		SubProjectId: requestPayload.SubProjectId,
		NoteId:       noteId,
	}

	jsonDataSubProject, _ := json.MarshalIndent(updateSubProject, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/add-sub-project-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonDataSubProject))

	if err != nil {
		log.Println("Post - CreateSubProjectNote", err)
		app.errorJSON(w, err)
		data.DeleteSubProjectNote(noteId)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	subProjectUpdateResponse, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - CreateSubProjectNote", err)
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
		log.Println("readJSON - GetAllSubProjectNotesBySubProjectId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetAllSubProjectNotesBySubProjectId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllSubProjectNotesBySubProjectId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	notes, err := app.Models.SubProjectNote.GetSubProjectNotesBySubProjectId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetAllSubProjectNotesBySubProjectId", err)
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
		log.Println("authenticated - GetAllSubProjectNotesByUserId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllSubProjectNotesByUserId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllSubProjectNotesByUserId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	notes, err := app.Models.SubProjectNote.GetSubProjectNotesByAuthorId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetAllSubProjectNotesByUserId", err)
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
		log.Println("authenticated - GetSubProjectNoteById", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetSubProjectNoteById")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetSubProjectNoteById", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.SubProjectNote.GetSubProjectNoteById(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetSubProjectNoteById", err)
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
		log.Println("authenticated - UpdateSubProjectNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - UpdateSubProjectNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateSubProjectNote", err)
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
		log.Println("postgres - UpdateSubProjectNote", err)
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
		log.Println("authenticated - DeleteSubProjectNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - DeleteSubProjectNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - DeleteSubProjectNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.SubProjectNote.GetSubProjectNoteById(requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - DeleteSubProjectNote", err)
		app.errorJSON(w, errors.New("failed to get subProject note by id"), http.StatusBadRequest)
		return
	}

	err = app.RemoveNoteFromSubProject(w, r, note.ID, requestPayload.SubProjectId)
	if err != nil {
		log.Println("RemoveNoteFromSubProject - DeleteSubProjectNote", err)
		app.errorJSON(w, err)
		return
	}

	err = data.DeleteSubProjectNote(requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - DeleteSubProjectNote", err)
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

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_sudo")
	if err != nil {
		log.Println("authenticated - RemoveNoteFromSubProject", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return err
	}

	if !authenticated {
		log.Println("!authenticated - RemoveNoteFromSubProject")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	jsonDataUser, _ := json.MarshalIndent(deleteSubProjectNote, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/delete-sub-project-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonDataUser))
	if err != nil {
		log.Println("POST - RemoveNoteFromSubProject", err)
		app.errorJSON(w, err)
		return err
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - RemoveNoteFromSubProject", err)
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
