package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"notes-service/cmd/data"
)

type NewProjectNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	Project     string `json:"project"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateProjectNotePayload struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	Project     string `json:"project"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type IDpayload struct {
	ID string `json:"id"`
}

type ReturnedProjectNotes struct {
	Notes []data.ProjectNote `json:"notes"`
}

type UpdateProjectNote struct {
	NoteId    string `json:"note_id"`
	ProjectId string `json:"project_id"`
}

type UpdateUserNotes struct {
	NoteId string `json:"note_id"`
	UserId string `json:"user_id"`
}

type DeleteNoteIdPayload struct {
	AuthorId  string `json:"author_id"`
	NoteId    string `json:"note_id"`
	ProjectId string `json:"project_id"`
}

// -------------------------------------------
// ------- START OF CREATE PROJECT NOTE  -----
// -------------------------------------------

func (app *Config) CreateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewProjectNote

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

	newNote := data.ProjectNote{
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		Project:     requestPayload.Project,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	noteId, err := app.Models.ProjectNote.InsertProjectNote(newNote)
	if err != nil {
		app.errorJSON(w, errors.New("could not create project note: "+err.Error()), http.StatusBadRequest)
		return
	}

	updateProject := UpdateProjectNote{
		ProjectId: requestPayload.Project,
		NoteId:    noteId,
	}

	jsonDataProject, _ := json.MarshalIndent(updateProject, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/add-project-note", bytes.NewBuffer(jsonDataProject))

	if err != nil {
		app.errorJSON(w, err)
		data.DeleteProjectNote(noteId)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	projectUpdateResponse, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		data.DeleteProjectNote(noteId)
		return
	}

	defer projectUpdateResponse.Body.Close()

	if projectUpdateResponse.StatusCode == http.StatusUnauthorized {
		data.DeleteProjectNote(noteId)
		return
	} else if projectUpdateResponse.StatusCode != http.StatusAccepted {
		data.DeleteProjectNote(noteId)
		return
	}

	updateUser := UpdateUserNotes{
		NoteId: noteId,
		UserId: requestPayload.AuthorId,
	}

	jsonDataUser, _ := json.MarshalIndent(updateUser, "", "")

	request, err = http.NewRequest("POST", "http://authentication-service/auth/add-user-note", bytes.NewBuffer(jsonDataUser))

	if err != nil {
		data.DeleteProjectNote(requestPayload.Project)
		app.RemoveNoteFromProject(w, r, requestPayload.AuthorId, requestPayload.Project)
		return
	}

	client = &http.Client{}

	userUpdateResponse, err := client.Do(request)
	if err != nil {
		data.DeleteProjectNote(noteId)
		app.RemoveNoteFromProject(w, r, requestPayload.AuthorId, requestPayload.Project)
		return
	}

	defer userUpdateResponse.Body.Close()

	if userUpdateResponse.StatusCode == http.StatusUnauthorized {
		data.DeleteProjectNote(noteId)
		app.RemoveNoteFromProject(w, r, requestPayload.AuthorId, requestPayload.Project)
		return
	} else if userUpdateResponse.StatusCode != http.StatusAccepted {
		data.DeleteProjectNote(noteId)
		app.RemoveNoteFromProject(w, r, requestPayload.AuthorId, requestPayload.Project)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("created note %s", requestPayload.Title),
		Data:    noteId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF CREATE PROJECT NOTE  -----------------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF GET ALL PROJECT NOTES (projectId) --
// ------------------------------------------------

func (app *Config) GetAllProjectNotesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	fmt.Println("GetAllNotesByProjectId")

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

	notes, err := app.Models.ProjectNote.GetProjectNotesByProjectId(requestPayload.ID)
	if err != nil {
		fmt.Print("notes", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.ProjectNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.ProjectNote{
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

	app.logItemViaRPC(w, noteSlice, RPCLogData{Action: "Get all project notes by project id [/notes/get-all-notes-by-project-id]", Name: "[notes-service] - Successfuly fetched all notes by project id"})
	app.writeJSON(w, http.StatusAccepted, noteSlice)
}

// ------------------------------------------------
// -- END OF GET ALL PROJECT NOTES (projectId) ----
// ------------------------------------------------

// ------------------------------------------------
// -- START OF GET ALL PROJECT NOTES (userId) -----
// ------------------------------------------------

func (app *Config) GetAllProjectNotesByUserId(w http.ResponseWriter, r *http.Request) {

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

	notes, err := app.Models.ProjectNote.GetProjectNotesByAuthorId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var noteSlice []data.ProjectNote
	for _, notePtr := range notes {
		note := *notePtr

		returnedNote := data.ProjectNote{
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

	app.logItemViaRPC(w, noteSlice, RPCLogData{Action: "Get all project notes by project id [/notes/get-all-notes-by-project-id]", Name: "[notes-service] - Successfuly fetched all notes by project id"})
	app.writeJSON(w, http.StatusAccepted, noteSlice)
}

// ------------------------------------------------
// -- END OF GET ALL PROJECT NOTES (userId) -------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF GET PROJECT NOTE BY ID -------------
// ------------------------------------------------

func (app *Config) GetProjectNoteById(w http.ResponseWriter, r *http.Request) {
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

	note, err := app.Models.ProjectNote.GetProjectNoteById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get project note by id"), http.StatusBadRequest)
		return
	}

	returnedNote := data.ProjectNote{
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

// ------------------------------------------------
// -- END OF GET PROJECT NOTE BY ID ---------------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF UPDATE PROJECT NOTE ----------------
// ------------------------------------------------

func (app *Config) UpdateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProjectNotePayload

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

	returnedNote := data.ProjectNote{
		ID:          requestPayload.ID,
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

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Project notes [/notes/update-project-note]", Name: "[notes-service] - Successful updated project-note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF UPDATE PROJECT NOTE ------------------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF DELETE PROJECT NOTE ----------------
// ------------------------------------------------

func (app *Config) DeleteProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteNoteIdPayload

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

	note, err := app.Models.ProjectNote.GetProjectNoteById(requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get project note by id"), http.StatusBadRequest)
		return
	}

	deleteUserNote := DeleteNoteIdPayload{
		NoteId:   note.ID,
		AuthorId: requestPayload.AuthorId,
	}

	jsonDataUser, _ := json.MarshalIndent(deleteUserNote, "", "")

	request, err := http.NewRequest("POST", "http://authentication-service/auth/delete-user-note", bytes.NewBuffer(jsonDataUser))

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
		app.errorJSON(w, errors.New("status unauthorized - delete project note from user"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service - delete project note from user"))
		return
	}

	err = app.RemoveNoteFromProject(w, r, note.ID, requestPayload.ProjectId)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = data.DeleteProjectNote(requestPayload.NoteId)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted note with Id: %s", fmt.Sprint(requestPayload.NoteId)),
		Data:    nil,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete project note [/auth/delete-user-note]", Name: "[authentication-service] - Successful deleted user note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

// ------------------------------------------------
// -- END OF DELETE PROJECT NOTE ------------------
// ------------------------------------------------

// ------------------------------------------------
// -- START OF REMOVE PROJECT NOTE FROM PROJECT ---
// ------------------------------------------------

func (app *Config) RemoveNoteFromProject(w http.ResponseWriter, r *http.Request, noteId string, projectId string) error {
	deleteProjectNote := UpdateProjectNote{
		NoteId:    noteId,
		ProjectId: projectId,
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

	jsonDataUser, _ := json.MarshalIndent(deleteProjectNote, "", "")

	request, err := http.NewRequest("POST", "http://project-service/project/delete-project-note", bytes.NewBuffer(jsonDataUser))

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
		app.errorJSON(w, errors.New("status unauthorized - delete note from project"))
		return err
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling project service - delete note from project"))
		return err
	}

	return nil
}

// ------------------------------------------------
// -- END OF REMOVE PROJECT NOTE FROM PROJECT -----
// ------------------------------------------------
