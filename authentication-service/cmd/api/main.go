package main

import (
	"authentication/cmd/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type IDpayload struct {
	ID string `json:"id"`
}

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("[ -- Starting authentication service -- ]")

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
		log.Println("sql.Open - openDB", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("db.Ping - openDB", err)
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
