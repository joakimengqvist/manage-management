package main

import (
	"encoding/json"
	"net/http"
	"net/rpc"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Send email [/email/send]", Name: "[mail-service] - Failed to read JSON payload"})
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = app.Mailer.sendSMTPMessage(msg)
	if err != nil {
		app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Send email [/email/send]", Name: "[mail-service] - Failed to send SMTP message (mail)"})
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Email is sent to " + requestPayload.To,
	}

	app.logItemViaRPC(w, requestPayload, RPCLogData{Action: "Send email success [/email/send]", Name: "[mail-service] - Successfully sent mail"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

type RPCPayload struct {
	Action string
	Name   string
	Data   string
}

type RPCLogData struct {
	Action string
	Name   string
}

func (app *Config) logItemViaRPC(w http.ResponseWriter, payload any, logData RPCLogData) {

	jsonData, _ := json.MarshalIndent(payload, "", "")

	client, err := rpc.Dial("tcp", "logger-service:5001")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	rpcPayload := RPCPayload{
		Action: logData.Action,
		Name:   logData.Name,
		Data:   string(jsonData),
	}

	err = client.Call("RPCServer.LogInfo", rpcPayload, nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
