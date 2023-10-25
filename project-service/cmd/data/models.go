package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Project:    Project{},
		SubProject: SubProject{},
	}
}

type Models struct {
	Project    Project
	SubProject SubProject
}

type NewProject struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Notes       []string  `json:"notes"`
	SubProjects []string  `json:"sub_projects"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
}

type PostgresProject struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Notes       string    `json:"notes"`
	SubProjects string    `json:"sub_projects"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
}

type NewSubProject struct {
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Priority          int       `json:"priority"`
	StartDate         time.Time `json:"start_date"`
	DueDate           time.Time `json:"due_date"`
	EstimatedDuration int       `json:"estimated_duration"`
	Notes             []string  `json:"notes"`
	CreatedBy         string    `json:"created_by"`
	UpdatedBy         string    `json:"updated_by"`
	Invoices          []string  `json:"invoices"`
	Incomes           []string  `json:"incomes"`
	Expenses          []string  `json:"expenses"`
}

type SubProject struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Priority          int       `json:"priority"`
	StartDate         time.Time `json:"start_date"`
	DueDate           time.Time `json:"due_date"`
	EstimatedDuration int       `json:"estimated_duration"`
	Notes             []string  `json:"notes"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedBy         string    `json:"updated_by"`
	Projects          []string  `json:"projects"`
	Invoices          []string  `json:"invoices"`
	Incomes           []string  `json:"incomes"`
	Expenses          []string  `json:"expenses"`
}

type PostgresSubProject struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Priority          int       `json:"priority"`
	StartDate         time.Time `json:"start_date"`
	DueDate           time.Time `json:"due_date"`
	EstimatedDuration int       `json:"estimated_duration"`
	Notes             string    `json:"notes"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedBy         string    `json:"updated_by"`
	Projects          string    `json:"projects"`
	Invoices          string    `json:"invoices"`
	Incomes           string    `json:"incomes"`
	Expenses          string    `json:"expenses"`
}
