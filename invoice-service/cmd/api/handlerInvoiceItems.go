package main

import (
	"errors"
	"fmt"
	"invoice-service/cmd/data"
	"log"
	"net/http"
)

func (app *Config) CreateInvoiceItem(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_write")
	if err != nil {
		log.Println("authenticated - CreateInvoiceItem", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - CreateInvoiceItem")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload data.InvoiceItem

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - CreateInvoiceItem", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	newInvoiceItem := data.InvoiceItem{
		ProductId:          requestPayload.ProductId,
		Quantity:           requestPayload.Quantity,
		OriginalPrice:      requestPayload.OriginalPrice,
		ActualPrice:        requestPayload.ActualPrice,
		DiscountPercentage: requestPayload.DiscountPercentage,
		DiscountAmount:     requestPayload.DiscountAmount,
		TaxPercentage:      requestPayload.TaxPercentage,
		OriginalTax:        requestPayload.OriginalTax,
		ActualTax:          requestPayload.ActualTax,
	}

	response, err := data.InsertInvoiceItem(newInvoiceItem, userId)
	if err != nil {
		log.Println("postgres - CreateInvoiceItem", err)
		app.errorJSON(w, errors.New("could not create invoice item: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created invoice item %s", requestPayload.ProductId),
		Data:    response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateInvoiceItem(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_write")
	if err != nil {
		log.Println("authenticated - UpdateInvoiceItem", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - UpdateInvoiceItem")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var requestPayload data.InvoiceItem

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - UpdateInvoiceItem", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	invoiceItem := data.InvoiceItem{
		ID:                 requestPayload.ID,
		ProductId:          requestPayload.ProductId,
		Quantity:           requestPayload.Quantity,
		OriginalPrice:      requestPayload.OriginalPrice,
		ActualPrice:        requestPayload.ActualPrice,
		DiscountPercentage: requestPayload.DiscountPercentage,
		DiscountAmount:     requestPayload.DiscountAmount,
		TaxPercentage:      requestPayload.TaxPercentage,
		OriginalTax:        requestPayload.OriginalTax,
		ActualTax:          requestPayload.ActualTax,
	}

	err = data.UpdateInvoiceItem(invoiceItem, userId)
	if err != nil {
		log.Println("postgres - UpdateInvoiceItem", err)
		app.errorJSON(w, errors.New("could not update invoice item: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("updated invoice item with Id: %s", fmt.Sprint(invoiceItem.ProductId)),
		Data:    invoiceItem,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoiceItems(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_read")
	if err != nil {
		log.Println("authenticated - GetAllnvoiceItems", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllnvoiceItems")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	invoiceItems, err := data.GetAllInvoiceItems()
	if err != nil {
		log.Println("postgres - GetAllInvoiceItems", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var invoiceItemsSlice []data.InvoiceItem
	for _, invoiceItemPtr := range invoiceItems {
		invoiceItem := *invoiceItemPtr

		returnedSlice := data.InvoiceItem{
			ID:                 invoiceItem.ID,
			ProductId:          invoiceItem.ProductId,
			Quantity:           invoiceItem.Quantity,
			OriginalPrice:      invoiceItem.OriginalPrice,
			ActualPrice:        invoiceItem.ActualPrice,
			DiscountPercentage: invoiceItem.DiscountPercentage,
			DiscountAmount:     invoiceItem.DiscountAmount,
			TaxPercentage:      invoiceItem.TaxPercentage,
			OriginalTax:        invoiceItem.OriginalTax,
			ActualTax:          invoiceItem.ActualTax,
			CreatedBy:          invoiceItem.CreatedBy,
			CreatedAt:          invoiceItem.CreatedAt,
			UpdatedBy:          invoiceItem.UpdatedBy,
			UpdatedAt:          invoiceItem.UpdatedAt,
		}

		invoiceItemsSlice = append(invoiceItemsSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "fetched all invoice items",
		Data:    invoiceItemsSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoiceItemsByIds(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDSpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_read")
	if err != nil {
		log.Println("authenticated - GetAllInvoiceItemsByIds", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllInvoiceItemsByIds")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllInvoiceItemsByIds", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	invoiceItems, err := data.GetAllInvoiceItemsByIds(app.convertToPostgresArray(requestPayload.IDs))
	if err != nil {
		log.Println("postgres - GetAllInvoiceItemsByIds", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var invoiceItemsSlice []data.InvoiceItem
	for _, invoiceItemsPtr := range invoiceItems {
		invoiceItem := *invoiceItemsPtr

		returnedSlice := data.InvoiceItem{
			ID:                 invoiceItem.ID,
			ProductId:          invoiceItem.ProductId,
			Quantity:           invoiceItem.Quantity,
			OriginalPrice:      invoiceItem.OriginalPrice,
			ActualPrice:        invoiceItem.ActualPrice,
			DiscountPercentage: invoiceItem.DiscountPercentage,
			DiscountAmount:     invoiceItem.DiscountAmount,
			TaxPercentage:      invoiceItem.TaxPercentage,
			OriginalTax:        invoiceItem.OriginalTax,
			ActualTax:          invoiceItem.ActualTax,
			CreatedBy:          invoiceItem.CreatedBy,
			CreatedAt:          invoiceItem.CreatedAt,
			UpdatedBy:          invoiceItem.UpdatedBy,
			UpdatedAt:          invoiceItem.UpdatedAt,
		}

		invoiceItemsSlice = append(invoiceItemsSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all invoice items by product id",
		Data:    invoiceItemsSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllInvoiceItemsByProductId(w http.ResponseWriter, r *http.Request) {

	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_read")
	if err != nil {
		log.Println("authenticated - GetAllInvoiceItemsByProductId", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetAllInvoiceItemsByProductId")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetAllInvoiceItemsByProductId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	invoiceItems, err := data.GetAllInvoiceItemsByProductId(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetAllInvoiceItemsByProductId", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var invoiceItemsSlice []data.InvoiceItem
	for _, invoiceItemsPtr := range invoiceItems {
		invoiceItem := *invoiceItemsPtr

		returnedSlice := data.InvoiceItem{
			ID:                 invoiceItem.ID,
			ProductId:          invoiceItem.ProductId,
			Quantity:           invoiceItem.Quantity,
			OriginalPrice:      invoiceItem.OriginalPrice,
			ActualPrice:        invoiceItem.ActualPrice,
			DiscountPercentage: invoiceItem.DiscountPercentage,
			DiscountAmount:     invoiceItem.DiscountAmount,
			TaxPercentage:      invoiceItem.TaxPercentage,
			OriginalTax:        invoiceItem.OriginalTax,
			ActualTax:          invoiceItem.ActualTax,
			CreatedBy:          invoiceItem.CreatedBy,
			CreatedAt:          invoiceItem.CreatedAt,
			UpdatedBy:          invoiceItem.UpdatedBy,
			UpdatedAt:          invoiceItem.UpdatedAt,
		}

		invoiceItemsSlice = append(invoiceItemsSlice, returnedSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Fetched all invoice items by product id",
		Data:    invoiceItemsSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetInvoiceItemById(w http.ResponseWriter, r *http.Request) {
	var requestPayload IDpayload

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "invoice_read")
	if err != nil {
		log.Println("authenticated - GetInvoiceItemById", err)
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		log.Println("!authenticated - GetInvoiceItemById")
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println("readJSON - GetInvoiceItemById", err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	invoiceItem, err := data.GetInvoiceItemById(requestPayload.ID)
	if err != nil {
		log.Println("postgres - GetInvoiceItemById", err)
		app.errorJSON(w, errors.New("failed to get invoice by id"), http.StatusBadRequest)
		return
	}

	returnedInvoiceItem := data.InvoiceItem{
		ID:                 invoiceItem.ID,
		ProductId:          invoiceItem.ProductId,
		Quantity:           invoiceItem.Quantity,
		OriginalPrice:      invoiceItem.OriginalPrice,
		ActualPrice:        invoiceItem.ActualPrice,
		DiscountPercentage: invoiceItem.DiscountPercentage,
		DiscountAmount:     invoiceItem.DiscountAmount,
		TaxPercentage:      invoiceItem.TaxPercentage,
		OriginalTax:        invoiceItem.OriginalTax,
		ActualTax:          invoiceItem.ActualTax,
		CreatedBy:          invoiceItem.CreatedBy,
		CreatedAt:          invoiceItem.CreatedAt,
		UpdatedBy:          invoiceItem.UpdatedBy,
		UpdatedAt:          invoiceItem.UpdatedAt,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched invoice by id successfull: %s", invoiceItem.ID),
		Data:    returnedInvoiceItem,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
