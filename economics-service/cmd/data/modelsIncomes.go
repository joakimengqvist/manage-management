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

	query := `select id, project_id, income_date, income_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, updated_by, updated_at
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
			&income.ProjectID,
			&income.IncomeDate,
			&income.IncomeCategory,
			&income.Vendor,
			&income.Description,
			&income.Amount,
			&income.Tax,
			&income.Status,
			&income.Currency,
			&income.PaymentMethod,
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

	query := `select id, project_id, income_date, income_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, updated_by, updated_at
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
			&income.ProjectID,
			&income.IncomeDate,
			&income.IncomeCategory,
			&income.Vendor,
			&income.Description,
			&income.Amount,
			&income.Tax,
			&income.Status,
			&income.Currency,
			&income.PaymentMethod,
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

func InsertIncome(income NewIncome) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `
		insert into incomes (
            project_id,
            income_date,
            income_category,
            vendor,
            description,
            amount,
			tax,
			status,
            currency,
            payment_method,
            created_by,
            created_at,
            updated_by,
            updated_at
        )
        values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) returning id`

	err := db.QueryRowContext(ctx, stmt,
		income.ProjectID,
		income.IncomeDate,
		income.IncomeCategory,
		income.Vendor,
		income.Description,
		income.Amount,
		income.Tax,
		income.Status,
		income.Currency,
		income.PaymentMethod,
		income.CreatedBy,
		time.Now(),
		income.UpdatedBy,
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}

func (u *Income) UpdateIncome(updatedByUserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update incomes set
		project_id = $1,
		income_date = $2,
		income_category = $3,
		vendor = $4,
		description = $5,
		amount = $6,
		tax = $7,
		status = $8,
		currency = $9,
		payment_method = $10,
		updated_by = $11,
		updated_at = $12
		where id = $13
	`

	_, err := db.ExecContext(ctx, stmt,
		u.ProjectID,
		u.IncomeDate,
		u.IncomeCategory,
		u.Vendor,
		u.Description,
		u.Amount,
		u.Tax,
		u.Status,
		u.Currency,
		u.PaymentMethod,
		u.UpdatedBy,
		time.Now(),
		u.ID,
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

	query := `select id, project_id, income_date, income_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, updated_by, updated_at
	from incomes where id = $1`

	row := db.QueryRowContext(ctx, query, IncomeId)

	var income Income
	err := row.Scan(
		&income.ID,
		&income.ProjectID,
		&income.IncomeDate,
		&income.IncomeCategory,
		&income.Vendor,
		&income.Description,
		&income.Amount,
		&income.Tax,
		&income.Status,
		&income.Currency,
		&income.PaymentMethod,
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
