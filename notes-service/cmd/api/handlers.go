package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/rpc"
)

type IDpayload struct {
	ID string `json:"id"`
}

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
		return false, err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		return false, nil
	} else if response.StatusCode != http.StatusAccepted {
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

func (app *Config) logItemViaRPC(w http.ResponseWriter, payload any, logData RPCLogData) {

	jsonData, _ := json.MarshalIndent(payload, "", "")

	client, err := rpc.Dial("tcp", "logger-service:5001")
	if err != nil {
		return
	}

	rpcPayload := RPCPayload{
		Action: logData.Action,
		Name:   logData.Name,
		Data:   string(jsonData),
	}

	err = client.Call("RPCServer.LogInfo", rpcPayload, nil)
	if err != nil {
		return
	}
}
