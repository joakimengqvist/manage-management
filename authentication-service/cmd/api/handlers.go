package main

import (
	"encoding/json"
	"net/http"
	"net/rpc"
)

type IDpayload struct {
	ID string `json:"id"`
}

type RPCLogData struct {
	Action string
	Name   string
}

type RPCPayload struct {
	Action string
	Name   string
	Data   string
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
