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
		ProjectNote:         ProjectNote{},
		SubProjectNote:      SubProjectNote{},
		IncomeNote:          IncomeNote{},
		ExpenseNote:         ExpenseNote{},
		ExternalCompanyNote: ExternalCompanyNote{},
	}
}

type Models struct {
	ProjectNote         ProjectNote
	SubProjectNote      SubProjectNote
	IncomeNote          IncomeNote
	ExpenseNote         ExpenseNote
	ExternalCompanyNote ExternalCompanyNote
}

type ProjectNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	Project     string    `json:"project"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SubProjectNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	SubProject  string    `json:"sub_project"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type IncomeNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	Income      string    `json:"income"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExpenseNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	Expense     string    `json:"expense"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExternalCompanyNote struct {
	ID              string    `json:"id"`
	AuthorId        string    `json:"author_id"`
	AuthorName      string    `json:"author_name"`
	AuthorEmail     string    `json:"author_email"`
	ExternalCompany string    `json:"external_company"`
	Title           string    `json:"title"`
	Note            string    `json:"note"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
