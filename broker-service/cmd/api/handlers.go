package main

import (
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
}

type IDpayload struct {
	ID string `json:"id"`
}

type IDSpayload struct {
	IDs []string `json:"ids"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker yeah",
	}

	app.writeJSON(w, http.StatusOK, payload)

}
