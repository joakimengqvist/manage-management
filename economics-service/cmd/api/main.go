package main

import (
	"bytes"
	"database/sql"
	"economics-service/cmd/data"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("Starting project service]")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Unable to connect to postgress.")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgress not yet ready...")
			counts++
		} else {
			log.Println("Connected to postress.")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for 2 seconds")
		time.Sleep(2 * time.Second)
	}
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
