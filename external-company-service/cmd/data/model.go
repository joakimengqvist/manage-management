package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		ExternalCompany:    ExternalCompany{},
		NewExternalCompany: NewExternalCompany{},
	}
}

type Models struct {
	ExternalCompany    ExternalCompany
	NewExternalCompany NewExternalCompany
}

type ExternalCompany struct {
	ID                        string    `json:"id"`
	CompanyName               string    `json:"company_name" sql:"not null"`
	CompanyRegistrationNumber string    `json:"company_registration_number"`
	ContactPerson             string    `json:"contact_person"`
	ContactEmail              string    `json:"contact_email"`
	ContactPhone              string    `json:"contact_phone"`
	Address                   string    `json:"address"`
	City                      string    `json:"city"`
	StateProvince             string    `json:"state_province"`
	Country                   string    `json:"country"`
	PostalCode                string    `json:"postal_code"`
	PaymentTerms              string    `json:"payment_terms"`
	BillingCurrency           string    `json:"billing_currency"`
	BankAccountInfo           string    `json:"bank_account_info"`
	TaxIdentificationNumber   string    `json:"tax_identification_number"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	Status                    string    `json:"status"`
	AssignedProjects          []string  `json:"assigned_projects"`
	InvoicePending            []string  `json:"invoice_pending"`
	InvoiceHistory            []string  `json:"invoice_history"`
	ContractualAgreements     []string  `json:"contractual_agreements"`
}

type ExternalCompanyPostgres struct {
	ID                        string    `json:"id"`
	CompanyName               string    `json:"company_name" sql:"not null"`
	CompanyRegistrationNumber string    `json:"company_registration_number"`
	ContactPerson             string    `json:"contact_person"`
	ContactEmail              string    `json:"contact_email"`
	ContactPhone              string    `json:"contact_phone"`
	Address                   string    `json:"address"`
	City                      string    `json:"city"`
	StateProvince             string    `json:"state_province"`
	Country                   string    `json:"country"`
	PostalCode                string    `json:"postal_code"`
	PaymentTerms              string    `json:"payment_terms"`
	BillingCurrency           string    `json:"billing_currency"`
	BankAccountInfo           string    `json:"bank_account_info"`
	TaxIdentificationNumber   string    `json:"tax_identification_number"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	Status                    string    `json:"status"`
	AssignedProjects          string    `json:"assigned_projects"`
	InvoicePending            string    `json:"invoice_pending"`
	InvoiceHistory            string    `json:"invoice_history"`
	ContractualAgreements     string    `json:"contractual_agreements"`
}

type NewExternalCompany struct {
	CompanyName               string    `json:"company_name" sql:"not null"`
	CompanyRegistrationNumber string    `json:"company_registration_number"`
	ContactPerson             string    `json:"contact_person"`
	ContactEmail              string    `json:"contact_email"`
	ContactPhone              string    `json:"contact_phone"`
	Address                   string    `json:"address"`
	City                      string    `json:"city"`
	StateProvince             string    `json:"state_province"`
	Country                   string    `json:"country"`
	PostalCode                string    `json:"postal_code"`
	PaymentTerms              string    `json:"payment_terms"`
	BillingCurrency           string    `json:"billing_currency"`
	BankAccountInfo           string    `json:"bank_account_info"`
	TaxIdentificationNumber   string    `json:"tax_identification_number"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	Status                    string    `json:"status"`
	AssignedProjects          string    `json:"assigned_projects"`
	InvoicePending            string    `json:"invoice_pending"`
	InvoiceHistory            string    `json:"invoice_history"`
	ContractualAgreements     string    `json:"contractual_agreements"`
}

func InsertExternalCompany(company NewExternalCompany) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string

	stmt := `insert into external_companies (company_name, company_registration_number, contact_person, contact_email, contact_phone, address, city, state_province, country, postal_code, payment_terms, billing_currency, bank_account_info, tax_identification_number, created_at, updated_at, status, assigned_projects, invoice_pending, invoice_history, contractual_agreements)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
		RETURNING id`

	err := db.QueryRowContext(ctx, stmt,
		company.CompanyName,
		company.CompanyRegistrationNumber,
		company.ContactPerson,
		company.ContactEmail,
		company.ContactPhone,
		company.Address,
		company.City,
		company.StateProvince,
		company.Country,
		company.PostalCode,
		company.PaymentTerms,
		company.BillingCurrency,
		company.BankAccountInfo,
		company.TaxIdentificationNumber,
		time.Now(),
		time.Now(),
		company.Status,
		company.AssignedProjects,
		company.InvoicePending,
		company.InvoiceHistory,
		company.ContractualAgreements,
	).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}

func GetAllExternalCompanies() ([]*ExternalCompanyPostgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, company_name, company_registration_number, contact_person, contact_email, contact_phone, address, city, state_province, country, postal_code, payment_terms, billing_currency, bank_account_info, tax_identification_number, created_at, updated_at, status, assigned_projects, invoice_pending, invoice_history, contractual_agreements
	from external_companies order by updated_at desc`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []*ExternalCompanyPostgres

	for rows.Next() {
		var company ExternalCompanyPostgres
		err := rows.Scan(
			&company.ID,
			&company.CompanyName,
			&company.CompanyRegistrationNumber,
			&company.ContactPerson,
			&company.ContactEmail,
			&company.ContactPhone,
			&company.Address,
			&company.City,
			&company.StateProvince,
			&company.Country,
			&company.PostalCode,
			&company.PaymentTerms,
			&company.BillingCurrency,
			&company.BankAccountInfo,
			&company.TaxIdentificationNumber,
			&company.CreatedAt,
			&company.UpdatedAt,
			&company.Status,
			&company.AssignedProjects,
			&company.InvoicePending,
			&company.InvoiceHistory,
			&company.ContractualAgreements,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		companies = append(companies, &company)
	}

	return companies, nil
}

func GetExternalCompanyById(id string) (*ExternalCompanyPostgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, company_name, company_registration_number, contact_person, contact_email, contact_phone, address, city, state_province, country, postal_code, payment_terms, billing_currency, bank_account_info, tax_identification_number, created_at, updated_at, status, assigned_projects, invoice_pending, invoice_history, contractual_agreements from external_companies where id = $1`

	var company ExternalCompanyPostgres
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&company.ID,
		&company.CompanyName,
		&company.CompanyRegistrationNumber,
		&company.ContactPerson,
		&company.ContactEmail,
		&company.ContactPhone,
		&company.Address,
		&company.City,
		&company.StateProvince,
		&company.Country,
		&company.PostalCode,
		&company.PaymentTerms,
		&company.BillingCurrency,
		&company.BankAccountInfo,
		&company.TaxIdentificationNumber,
		&company.CreatedAt,
		&company.UpdatedAt,
		&company.Status,
		&company.AssignedProjects,
		&company.InvoicePending,
		&company.InvoiceHistory,
		&company.ContractualAgreements,
	)

	if err != nil {
		return nil, err
	}

	return &company, nil
}
