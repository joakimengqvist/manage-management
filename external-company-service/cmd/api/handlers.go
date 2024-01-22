package main

import (
	"errors"
	"external-company-service/cmd/data"
	"fmt"
	"log"
	"net/http"
)

type ExternalCompanyId struct {
	ID string `json:"id"`
}

func (app *Config) CreateExternalCompany(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.ExternalCompany

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "external_company_write")
	if err != nil {
		log.Println("authenticated - CreateExternalCompany", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - CreateExternalCompany")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateExternalCompany", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newCompany := data.NewExternalCompany{
		CompanyName:               requestPayload.CompanyName,
		CompanyRegistrationNumber: requestPayload.CompanyRegistrationNumber,
		ContactPerson:             requestPayload.ContactPerson,
		ContactEmail:              requestPayload.ContactEmail,
		ContactPhone:              requestPayload.ContactPhone,
		Address:                   requestPayload.Address,
		City:                      requestPayload.City,
		StateProvince:             requestPayload.StateProvince,
		Country:                   requestPayload.Country,
		PostalCode:                requestPayload.PostalCode,
		PaymentTerms:              requestPayload.PaymentTerms,
		BillingCurrency:           requestPayload.BillingCurrency,
		BankAccountInfo:           requestPayload.BankAccountInfo,
		TaxIdentificationNumber:   requestPayload.TaxIdentificationNumber,
		CreatedBy:                 userId,
		UpdatedBy:                 userId,
		Status:                    requestPayload.Status,
		AssignedProjects:          app.convertToPostgresArray(requestPayload.AssignedProjects),
		Invoices:                  app.convertToPostgresArray(requestPayload.Invoices),
		ContractualAgreements:     app.convertToPostgresArray(requestPayload.ContractualAgreements),
	}

	response, err := data.InsertExternalCompany(newCompany)
	if err != nil {
		log.Println("postgres - CreateExternalCompany", err)
		app.errorJSON(w, errors.New("could not create external company: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created new external company %s", requestPayload.CompanyName),
		Data:    response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateExternalCompany(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.ExternalCompany

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "external_company_write")
	if err != nil {
		log.Println("authenticated - UpdateExternalCompany", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - UpdateExternalCompany")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateExternalCompany", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	updatedCompany := data.ExternalCompanyPostgres{
		ID:                        requestPayload.ID,
		CompanyName:               requestPayload.CompanyName,
		CompanyRegistrationNumber: requestPayload.CompanyRegistrationNumber,
		ContactPerson:             requestPayload.ContactPerson,
		ContactEmail:              requestPayload.ContactEmail,
		ContactPhone:              requestPayload.ContactPhone,
		Address:                   requestPayload.Address,
		City:                      requestPayload.City,
		StateProvince:             requestPayload.StateProvince,
		Country:                   requestPayload.Country,
		PostalCode:                requestPayload.PostalCode,
		PaymentTerms:              requestPayload.PaymentTerms,
		BillingCurrency:           requestPayload.BillingCurrency,
		BankAccountInfo:           requestPayload.BankAccountInfo,
		TaxIdentificationNumber:   requestPayload.TaxIdentificationNumber,
		UpdatedBy:                 userId,
		Status:                    requestPayload.Status,
		AssignedProjects:          app.convertToPostgresArray(requestPayload.AssignedProjects),
		Invoices:                  app.convertToPostgresArray(requestPayload.Invoices),
		ContractualAgreements:     app.convertToPostgresArray(requestPayload.ContractualAgreements),
	}

	err = data.UpdateExternalCompany(updatedCompany)
	if err != nil {
		log.Println("postgres - UpdateExternalCompany", err)
		app.errorJSON(w, errors.New("could not update external company: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Updated external company %s", requestPayload.CompanyName),
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) AddInvoiceToCompany(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.InvoicePayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "external_company_write")
	if err != nil {
		log.Println("authenticated - AddInvoiceToCompany", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - AddInvoiceToCompany")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - AddInvoiceToCompany", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = data.AppendInvoiceToCompany(requestPayload.InvoiceId, requestPayload.CompanyId)
	if err != nil {
		log.Println("postgres - AddInvoiceToCompany", err)
		app.errorJSON(w, errors.New("could not update external company: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Updated external company",
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) RemoveInvoiceToCompany(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.InvoicePayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "external_company_write")
	if err != nil {
		log.Println("authenticated - RemoveInvoiceToCompany", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - RemoveInvoiceToCompany")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - RemoveInvoiceToCompany", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = data.RemoveInvoiceFromCompany(requestPayload.InvoiceId, requestPayload.CompanyId)
	if err != nil {
		log.Println("postgres - RemoveInvoiceToCompany", err)
		app.errorJSON(w, errors.New("could not update external company: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Updated external company",
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllExternalCompanies(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "external_company_read")
	if err != nil {
		log.Println("authenticated - GetAllExternalCompanies", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllExternalCompanies")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	companies, err := data.GetAllExternalCompanies()
	if err != nil {
		log.Println("postgres - GetAllExternalCompanies", err)
		app.errorJSON(w, errors.New("could not fetch external companies"), http.StatusUnauthorized)
		return
	}

	var companiesSlice []data.ExternalCompany
	for _, companyPtr := range companies {
		company := *companyPtr

		returnedSlice := data.ExternalCompany{
			ID:                        company.ID,
			CompanyName:               company.CompanyName,
			CompanyRegistrationNumber: company.CompanyRegistrationNumber,
			ContactPerson:             company.ContactPerson,
			ContactEmail:              company.ContactEmail,
			ContactPhone:              company.ContactPhone,
			Address:                   company.Address,
			City:                      company.City,
			StateProvince:             company.StateProvince,
			Country:                   company.Country,
			PostalCode:                company.PostalCode,
			PaymentTerms:              company.PaymentTerms,
			BillingCurrency:           company.BillingCurrency,
			BankAccountInfo:           company.BankAccountInfo,
			TaxIdentificationNumber:   company.TaxIdentificationNumber,
			CreatedAt:                 company.CreatedAt,
			CreatedBy:                 company.CreatedBy,
			UpdatedAt:                 company.UpdatedAt,
			UpdatedBy:                 company.UpdatedBy,
			Status:                    company.Status,
			AssignedProjects:          app.parsePostgresArray(company.AssignedProjects),
			Invoices:                  app.parsePostgresArray(company.Invoices),
			ContractualAgreements:     app.parsePostgresArray(company.ContractualAgreements),
		}

		companiesSlice = append(companiesSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Successfully fetched all external companies",
		Data:    companiesSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetExternalCompanyById(w http.ResponseWriter, r *http.Request) {
	var requestPayload ExternalCompanyId

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "external_company_read")
	if err != nil {
		log.Println("authenticated - GetExternalCompanyById", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetExternalCompanyById")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetExternalCompanyById", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	company, err := data.GetExternalCompanyById(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetExternalCompanyById", err)
		app.errorJSON(w, errors.New("failed to get user by id"), http.StatusBadRequest)
		return
	}

	returnedCompany := data.ExternalCompany{
		ID:                        company.ID,
		CompanyName:               company.CompanyName,
		CompanyRegistrationNumber: company.CompanyRegistrationNumber,
		ContactPerson:             company.ContactPerson,
		ContactEmail:              company.ContactEmail,
		ContactPhone:              company.ContactPhone,
		Address:                   company.Address,
		City:                      company.City,
		StateProvince:             company.StateProvince,
		Country:                   company.Country,
		PostalCode:                company.PostalCode,
		PaymentTerms:              company.PaymentTerms,
		BillingCurrency:           company.BillingCurrency,
		BankAccountInfo:           company.BankAccountInfo,
		TaxIdentificationNumber:   company.TaxIdentificationNumber,
		CreatedAt:                 company.CreatedAt,
		CreatedBy:                 company.CreatedBy,
		UpdatedAt:                 company.UpdatedAt,
		UpdatedBy:                 company.UpdatedBy,
		Status:                    company.Status,
		AssignedProjects:          app.parsePostgresArray(company.AssignedProjects),
		Invoices:                  app.parsePostgresArray(company.Invoices),
		ContractualAgreements:     app.parsePostgresArray(company.ContractualAgreements),
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched external company: %s ", company.CompanyName),
		Data:    returnedCompany,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
