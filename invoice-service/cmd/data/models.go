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
		Invoice:     Invoice{},
		InvoiceItem: InvoiceItem{},
	}
}

type Models struct {
	Invoice     Invoice
	InvoiceItem InvoiceItem
}

type Invoice struct {
	ID                 string    `json:"id,omitempty"`
	CompanyId          string    `json:"company_id"`
	ProjectId          string    `json:"project_id"`
	SubProjectId       string    `json:"sub_project_id"`
	IncomeId           string    `json:"income_id,omitempty"`
	InvoiceDisplayName string    `json:"invoice_display_name"`
	StatisticsInvoice  bool      `json:"statistics_invoice"`
	InvoiceDescription string    `json:"invoice_description"`
	InvoiceItems       []string  `json:"invoice_items"`
	OriginalPrice      float64   `json:"original_price"`
	ActualPrice        float64   `json:"actual_price"`
	DiscountPercentage float64   `json:"discount_percentage"`
	DiscountAmount     float64   `json:"discount_amount"`
	OriginalTax        float64   `json:"original_tax"`
	ActualTax          float64   `json:"actual_tax"`
	InvoiceDate        time.Time `json:"invoice_date"`
	DueDate            time.Time `json:"due_date"`
	Paid               bool      `json:"paid"`
	Status             string    `json:"status"`
	PaymentDate        time.Time `json:"payment_date"`
	CreatedBy          string    `json:"created_by,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedBy          string    `json:"updated_by,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

type InvoicePostgres struct {
	ID                 string    `json:"id,omitempty"`
	CompanyId          string    `json:"company_id"`
	ProjectId          string    `json:"project_id"`
	SubProjectId       string    `json:"sub_project_id"`
	IncomeId           string    `json:"income_id,omitempty"`
	InvoiceDisplayName string    `json:"invoice_display_name"`
	InvoiceDescription string    `json:"invoice_description"`
	StatisticsInvoice  bool      `json:"statistics_invoice"`
	InvoiceItems       string    `json:"invoice_items"`
	OriginalPrice      float64   `json:"original_price"`
	ActualPrice        float64   `json:"actual_price"`
	DiscountPercentage float64   `json:"discount_percentage"`
	DiscountAmount     float64   `json:"discount_amount"`
	OriginalTax        float64   `json:"original_tax"`
	ActualTax          float64   `json:"actual_tax"`
	InvoiceDate        time.Time `json:"invoice_date"`
	DueDate            time.Time `json:"due_date"`
	Paid               bool      `json:"paid"`
	Status             string    `json:"status"`
	PaymentDate        time.Time `json:"payment_date"`
	CreatedBy          string    `json:"created_by,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedBy          string    `json:"updated_by,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

type InvoiceItem struct {
	ID                 string    `json:"id,omitempty"`
	ProductId          string    `json:"product_id"`
	Quantity           int       `json:"quantity"`
	DiscountPercentage float64   `json:"discount_percentage"`
	DiscountAmount     float64   `json:"discount_amount"`
	OriginalPrice      float64   `json:"original_price"`
	ActualPrice        float64   `json:"actual_price"`
	TaxPercentage      float64   `json:"tax_percentage"`
	OriginalTax        float64   `json:"original_tax"`
	ActualTax          float64   `json:"actual_tax"`
	CreatedBy          string    `json:"created_by,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedBy          string    `json:"updated_by,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}
