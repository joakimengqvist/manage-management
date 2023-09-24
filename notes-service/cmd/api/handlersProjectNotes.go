package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"notes-service/cmd/data"
)

type NewNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	Project     string `json:"project"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateNote struct {
	Id          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	Project     string `json:"project"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type NoteIdPayload struct {
	Id string `json:"id"`
}

type ProjectId struct {
	ProjectId string `json:"projectId"`
}

type ReturnedNotes struct {
	Notes []data.Note `json:"notes"`
}

type UpdateProjectNote struct {
	NoteId    string `json:"noteId"`
	ProjectId string `json:"projectId"`
}

func (app *Config) CreateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewNote

	// userId := r.Header.Get("X-User-Id")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newNote := data.Note{
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		Project:     requestPayload.Project,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	response, err := app.Models.Note.InsertProjectNote(newNote)
	if err != nil {
		app.errorJSON(w, errors.New("could not create project note: "+err.Error()), http.StatusBadRequest)
		return
	}

	updateProject := UpdateProjectNote{
		ProjectId: requestPayload.Project,
		NoteId:    response,
	}

	jsonData, _ := json.MarshalIndent(updateProject, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/update-project-notes", bytes.NewBuffer(jsonData))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}

	projectUpdateResponse, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		// delete project note here
		return
	}

	defer projectUpdateResponse.Body.Close()

	if projectUpdateResponse.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - update project with new project note"))
		// delete project note here
		return
	} else if projectUpdateResponse.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling project service - update project with new project note"))
		// delete project note here
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("created note %s", requestPayload.Title),
		Data:    response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllNotesByProductId(w http.ResponseWriter, r *http.Request) {

	var requestPayload ProjectId

	// userId := r.Header.Get("X-User-Id")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	notes, err := app.Models.Note.GetProjectNotesByProjectId(requestPayload.ProjectId)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.Note
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.Note{
			ID:          note.ID,
			AuthorId:    note.AuthorId,
			AuthorName:  note.AuthorName,
			AuthorEmail: note.AuthorEmail,
			Project:     note.Project,
			Title:       note.Title,
			Note:        note.Note,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	app.logItemViaRPC(w, nil, RPCLogData{Action: "Get all project notes by project id [/notes/get-all-notes-by-project-id]", Name: "[notes-service] - Successfuly fetched all notes by project id"})
	app.writeJSON(w, http.StatusAccepted, noteSlice)
}

func (app *Config) GetProjectNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload NoteIdPayload

	// userId := r.Header.Get("X-User-Id")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.Note.GetProjectNoteById(requestPayload.Id)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get project note by id"), http.StatusBadRequest)
		return
	}

	returnedNote := data.Note{
		ID:          note.ID,
		AuthorId:    note.AuthorId,
		AuthorName:  note.AuthorName,
		AuthorEmail: note.AuthorEmail,
		Project:     note.Project,
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

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get project note by id [/notes/get-project-note-by-id]", Name: "[notes-service] - Successfuly fetched project note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateNote

	//VuserId := r.Header.Get("X-User-Id")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	returnedNote := data.Note{
		ID:          requestPayload.Id,
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		Project:     requestPayload.Project,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	err = returnedNote.UpdateProjectNote()
	if err != nil {
		app.errorJSON(w, errors.New("could not update project note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated note with Id: %s", fmt.Sprint(returnedNote.ID)),
		Data:    returnedNote,
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Authenticate [/notes/update-project-note]", Name: "[notes-service] - Successful updated project-note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}
