package data

import (
	"context"
	"database/sql"
	"fmt"
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
	CreatedAt                 time.Time `json:"created_at,omitempty"`
	CreatedBy                 string    `json:"created_by,omitempty"`
	UpdatedAt                 time.Time `json:"updated_at,omitempty"`
	UpdatedBy                 string    `json:"updated_by,omitempty"`
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
	CreatedBy                 string    `json:"created_by"`
	UpdatedAt                 time.Time `json:"updated_at"`
	UpdatedBy                 string    `json:"updated_by"`
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
	CreatedBy                 string    `json:"created_by"`
	UpdatedAt                 time.Time `json:"updated_at"`
	UpdatedBy                 string    `json:"updated_by"`
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

	stmt := `insert into external_companies (company_name, company_registration_number, contact_person, contact_email, contact_phone, address, city, state_province, country, postal_code, payment_terms, billing_currency, bank_account_info, tax_identification_number, created_at, created_by, updated_at, updated_by, status, assigned_projects, invoice_pending, invoice_history, contractual_agreements)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
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
		company.CreatedBy,
		time.Now(),
		company.CreatedBy,
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

func UpdateExternalCompany(p ExternalCompanyPostgres) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update external_companies set
		company_name = $1,
		company_registration_number = $2,
		contact_person = $3,
		contact_email = $4,
		contact_phone = $5,
		address = $6,
		city = $7,
		state_province = $8,
		country = $9,
		postal_code = $10,
		payment_terms = $11,
		billing_currency = $12,
		bank_account_info = $13,
		tax_identification_number = $14,
		updated_at = $15,
		updated_by = $16,
		status = $17,
		assigned_projects = $18,
		invoice_pending = $19,
		invoice_history = $20,
		contractual_agreements = $21
		where id = $22
	`

	_, err := db.ExecContext(ctx, stmt,
		p.CompanyName,
		p.CompanyRegistrationNumber,
		p.ContactPerson,
		p.ContactEmail,
		p.ContactPhone,
		p.Address,
		p.City,
		p.StateProvince,
		p.Country,
		p.PostalCode,
		p.PaymentTerms,
		p.BillingCurrency,
		p.BankAccountInfo,
		p.TaxIdentificationNumber,
		time.Now(),
		p.UpdatedBy,
		p.Status,
		p.AssignedProjects,
		p.InvoicePending,
		p.InvoiceHistory,
		p.ContractualAgreements,
		p.ID,
	)

	if err != nil {
		fmt.Println("Error updating sub project", err)
		return err
	}

	return nil
}

func GetAllExternalCompanies() ([]*ExternalCompanyPostgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, company_name, company_registration_number, contact_person, contact_email, contact_phone, address, city, state_province, country, postal_code, payment_terms, billing_currency, bank_account_info, tax_identification_number, created_at, created_by, updated_at, updated_by, status, assigned_projects, invoice_pending, invoice_history, contractual_agreements
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
			&company.CreatedBy,
			&company.UpdatedAt,
			&company.UpdatedBy,
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

	query := `select id, company_name, company_registration_number, contact_person, contact_email, contact_phone, address, city, state_province, country, postal_code, payment_terms, billing_currency, bank_account_info, tax_identification_number, created_at, created_by, updated_at, updated_by, status, assigned_projects, invoice_pending, invoice_history, contractual_agreements from external_companies where id = $1`

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
		&company.CreatedBy,
		&company.UpdatedAt,
		&company.UpdatedBy,
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
