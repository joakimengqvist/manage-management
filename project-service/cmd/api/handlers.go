package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type IDpayload struct {
	ID string `json:"id"`
}

type IdsPayload struct {
	Ids []string `json:"ids"`
}

type PrivilegeCheckPayload struct {
	UserId string `json:"userId"`
	Action string `json:"action"`
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

func (app *Config) CheckPrivilege(w http.ResponseWriter, userId string, privilege string) (bool, error) {

	payload := PrivilegeCheckPayload{
		UserId: userId,
		Action: privilege,
	}

	jsonData, _ := json.MarshalIndent(payload, "", "")

	endpoint := "http://" + os.Getenv("AUTHENTICATION_SERVICE_SERVICE_HOST") + "/auth/check-privilege"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Println("POST - CheckPrivilege", err)
		app.errorJSON(w, err)
		return false, err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("client.do - CheckPrivilege", err)
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

/*

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

*/
