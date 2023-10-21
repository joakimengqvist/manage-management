package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// -------------------------------------------
// ------- START OF SEND EMAIL  --------------
// -------------------------------------------

func (app *Config) SendEmail(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Send email [/email/send]", Name: "[broker-service] - Send email initiated"})

	app.sendMail(w, requestPayload)
}

func (app *Config) sendMail(w http.ResponseWriter, rpl RequestPayload) {
	jsonData, _ := json.MarshalIndent(rpl.Mail, "", "")

	app.logItemViaRPC(w, jsonData, RPCLogData{Action: "Send email [/email/send]", Name: "[broker-service]"})

	mailServiceURL := "http://mail-service/send"

	request, err := http.NewRequest("POST", mailServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	}

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.logItemViaRPC(w, jsonFromService, RPCLogData{Action: "Send email success [/email/send]", Name: "[broker-service] - succesfully sent email"})
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}

// -------------------------------------------
// ------- END OF SEND EMAIL  ----------------
// -------------------------------------------
