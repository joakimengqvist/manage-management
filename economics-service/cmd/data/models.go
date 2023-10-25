package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Expense:    Expense{},
		NewExpense: NewExpense{},
		Income:     Income{},
		NewIncome:  NewIncome{},
	}
}

type Models struct {
	Expense    Expense
	NewExpense NewExpense
	Income     Income
	NewIncome  NewIncome
}

type NewExpense struct {
	ProjectID       string    `json:"project_id"`
	ExpenseDate     time.Time `json:"expense_date"`
	ExpenseCategory string    `json:"expense_category"`
	Vendor          string    `json:"vendor"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Tax             float64   `json:"tax"`
	Status          string    `json:"status"`
	Currency        string    `json:"currency"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedBy       string    `json:"created_by"`
	UpdatedBy       string    `json:"updated_by"`
}

type NewIncome struct {
	ProjectID      string    `json:"project_id"`
	IncomeDate     time.Time `json:"income_date"`
	IncomeCategory string    `json:"income_category"`
	Vendor         string    `json:"vendor"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Tax            float64   `json:"tax"`
	Status         string    `json:"status"`
	Currency       string    `json:"currency"`
	PaymentMethod  string    `json:"payment_method"`
	CreatedBy      string    `json:"created_by"`
	UpdatedBy      string    `json:"updated_by"`
}

type Expense struct {
	ID              string    `json:"id"`
	ProjectID       string    `json:"project_id"`
	ExpenseDate     time.Time `json:"expense_date"`
	ExpenseCategory string    `json:"expense_category"`
	Vendor          string    `json:"vendor"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Tax             float64   `json:"tax"`
	Status          string    `json:"status"`
	Currency        string    `json:"currency"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedBy       string    `json:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Income struct {
	ID             string    `json:"id"`
	ProjectID      string    `json:"project_id"`
	IncomeDate     time.Time `json:"income_date"`
	IncomeCategory string    `json:"income_category"`
	Vendor         string    `json:"vendor"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Tax            float64   `json:"tax"`
	Status         string    `json:"status"`
	Currency       string    `json:"currency"`
	PaymentMethod  string    `json:"payment_method"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}
