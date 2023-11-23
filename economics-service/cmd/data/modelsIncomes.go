package data

import (
	"context"
	"fmt"
	"log"
	"time"
)

func GetAllIncomes() ([]*Income, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, invoice_id, project_id, income_date, income_category, statistics_income, vendor, description, amount, tax, status, currency, created_by, created_at, updated_by, updated_at
	from incomes order by income_date desc`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incomes []*Income

	for rows.Next() {
		var income Income
		err := rows.Scan(
			&income.ID,
			&income.InvoiceID,
			&income.ProjectID,
			&income.IncomeDate,
			&income.IncomeCategory,
			&income.StatisticsIncome,
			&income.Vendor,
			&income.Description,
			&income.Amount,
			&income.Tax,
			&income.Status,
			&income.Currency,
			&income.CreatedBy,
			&income.CreatedAt,
			&income.UpdatedBy,
			&income.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		incomes = append(incomes, &income)
	}

	return incomes, nil
}

func GetAllProjectIncomesByProjectId(projectId string) ([]*Income, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, invoice_id, project_id, income_date, income_category, statistics_income, vendor, description, amount, tax, status, currency, created_by, created_at, updated_by, updated_at
	from incomes where project_id = $1 order by income_date desc`

	rows, err := db.QueryContext(ctx, query, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incomes []*Income

	for rows.Next() {
		var income Income
		err := rows.Scan(
			&income.ID,
			&income.InvoiceID,
			&income.ProjectID,
			&income.IncomeDate,
			&income.IncomeCategory,
			&income.StatisticsIncome,
			&income.Vendor,
			&income.Description,
			&income.Amount,
			&income.Tax,
			&income.Status,
			&income.Currency,
			&income.CreatedBy,
			&income.CreatedAt,
			&income.UpdatedBy,
			&income.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		incomes = append(incomes, &income)
	}

	return incomes, nil
}

func InsertIncome(income NewIncome, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `
		insert into incomes (
            project_id,
			invoice_id,
            income_date,
            income_category,
			statistics_income,
            vendor,
            description,
            amount,
			tax,
			status,
            currency,
            created_by,
            created_at,
            updated_by,
            updated_at
        )
        values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) returning id`

	err := db.QueryRowContext(ctx, stmt,
		income.ProjectID,
		income.InvoiceID,
		income.IncomeDate,
		income.IncomeCategory,
		income.StatisticsIncome,
		income.Vendor,
		income.Description,
		income.Amount,
		income.Tax,
		income.Status,
		income.Currency,
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

func (income *Income) UpdateIncome(updatedByUserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update incomes set
		project_id = $1,
		invoice_id = $2
		income_date = $3,
		income_category = $4,
		statistics_income = $5,
		vendor = $6,
		description = $7,
		amount = $8,
		tax = $9,
		status = $10,
		currency = $11,
		updated_by = $12,
		updated_at = $13
		where id = $14
	`

	_, err := db.ExecContext(ctx, stmt,
		income.ProjectID,
		income.InvoiceID,
		income.IncomeDate,
		income.IncomeCategory,
		income.StatisticsIncome,
		income.Vendor,
		income.Description,
		income.Amount,
		income.Tax,
		income.Status,
		income.Currency,
		income.UpdatedBy,
		time.Now(),
		income.ID,
	)

	if err != nil {
		fmt.Println("Error updating project", err)
		return err
	}

	return nil
}

func GetIncomeById(IncomeId string) (*Income, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, invoice_id, project_id, income_date, income_category, statistics_income, vendor, description, amount, tax, status, currency, created_by, created_at, updated_by, updated_at
	from incomes where id = $1`

	row := db.QueryRowContext(ctx, query, IncomeId)

	var income Income
	err := row.Scan(
		&income.ID,
		&income.InvoiceID,
		&income.ProjectID,
		&income.IncomeDate,
		&income.IncomeCategory,
		&income.StatisticsIncome,
		&income.Vendor,
		&income.Description,
		&income.Amount,
		&income.Tax,
		&income.Status,
		&income.Currency,
		&income.CreatedBy,
		&income.CreatedAt,
		&income.UpdatedBy,
		&income.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &income, nil
}
