package main

import (
	"net/http"
)

type IDpayload struct {
	ID string `json:"id"`
}

type IDSpayload struct {
	IDs []string `json:"ids"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	app.writeJSON(w, http.StatusOK, payload)

}
