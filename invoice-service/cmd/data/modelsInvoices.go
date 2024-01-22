package data

import (
	"context"
	"log"
	"time"
)

func InsertInvoice(invoice InvoicePostgres, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string

	stmt := `
	insert into invoices (
		company_id,
		project_id,
		sub_project_id,
		invoice_display_name,
		invoice_description,
		statistics_invoice,
		invoice_items,
		original_price,
		actual_price,
		discount_percentage,
		discount_amount,
		original_tax,
		actual_tax,
		invoice_date,
		due_date,
		payment_date,
		paid,
		status,
		created_by,
		created_at,
		updated_by,
		updated_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) returning id`

	err := db.QueryRowContext(ctx, stmt,
		invoice.CompanyId,
		invoice.ProjectId,
		invoice.SubProjectId,
		invoice.InvoiceDisplayName,
		invoice.InvoiceDescription,
		invoice.StatisticsInvoice,
		invoice.InvoiceItems,
		invoice.OriginalPrice,
		invoice.ActualPrice,
		invoice.DiscountPercentage,
		invoice.DiscountAmount,
		invoice.OriginalTax,
		invoice.ActualTax,
		invoice.InvoiceDate,
		invoice.DueDate,
		invoice.PaymentDate,
		invoice.Paid,
		invoice.Status,
		userId,
		time.Now(),
		userId,
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}

func GetInvoiceById(invoiceId string) (*InvoicePostgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, company_id, project_id, sub_project_id, income_id, invoice_display_name, invoice_description, statistics_invoice, invoice_items, original_price, actual_price, discount_percentage, discount_amount, original_tax, actual_tax, invoice_date, due_date, paid, status, payment_date, created_by, created_at, updated_by, updated_at from invoices where id = $1`

	row := db.QueryRowContext(ctx, query, invoiceId)

	var invoice InvoicePostgres
	err := row.Scan(
		&invoice.ID,
		&invoice.CompanyId,
		&invoice.ProjectId,
		&invoice.SubProjectId,
		&invoice.IncomeId,
		&invoice.InvoiceDisplayName,
		&invoice.InvoiceDescription,
		&invoice.StatisticsInvoice,
		&invoice.InvoiceItems,
		&invoice.OriginalPrice,
		&invoice.ActualPrice,
		&invoice.DiscountPercentage,
		&invoice.DiscountAmount,
		&invoice.OriginalTax,
		&invoice.ActualTax,
		&invoice.InvoiceDate,
		&invoice.DueDate,
		&invoice.Paid,
		&invoice.Status,
		&invoice.PaymentDate,
		&invoice.CreatedBy,
		&invoice.CreatedAt,
		&invoice.UpdatedBy,
		&invoice.UpdatedAt,
	)
	if err != nil {
		log.Println("Error scanning", err)
		return nil, err
	}

	return &invoice, nil
}

func GetAllInvoicesByProjectId(subProjectId string) ([]*InvoicePostgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select 
		id,
		company_id,
		project_id,
		sub_project_id,
		income_id,
		invoice_display_name,
		invoice_description,
		statistics_invoice,
		invoice_items,
		original_price,
		actual_price,
		discount_percentage,
		discount_amount,
		original_tax,
		actual_tax,
		invoice_date,
		due_date,
		payment_date,
		paid,
		status,
		created_by,
		created_at, 
		updated_by, 
		updated_at
		from invoices where project_id = $1`

	rows, err := db.QueryContext(ctx, query, subProjectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*InvoicePostgres

	for rows.Next() {
		var invoice InvoicePostgres
		err := rows.Scan(
			&invoice.ID,
			&invoice.CompanyId,
			&invoice.ProjectId,
			&invoice.SubProjectId,
			&invoice.IncomeId,
			&invoice.InvoiceDisplayName,
			&invoice.InvoiceDescription,
			&invoice.StatisticsInvoice,
			&invoice.InvoiceItems,
			&invoice.OriginalPrice,
			&invoice.ActualPrice,
			&invoice.DiscountPercentage,
			&invoice.DiscountAmount,
			&invoice.OriginalTax,
			&invoice.ActualTax,
			&invoice.InvoiceDate,
			&invoice.DueDate,
			&invoice.PaymentDate,
			&invoice.Paid,
			&invoice.Status,
			&invoice.CreatedBy,
			&invoice.CreatedAt,
			&invoice.UpdatedBy,
			&invoice.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

func GetAllInvoicesBySubProjectId(subProjectId string) ([]*InvoicePostgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, company_id, project_id, sub_project_id, income_id, invoice_display_name, invoice_description, statistics_invoice, invoice_items, original_price, actual_price, discount_percentage, discount_amount, original_tax, actual_tax, invoice_date, due_date, paid, status, payment_date, created_by, created_at, updated_by, updated_at from invoices where sub_project_id = $1`

	rows, err := db.QueryContext(ctx, query, subProjectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*InvoicePostgres

	for rows.Next() {
		var invoice InvoicePostgres
		err := rows.Scan(
			&invoice.ID,
			&invoice.CompanyId,
			&invoice.ProjectId,
			&invoice.SubProjectId,
			&invoice.IncomeId,
			&invoice.InvoiceDisplayName,
			&invoice.InvoiceDescription,
			&invoice.StatisticsInvoice,
			&invoice.InvoiceItems,
			&invoice.OriginalPrice,
			&invoice.ActualPrice,
			&invoice.DiscountPercentage,
			&invoice.DiscountAmount,
			&invoice.OriginalTax,
			&invoice.ActualTax,
			&invoice.InvoiceDate,
			&invoice.DueDate,
			&invoice.Paid,
			&invoice.Status,
			&invoice.PaymentDate,
			&invoice.CreatedBy,
			&invoice.CreatedAt,
			&invoice.UpdatedBy,
			&invoice.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

func GetAllInvoices() ([]*InvoicePostgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, company_id, project_id, sub_project_id, income_id, invoice_display_name, invoice_description, statistics_invoice, invoice_items, original_price, actual_price, discount_percentage, discount_amount, original_tax, actual_tax, invoice_date, due_date, paid, status, payment_date, created_by, created_at, updated_by, updated_at from invoices`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*InvoicePostgres

	for rows.Next() {
		var invoice InvoicePostgres
		err := rows.Scan(
			&invoice.ID,
			&invoice.CompanyId,
			&invoice.ProjectId,
			&invoice.SubProjectId,
			&invoice.IncomeId,
			&invoice.InvoiceDisplayName,
			&invoice.InvoiceDescription,
			&invoice.StatisticsInvoice,
			&invoice.InvoiceItems,
			&invoice.OriginalPrice,
			&invoice.ActualPrice,
			&invoice.DiscountPercentage,
			&invoice.DiscountAmount,
			&invoice.OriginalTax,
			&invoice.ActualTax,
			&invoice.InvoiceDate,
			&invoice.DueDate,
			&invoice.Paid,
			&invoice.Status,
			&invoice.PaymentDate,
			&invoice.CreatedBy,
			&invoice.CreatedAt,
			&invoice.UpdatedBy,
			&invoice.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

func UpdateInvoice(invoice InvoicePostgres, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update invoices set 
		company_id = $1,
		project_id = $2,
		sub_project_id = $3,
		invoice_display_name = $4,
		invoice_description = $5,
		statistics_invoice = $6,
		invoice_items = $7,
		original_price = $8,
		actual_price = $9,
		discount_percentage = $10,
		discount_amount = $11,
		original_tax = $12,
		actual_tax = $13,
		invoice_date = $14,
		due_date = $15,
		paid = $16,
		status = $17,
		payment_date = $18,
		updated_by = $19,
		updated_at = $20
		where id = $21`

	_, err := db.ExecContext(ctx, stmt,
		invoice.CompanyId,
		invoice.ProjectId,
		invoice.SubProjectId,
		invoice.InvoiceDisplayName,
		invoice.InvoiceDescription,
		invoice.StatisticsInvoice,
		invoice.InvoiceItems,
		invoice.OriginalPrice,
		invoice.ActualPrice,
		invoice.DiscountPercentage,
		invoice.DiscountAmount,
		invoice.OriginalTax,
		invoice.ActualTax,
		invoice.InvoiceDate,
		invoice.DueDate,
		invoice.Paid,
		invoice.Status,
		invoice.PaymentDate,
		userId,
		time.Now(),
		invoice.ID,
	)
	if err != nil {
		log.Println("Error updating invoice", err)
		return err
	}

	return nil
}

func UpdateIncomeId(incomeId string, invoiceId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update invoices set 
		income_id = $1
		where id = $2`

	_, err := db.ExecContext(ctx, stmt,
		incomeId,
		invoiceId,
	)
	if err != nil {
		log.Println("Error updating invoice", err)
		return err
	}

	return nil
}

func GetAllInvoicesByIds(ids string) ([]*InvoicePostgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select
		id,
		company_id,
		project_id,
		sub_project_id,
		income_id,
		invoice_display_name,
		invoice_description,
		statistics_invoice,
		invoice_items,
		original_price,
		actual_price,
		discount_percentage,
		discount_amount,
		original_tax,
		actual_tax,
		invoice_date,
		due_date,
		payment_date,
		paid,
		status,
		created_by,
		created_at, 
		updated_by, 
		updated_at
        from invoices
        where id = ANY($1)
		order by created_at asc
    `

	rows, err := db.QueryContext(ctx, query, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*InvoicePostgres

	for rows.Next() {
		var invoice InvoicePostgres
		err := rows.Scan(
			&invoice.ID,
			&invoice.CompanyId,
			&invoice.ProjectId,
			&invoice.SubProjectId,
			&invoice.IncomeId,
			&invoice.InvoiceDisplayName,
			&invoice.InvoiceDescription,
			&invoice.StatisticsInvoice,
			&invoice.InvoiceItems,
			&invoice.OriginalPrice,
			&invoice.ActualPrice,
			&invoice.DiscountPercentage,
			&invoice.DiscountAmount,
			&invoice.OriginalTax,
			&invoice.ActualTax,
			&invoice.InvoiceDate,
			&invoice.DueDate,
			&invoice.PaymentDate,
			&invoice.Paid,
			&invoice.Status,
			&invoice.CreatedBy,
			&invoice.CreatedAt,
			&invoice.UpdatedBy,
			&invoice.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}
