package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
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

	endpoint := "http://" + os.Getenv("AUTHENTICATION_SERVICE_SERVICE_HOST") + "/auth/check-privilege"

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))

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
