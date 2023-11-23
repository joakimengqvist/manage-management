package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"invoice-service/cmd/data"
	"net/http"
	"time"
)

type NewIncome struct {
	ProjectID        string    `json:"project_id"`
	InvoiceID        string    `json:"invoice_id,omitempty"`
	IncomeDate       time.Time `json:"income_date"`
	IncomeCategory   string    `json:"income_category"`
	StatisticsIncome bool      `json:"statistics_income"`
	Vendor           string    `json:"vendor"`
	Description      string    `json:"description"`
	Amount           float64   `json:"amount"`
	Tax              float64   `json:"tax"`
	Status           string    `json:"status"`
	Currency         string    `json:"currency"`
}

func (app *Config) CreateInvoice(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload data.Invoice

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Println("readjson error", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newInvoice := data.InvoicePostgres{
		CompanyId:          requestPayload.CompanyId,
		ProjectId:          requestPayload.ProjectId,
		SubProjectId:       requestPayload.SubProjectId,
		InvoiceDisplayName: requestPayload.InvoiceDisplayName,
		InvoiceDescription: requestPayload.InvoiceDescription,
		StatisticsInvoice:  requestPayload.StatisticsInvoice,
		InvoiceItems:       app.convertToPostgresArray(requestPayload.InvoiceItems),
		OriginalPrice:      requestPayload.OriginalPrice,
		ActualPrice:        requestPayload.ActualPrice,
		DiscountPercentage: requestPayload.DiscountPercentage,
		DiscountAmount:     requestPayload.DiscountAmount,
		OriginalTax:        requestPayload.OriginalTax,
		ActualTax:          requestPayload.ActualTax,
		InvoiceDate:        requestPayload.InvoiceDate,
		DueDate:            requestPayload.DueDate,
		Paid:               requestPayload.Paid,
		Status:             requestPayload.Status,
		PaymentDate:        requestPayload.PaymentDate,
	}

	fmt.Println("newInvoice", newInvoice)

	response, err := data.InsertInvoice(newInvoice, userId)
	if err != nil {
		fmt.Println("response", err)
		app.errorJSON(w, errors.New("could not create invoice: "+err.Error()), http.StatusBadRequest)
		return
	}

	income := NewIncome{
		ProjectID:        requestPayload.ProjectId,
		InvoiceID:        response,
		IncomeDate:       requestPayload.InvoiceDate,
		IncomeCategory:   "invoice",
		StatisticsIncome: requestPayload.StatisticsInvoice,
		Vendor:           requestPayload.CompanyId,
		Description:      requestPayload.InvoiceDescription,
		Amount:           requestPayload.ActualPrice,
		Tax:              requestPayload.ActualTax,
		Status:           requestPayload.Status,
		Currency:         "SEK",
	}

	jsonData, err := json.Marshal(income)
	if err != nil {
		app.errorJSON(w, err)
	}

	request, err := http.NewRequest("POST", "http://economics-service/economics/create-income", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("X-User-Id", userId)

	client := &http.Client{}

	incomeResponse, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer incomeResponse.Body.Close()

	if incomeResponse.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("status unauthorized - create income"))
		return
	} else if incomeResponse.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling economics service - create income"))
		return
	}

	var jsonFromService jsonResponseCreateIncome

	err = json.NewDecoder(incomeResponse.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = data.UpdateIncomeId(jsonFromService.Data, response)

	if err != nil {
		fmt.Println("response", err)
		app.errorJSON(w, errors.New("could not create invoice: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created invoice %s", requestPayload.InvoiceDisplayName),
		Data:    response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateInvoice(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_write")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload data.Invoice

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	invoice := data.InvoicePostgres{
		ID:                 requestPayload.ID,
		CompanyId:          requestPayload.CompanyId,
		ProjectId:          requestPayload.ProjectId,
		SubProjectId:       requestPayload.SubProjectId,
		InvoiceDisplayName: requestPayload.InvoiceDisplayName,
		InvoiceDescription: requestPayload.InvoiceDescription,
		StatisticsInvoice:  requestPayload.StatisticsInvoice,
		InvoiceItems:       app.convertToPostgresArray(requestPayload.InvoiceItems),
		OriginalPrice:      requestPayload.OriginalPrice,
		ActualPrice:        requestPayload.ActualPrice,
		DiscountPercentage: requestPayload.DiscountPercentage,
		DiscountAmount:     requestPayload.DiscountAmount,
		OriginalTax:        requestPayload.OriginalTax,
		ActualTax:          requestPayload.ActualTax,
		InvoiceDate:        requestPayload.InvoiceDate,
		DueDate:            requestPayload.DueDate,
		Paid:               requestPayload.Paid,
		Status:             requestPayload.Status,
		PaymentDate:        requestPayload.PaymentDate,
	}

	err = data.UpdateInvoice(invoice, userId)
	if err != nil {
		app.errorJSON(w, errors.New("could not update invoice: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated invoice with Id: %s", fmt.Sprint(invoice.ID)),
		Data:    invoice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoices(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_read")
	if err != nil {
		fmt.Println("GetAllInvoices - authenticated error", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		fmt.Println("GetAllInvoices - not authenticated")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	invoices, err := data.GetAllInvoices()
	if err != nil {
		fmt.Println("GetAllProjectInvoices - invoices error", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var invoicesSlice []data.Invoice
	for _, invoicesPtr := range invoices {
		invoice := *invoicesPtr

		returnedSlice := data.Invoice{
			ID:                 invoice.ID,
			CompanyId:          invoice.CompanyId,
			ProjectId:          invoice.ProjectId,
			SubProjectId:       invoice.SubProjectId,
			IncomeId:           invoice.IncomeId,
			InvoiceDisplayName: invoice.InvoiceDisplayName,
			InvoiceDescription: invoice.InvoiceDescription,
			StatisticsInvoice:  invoice.StatisticsInvoice,
			InvoiceItems:       app.parsePostgresArray(invoice.InvoiceItems),
			OriginalPrice:      invoice.OriginalPrice,
			ActualPrice:        invoice.ActualPrice,
			DiscountPercentage: invoice.DiscountPercentage,
			DiscountAmount:     invoice.DiscountAmount,
			OriginalTax:        invoice.OriginalTax,
			ActualTax:          invoice.ActualTax,
			InvoiceDate:        invoice.InvoiceDate,
			DueDate:            invoice.DueDate,
			Paid:               invoice.Paid,
			Status:             invoice.Status,
			PaymentDate:        invoice.PaymentDate,
			CreatedBy:          invoice.CreatedBy,
			CreatedAt:          invoice.CreatedAt,
			UpdatedBy:          invoice.UpdatedBy,
			UpdatedAt:          invoice.UpdatedAt,
		}

		invoicesSlice = append(invoicesSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "fetched all exepnses",
		Data:    invoicesSlice,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get all privileges [/auth/get-all-privileges]", Name: "[authentication-service] - Successfuly fetched all privileges"})
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoicesByProjectId(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	invoices, err := data.GetAllInvoicesByProjectId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var invoicesSlice []data.Invoice
	for _, invoicesPtr := range invoices {
		invoice := *invoicesPtr

		returnedSlice := data.Invoice{
			ID:                 invoice.ID,
			CompanyId:          invoice.CompanyId,
			ProjectId:          invoice.ProjectId,
			SubProjectId:       invoice.SubProjectId,
			IncomeId:           invoice.IncomeId,
			InvoiceDisplayName: invoice.InvoiceDisplayName,
			InvoiceDescription: invoice.InvoiceDescription,
			StatisticsInvoice:  invoice.StatisticsInvoice,
			InvoiceItems:       app.parsePostgresArray(invoice.InvoiceItems),
			OriginalPrice:      invoice.OriginalPrice,
			ActualPrice:        invoice.ActualPrice,
			DiscountPercentage: invoice.DiscountPercentage,
			DiscountAmount:     invoice.DiscountAmount,
			OriginalTax:        invoice.OriginalTax,
			ActualTax:          invoice.ActualTax,
			InvoiceDate:        invoice.InvoiceDate,
			DueDate:            invoice.DueDate,
			Paid:               invoice.Paid,
			Status:             invoice.Status,
			PaymentDate:        invoice.PaymentDate,
			CreatedBy:          invoice.CreatedBy,
			CreatedAt:          invoice.CreatedAt,
			UpdatedBy:          invoice.UpdatedBy,
			UpdatedAt:          invoice.UpdatedAt,
		}

		invoicesSlice = append(invoicesSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all invoices by project id",
		Data:    invoicesSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoicesBySubProjectId(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	invoices, err := data.GetAllInvoicesBySubProjectId(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var invoicesSlice []data.Invoice
	for _, invoicesPtr := range invoices {
		invoice := *invoicesPtr

		returnedSlice := data.Invoice{
			ID:                 invoice.ID,
			CompanyId:          invoice.CompanyId,
			ProjectId:          invoice.ProjectId,
			SubProjectId:       invoice.SubProjectId,
			IncomeId:           invoice.IncomeId,
			InvoiceDisplayName: invoice.InvoiceDisplayName,
			InvoiceDescription: invoice.InvoiceDescription,
			StatisticsInvoice:  invoice.StatisticsInvoice,
			InvoiceItems:       app.parsePostgresArray(invoice.InvoiceItems),
			OriginalPrice:      invoice.OriginalPrice,
			ActualPrice:        invoice.ActualPrice,
			DiscountPercentage: invoice.DiscountPercentage,
			DiscountAmount:     invoice.DiscountAmount,
			OriginalTax:        invoice.OriginalTax,
			ActualTax:          invoice.ActualTax,
			InvoiceDate:        invoice.InvoiceDate,
			DueDate:            invoice.DueDate,
			Paid:               invoice.Paid,
			Status:             invoice.Status,
			PaymentDate:        invoice.PaymentDate,
			CreatedBy:          invoice.CreatedBy,
			CreatedAt:          invoice.CreatedAt,
			UpdatedBy:          invoice.UpdatedBy,
			UpdatedAt:          invoice.UpdatedAt,
		}

		invoicesSlice = append(invoicesSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all invoices by sub project id",
		Data:    invoicesSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetInvoiceById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	invoice, err := data.GetInvoiceById(requestPayload.ID)
	if err != nil {
		fmt.Println("error getting", err)
		app.errorJSON(w, errors.New("failed to get invoice by id"), http.StatusBadRequest)
		return
	}

	returnedInvoice := data.Invoice{
		ID:                 invoice.ID,
		CompanyId:          invoice.CompanyId,
		ProjectId:          invoice.ProjectId,
		SubProjectId:       invoice.SubProjectId,
		IncomeId:           invoice.IncomeId,
		InvoiceDisplayName: invoice.InvoiceDisplayName,
		InvoiceDescription: invoice.InvoiceDescription,
		StatisticsInvoice:  invoice.StatisticsInvoice,
		InvoiceItems:       app.parsePostgresArray(invoice.InvoiceItems),
		OriginalPrice:      invoice.OriginalPrice,
		ActualPrice:        invoice.ActualPrice,
		DiscountPercentage: invoice.DiscountPercentage,
		DiscountAmount:     invoice.DiscountAmount,
		OriginalTax:        invoice.OriginalTax,
		ActualTax:          invoice.ActualTax,
		InvoiceDate:        invoice.InvoiceDate,
		DueDate:            invoice.DueDate,
		Paid:               invoice.Paid,
		Status:             invoice.Status,
		PaymentDate:        invoice.PaymentDate,
		CreatedBy:          invoice.CreatedBy,
		CreatedAt:          invoice.CreatedAt,
		UpdatedBy:          invoice.UpdatedBy,
		UpdatedAt:          invoice.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched invoice by id successfull: %s", invoice.ID),
		Data:    returnedInvoice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
