package main

import (
	"errors"
	"fmt"
	"net/http"
	"product-service/cmd/data"
)

type ProductId struct {
	ID string `json:"id"`
}

func (app *Config) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.Product

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "product_write")
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

	newCompany := data.NewProduct{
		Name:          requestPayload.Name,
		Description:   requestPayload.Description,
		Category:      requestPayload.Category,
		Price:         requestPayload.Price,
		TaxPercentage: requestPayload.TaxPercentage,
		CreatedBy:     userId,
	}

	response, err := data.InsertProduct(newCompany)
	if err != nil {
		app.errorJSON(w, errors.New("could not create product: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created new product %s", requestPayload.Name),
		Data:    response,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.Product

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "product_write")
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

	updatedCompany := data.Product{
		ID:            requestPayload.ID,
		Name:          requestPayload.Name,
		Description:   requestPayload.Description,
		Category:      requestPayload.Category,
		Price:         requestPayload.Price,
		TaxPercentage: requestPayload.TaxPercentage,
		UpdatedBy:     userId,
	}

	err = data.UpdateProduct(updatedCompany)
	if err != nil {
		app.errorJSON(w, errors.New("could not update product: "+err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Updated product %s", requestPayload.Name),
		Data:    nil,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "product_read")
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	if !authenticated {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	companies, err := data.GetAllProducts()
	if err != nil {
		app.errorJSON(w, errors.New("could not fetch products"), http.StatusUnauthorized)
		return
	}

	var productsSlice []data.Product
	for _, productPtr := range companies {
		product := *productPtr

		productSlice := data.Product{
			ID:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			Category:      product.Category,
			Price:         product.Price,
			TaxPercentage: product.TaxPercentage,
			CreatedAt:     product.CreatedAt,
			CreatedBy:     product.CreatedBy,
			UpdatedAt:     product.UpdatedAt,
			UpdatedBy:     product.UpdatedBy,
		}

		productsSlice = append(productsSlice, productSlice)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Successfully fetched all products",
		Data:    productsSlice,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetProductById(w http.ResponseWriter, r *http.Request) {
	var requestPayload ProductId

	userId := r.Header.Get("X-User-Id")
	authenticated, err := app.CheckPrivilege(w, userId, "product_read")
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

	product, err := data.GetProductById(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("failed to get user by id"), http.StatusBadRequest)
		return
	}

	returnedCompany := data.Product{
		ID:            product.ID,
		Name:          product.Name,
		Description:   product.Description,
		Category:      product.Category,
		Price:         product.Price,
		TaxPercentage: product.TaxPercentage,
		CreatedAt:     product.CreatedAt,
		CreatedBy:     product.CreatedBy,
		UpdatedAt:     product.UpdatedAt,
		UpdatedBy:     product.UpdatedBy,
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Fetched product: %s ", product.Name),
		Data:    returnedCompany,
	}

	app.logItemViaRPC(w, payload, RPCLogData{Action: "Get product by id", Name: "[product-service] - Successfuly fetched product"})
	app.writeJSON(w, http.StatusAccepted, payload)
}
