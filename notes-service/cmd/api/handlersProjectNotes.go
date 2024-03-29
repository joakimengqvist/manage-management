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

type NewProjectNote struct {
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	ProjectId   string `json:"project_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type UpdateProjectNotePayload struct {
	ID          string `json:"id"`
	AuthorId    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	ProjectId   string `json:"project_id"`
	Title       string `json:"title"`
	Note        string `json:"note"`
}

type ReturnedProjectNotes struct {
	Notes []data.ProjectNote `json:"notes"`
}

type UpdateProjectNote struct {
	NoteId    string `json:"note_id"`
	ProjectId string `json:"project_id"`
}

type DeleteProjectNoteIdPayload struct {
	AuthorId  string `json:"author_id"`
	NoteId    string `json:"note_id"`
	ProjectId string `json:"project_id"`
}

func (app *Config) CreateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload NewProjectNote

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_write")
	if err != nil {
		log.Println("authenticated - CreateProjectNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - CreateProjectNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateProjectNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newNote := data.ProjectNote{
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		ProjectId:   requestPayload.ProjectId,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	noteId, err := app.Models.ProjectNote.InsertProjectNote(newNote)
	if err != nil {
		log.Println("postgres - CreateProjectNote", err)
		app.errorJSON(w, errors.New("could not create project note: "+err.Error()), http.StatusBadRequest)
		return
	}

	updateProject := UpdateProjectNote{
		ProjectId: requestPayload.ProjectId,
		NoteId:    noteId,
	}

	jsonDataProject, _ := json.MarshalIndent(updateProject, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/add-project-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonDataProject))

	if err != nil {
		log.Println("POST - CreateProjectNote", err)
		app.errorJSON(w, err)
		data.DeleteProjectNote(noteId)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	projectUpdateResponse, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - CreateProjectNote", err)
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

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("created note %s", requestPayload.Title),
		Data:    noteId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllProjectNotesByProjectId(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	log.Println("GetAllNotesByProjectId")

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllProjectNotesByProjectId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetAllProjectNotesByProjectId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllProjectNotesByProjectId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	notes, err := app.Models.ProjectNote.GetProjectNotesByProjectId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetAllProjectNotesByProjectId", err)
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
			ProjectId:   note.ProjectId,
			Title:       note.Title,
			Note:        note.Note,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all project notes by project id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllProjectNotesByUserId(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetAllProjectNotesByUserId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllProjectNotesByUserId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllProjectNotesByUserId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	notes, err := app.Models.ProjectNote.GetProjectNotesByAuthorId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetAllProjectNotesByUserId", err)
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
			ProjectId:   note.ProjectId,
			Title:       note.Title,
			Note:        note.Note,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}

		noteSlice = append(noteSlice, returnedNote)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "get all project notes by user id",
		Data:    noteSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetProjectNoteById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - GetProjectNoteById", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetProjectNoteById")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetProjectNoteById", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.ProjectNote.GetProjectNoteById(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetProjectNoteById", err)
		app.errorJSON(w, errors.New("failed to get project note by id"), http.StatusBadRequest)
		return
	}

	returnedNote := data.ProjectNote{
		ID:          note.ID,
		AuthorId:    note.AuthorId,
		AuthorName:  note.AuthorName,
		AuthorEmail: note.AuthorEmail,
		ProjectId:   note.ProjectId,
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

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Get project note by id [/notes/get-project-note-by-id]", Name: "[notes-service] - Successfuly fetched project note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateProjectNotePayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_read")
	if err != nil {
		log.Println("authenticated - UpdateProjectNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - UpdateProjectNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateProjectNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	returnedNote := data.ProjectNote{
		ID:          requestPayload.ID,
		AuthorId:    requestPayload.AuthorId,
		AuthorName:  requestPayload.AuthorName,
		AuthorEmail: requestPayload.AuthorEmail,
		ProjectId:   requestPayload.ProjectId,
		Title:       requestPayload.Title,
		Note:        requestPayload.Note,
	}

	err = returnedNote.UpdateProjectNote()
	if err != nil {
		log.Println("postgres - UpdateProjectNote", err)
		app.errorJSON(w, errors.New("could not update project note: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated note with Id: %s", fmt.Sprint(returnedNote.ID)),
		Data:    returnedNote,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Project notes [/notes/update-project-note]", Name: "[notes-service] - Successful updated project-note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteProjectNote(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteProjectNoteIdPayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_sudo")
	if err != nil {
		log.Println("authenticated - DeleteProjectNote", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - DeleteProjectNote")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - DeleteProjectNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	note, err := app.Models.ProjectNote.GetProjectNoteById(requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - DeleteProjectNote", err)
		app.errorJSON(w, errors.New("failed to get project note by id"), http.StatusBadRequest)
		return
	}

	err = app.RemoveNoteFromProject(w, r, note.ID, requestPayload.ProjectId)
	if err != nil {
		log.Println("RemoveNoteFromProject - DeleteProjectNote", err)
		app.errorJSON(w, err)
		return
	}

	err = data.DeleteProjectNote(requestPayload.NoteId)
	if err != nil {
		log.Println("postgres - DeleteProjectNote", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("deleted note with Id: %s", fmt.Sprint(requestPayload.NoteId)),
		Data:    nil,
	}

	// app.logItemViaRPC(w, payload, RPCLogData{Action: "Delete project note [/auth/delete-user-note]", Name: "[authentication-service] - Successful deleted user note"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) RemoveNoteFromProject(w http.ResponseWriter, r *http.Request, noteId string, projectId string) error {
	deleteProjectNote := UpdateProjectNote{
		NoteId:    noteId,
		ProjectId: projectId,
	}

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "note_sudo")
	if err != nil {
		log.Println("authenticated - RemoveNoteFromProject", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return err
	}

	if !authenticated {
		log.Println("!authenticated - RemoveNoteFromProject")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return errors.New("status unauthorized")
	}

	jsonDataUser, _ := json.MarshalIndent(deleteProjectNote, "", "")

	endpoint := "http://" + os.Getenv("PROJECT_SERVICE_SERVICE_HOST") + "/project/delete-project-note"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonDataUser))

	if err != nil {
		log.Println("POST - RemoveNoteFromProject", err)
		app.errorJSON(w, err)
		return err
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.Do - RemoveNoteFromProject", err)
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
