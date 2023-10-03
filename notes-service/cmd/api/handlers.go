package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/rpc"
)

type PrivilegeCheckPayload struct {
	UserId string `json:"userId"`
	Action string `json:"action"`
}

func (app *Config) CheckPrivilege(w http.ResponseWriter, userId string, privilege string) (bool, error) {

	payload := PrivilegeCheckPayload{
		UserId: userId,
		Action: privilege,
	}

	jsonData, _ := json.MarshalIndent(payload, "", "")

	request, err := http.NewRequest("POST", "http://authentication-service/auth/check-privilege", bytes.NewBuffer(jsonData))

	if err != nil {
		app.errorJSON(w, err)
		return false, err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return false, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - check privilege project "+payload.Action))
		return false, nil
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling authentication service - check privilege project "+payload.Action))
		return false, nil
	}

	return true, nil
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

// -------------------------------------------
// ------- START OF LOG ITEM (RPC)  ----------
// -------------------------------------------

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

// -------------------------------------------
// ------- END OF LOG ITEM (RPC)  ------------
// -------------------------------------------
