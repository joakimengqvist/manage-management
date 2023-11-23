package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {

	out, err := json.Marshal(data)
	if err != nil {
		log.Println("JSONMARSHAL", err)
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		log.Println("WRITE HEADER", err)
		return err
	}

	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)
}

func (app *Config) parsePostgresArray(postgresArray string) []string {
	postgresArray = strings.Trim(postgresArray, "{}")

	if len(postgresArray) < 4 {
		return []string{}
	}

	arrayElements := strings.Split(postgresArray, ",")

	return arrayElements
}

func (app *Config) convertToPostgresArray(arrayElements []string) string {
	postgresArray := strings.Join(arrayElements, ",")
	postgresArray = "{" + postgresArray + "}"
	return postgresArray
}

func (app *Config) convertToPostgresArrayInteger(arrayElements []int) string {
	arrayString := ""
	for i, id := range arrayElements {
		arrayString += strconv.Itoa(id)
		if i < (len(arrayElements) - 1) {
			arrayString += ","
		}
	}
	postgresArray := "{" + arrayString + "}"
	return postgresArray
}

func (app *Config) parsePostgresArrayInteger(postgresArray string) []int {
	postgresArray = strings.Trim(postgresArray, "{}")
	arrayElements := strings.Split(postgresArray, ",")

	var arrayInt []int
	for _, id := range arrayElements {
		intId, _ := strconv.Atoi(id)
		arrayInt = append(arrayInt, intId)
	}

	return arrayInt
}
