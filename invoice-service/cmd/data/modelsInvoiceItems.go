package data

import (
	"context"
	"fmt"
	"log"
	"time"
)

func InsertInvoiceItem(invoiceItem InvoiceItem, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string

	stmt := `insert into invoice_items (
		product_id,
		quantity,
		original_price,
		actual_price,
		discount_percentage,
		discount_amount,
		tax_percentage,
		original_tax,
		actual_tax,
		created_by,
		created_at,
		updated_by,
		updated_at
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) returning id`

	err := db.QueryRowContext(ctx, stmt,
		invoiceItem.ProductId,
		invoiceItem.Quantity,
		invoiceItem.OriginalPrice,
		invoiceItem.ActualPrice,
		invoiceItem.DiscountPercentage,
		invoiceItem.DiscountAmount,
		invoiceItem.TaxPercentage,
		invoiceItem.OriginalTax,
		invoiceItem.ActualTax,
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

func GetInvoiceItemById(invoiceItemId string) (*InvoiceItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select 
		id,
		product_id,
		quantity,
		original_price,
		actual_price,
		discount_percentage,
		discount_amount,
		tax_percentage,
		original_tax,
		actual_tax,
		created_by,
		created_at,
		updated_by,
		updated_at
		from invoice_items where id = $1`

	row := db.QueryRowContext(ctx, query, invoiceItemId)

	var invoiceItem InvoiceItem
	err := row.Scan(
		&invoiceItem.ID,
		&invoiceItem.ProductId,
		&invoiceItem.Quantity,
		&invoiceItem.OriginalPrice,
		&invoiceItem.ActualPrice,
		&invoiceItem.DiscountPercentage,
		&invoiceItem.DiscountAmount,
		&invoiceItem.TaxPercentage,
		&invoiceItem.OriginalTax,
		&invoiceItem.ActualTax,
		&invoiceItem.CreatedBy,
		&invoiceItem.CreatedAt,
		&invoiceItem.UpdatedBy,
		&invoiceItem.UpdatedAt,
	)
	if err != nil {
		log.Println("Error scanning", err)
		return nil, err
	}

	return &invoiceItem, nil
}

func GetInvoiceItemsByInvoiceId(invoiceId string) ([]*InvoiceItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select 
		id,
		product_id,
		quantity,
		original_price,
		actual_price,
		discount_percentage,
		discount_amount,
		tax_percentage,
		original_tax,
		actual_tax,
		created_by,
		created_at,
		updated_by,
		updated_at
		from invoice_items where invoice_id = $1`

	rows, err := db.QueryContext(ctx, query, invoiceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoiceItems []*InvoiceItem

	for rows.Next() {
		var invoiceItem InvoiceItem
		err := rows.Scan(
			&invoiceItem.ID,
			&invoiceItem.ProductId,
			&invoiceItem.Quantity,
			&invoiceItem.OriginalPrice,
			&invoiceItem.ActualPrice,
			&invoiceItem.DiscountPercentage,
			&invoiceItem.DiscountAmount,
			&invoiceItem.TaxPercentage,
			&invoiceItem.OriginalTax,
			&invoiceItem.ActualTax,
			&invoiceItem.CreatedBy,
			&invoiceItem.CreatedAt,
			&invoiceItem.UpdatedBy,
			&invoiceItem.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		invoiceItems = append(invoiceItems, &invoiceItem)
	}

	return invoiceItems, nil
}

func GetAllInvoiceItems() ([]*InvoiceItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select 
		id, 
		product_id,
		quantity,
		original_price,
		actual_price,
		discount_percentage,
		discount_amount,
		tax_percentage,
		original_tax,
		actual_tax,
		created_by,
		created_at,
		updated_by,
		updated_at
		from invoice_items`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoiceItems []*InvoiceItem

	for rows.Next() {
		var invoiceItem InvoiceItem
		err := rows.Scan(
			&invoiceItem.ID,
			&invoiceItem.ProductId,
			&invoiceItem.Quantity,
			&invoiceItem.OriginalPrice,
			&invoiceItem.ActualPrice,
			&invoiceItem.DiscountPercentage,
			&invoiceItem.DiscountAmount,
			&invoiceItem.TaxPercentage,
			&invoiceItem.OriginalTax,
			&invoiceItem.ActualTax,
			&invoiceItem.CreatedBy,
			&invoiceItem.CreatedAt,
			&invoiceItem.UpdatedBy,
			&invoiceItem.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		invoiceItems = append(invoiceItems, &invoiceItem)
	}

	return invoiceItems, nil
}

func GetAllInvoiceItemsByIds(ids string) ([]*InvoiceItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select
		id, 
		product_id,
		quantity,
		original_price,
		actual_price,
		discount_percentage,
		discount_amount,
		tax_percentage,
		original_tax,
		actual_tax,
		created_by,
		created_at,
		updated_by,
		updated_at
        from invoice_items
        where id = ANY($1)
		order by created_at asc
    `

	rows, err := db.QueryContext(ctx, query, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoiceItems []*InvoiceItem

	for rows.Next() {
		var invoiceItem InvoiceItem
		err := rows.Scan(
			&invoiceItem.ID,
			&invoiceItem.ProductId,
			&invoiceItem.Quantity,
			&invoiceItem.OriginalPrice,
			&invoiceItem.ActualPrice,
			&invoiceItem.DiscountPercentage,
			&invoiceItem.DiscountAmount,
			&invoiceItem.TaxPercentage,
			&invoiceItem.OriginalTax,
			&invoiceItem.ActualTax,
			&invoiceItem.CreatedBy,
			&invoiceItem.CreatedAt,
			&invoiceItem.UpdatedBy,
			&invoiceItem.UpdatedAt,
		)
		if err != nil {
			fmt.Println("Error scanning", err)
			return nil, err
		}

		invoiceItems = append(invoiceItems, &invoiceItem)
	}

	return invoiceItems, nil
}

func UpdateInvoiceItem(invoiceItem InvoiceItem, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update invoice_items set 
		product_id = $1,
		quantity = $2,
		original_price = $3,
		actual_price = $4,
		discount_percentage = $5,
		discount_amount = $6,
		tax_percentage = $7,
		original_tax = $8,
		actual_tax = $9,
		updated_by = $10,
		updated_at = $11
		where id = $12`

	_, err := db.ExecContext(ctx, stmt,
		invoiceItem.ProductId,
		invoiceItem.Quantity,
		invoiceItem.OriginalPrice,
		invoiceItem.ActualPrice,
		invoiceItem.DiscountPercentage,
		invoiceItem.DiscountAmount,
		invoiceItem.TaxPercentage,
		invoiceItem.OriginalTax,
		invoiceItem.ActualTax,
		userId,
		time.Now(),
		invoiceItem.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetAllInvoiceItemsByProductId(productId string) ([]*InvoiceItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select 
		id, 
		product_id,
		quantity,
		original_price,
		actual_price,
		discount_percentage,
		discount_amount,
		tax_percentage,
		original_tax,
		actual_tax,
		created_by,
		created_at,
		updated_by,
		updated_at 
		from invoice_items where product_id = $1`

	rows, err := db.QueryContext(ctx, query, productId)
	if err != nil {
		log.Println("Error querying", err)
	}
	defer rows.Close()

	var invoiceItems []*InvoiceItem

	for rows.Next() {
		var invoiceItem InvoiceItem
		err := rows.Scan(
			&invoiceItem.ID,
			&invoiceItem.ProductId,
			&invoiceItem.Quantity,
			&invoiceItem.OriginalPrice,
			&invoiceItem.ActualPrice,
			&invoiceItem.DiscountPercentage,
			&invoiceItem.DiscountAmount,
			&invoiceItem.TaxPercentage,
			&invoiceItem.OriginalTax,
			&invoiceItem.ActualTax,
			&invoiceItem.CreatedBy,
			&invoiceItem.CreatedAt,
			&invoiceItem.UpdatedBy,
			&invoiceItem.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
		}

		invoiceItems = append(invoiceItems, &invoiceItem)
	}

	return invoiceItems, nil
}
