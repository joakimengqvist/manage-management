package main

import (
	"log-service/data"
	"net/http"
)

type logJsonResponse struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	Data   string `json:"data"`
}

// Function is not being used
func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload logJsonResponse

	_ = app.readJSON(w, r, &requestPayload)

	event := data.LogEntry{
		Action: requestPayload.Action,
		Name:   requestPayload.Name,
		Data:   requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(&event)
	if err != nil {
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}
