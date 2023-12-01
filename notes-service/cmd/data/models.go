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
		ProductNote:         ProductNote{},
		InvoiceNote:         InvoiceNote{},
		InvoiceItemNote:     InvoiceItemNote{},
	}
}

type Models struct {
	ProjectNote         ProjectNote
	SubProjectNote      SubProjectNote
	IncomeNote          IncomeNote
	ExpenseNote         ExpenseNote
	ExternalCompanyNote ExternalCompanyNote
	ProductNote         ProductNote
	InvoiceNote         InvoiceNote
	InvoiceItemNote     InvoiceItemNote
}

type ProjectNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	ProjectId   string    `json:"project_id"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SubProjectNote struct {
	ID           string    `json:"id"`
	AuthorId     string    `json:"author_id"`
	AuthorName   string    `json:"author_name"`
	AuthorEmail  string    `json:"author_email"`
	SubProjectId string    `json:"sub_project_id"`
	Title        string    `json:"title"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type IncomeNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	IncomeId    string    `json:"income_id"`
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
	ExpenseId   string    `json:"expense_id"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ExternalCompanyNote struct {
	ID                string    `json:"id"`
	AuthorId          string    `json:"author_id"`
	AuthorName        string    `json:"author_name"`
	AuthorEmail       string    `json:"author_email"`
	ExternalCompanyId string    `json:"external_company_id"`
	Title             string    `json:"title"`
	Note              string    `json:"note"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type ProductNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	ProductId   string    `json:"product_id"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InvoiceNote struct {
	ID          string    `json:"id"`
	AuthorId    string    `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	InvoiceId   string    `json:"invoice_id"`
	Title       string    `json:"title"`
	Note        string    `json:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InvoiceItemNote struct {
	ID            string    `json:"id"`
	AuthorId      string    `json:"author_id"`
	AuthorName    string    `json:"author_name"`
	AuthorEmail   string    `json:"author_email"`
	InvoiceItemId string    `json:"invoice_item_id"`
	Title         string    `json:"title"`
	Note          string    `json:"note"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
