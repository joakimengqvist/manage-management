package data

import (
	"context"
	"fmt"
	"log"
	"time"
)

func GetAllExpenses() ([]*Expense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, project_id, expense_date, expense_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, updated_by, updated_at
	from expenses order by expense_date desc`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*Expense

	for rows.Next() {
		var expense Expense
		err := rows.Scan(
			&expense.ID,
			&expense.ProjectID,
			&expense.ExpenseDate,
			&expense.ExpenseCategory,
			&expense.Vendor,
			&expense.Description,
			&expense.Amount,
			&expense.Tax,
			&expense.Status,
			&expense.Currency,
			&expense.PaymentMethod,
			&expense.CreatedBy,
			&expense.CreatedAt,
			&expense.UpdatedBy,
			&expense.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		expenses = append(expenses, &expense)
	}

	return expenses, nil
}

func GetAllExpensesByProjectId(projectId string) ([]*Expense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, project_id, expense_date, expense_category, vendor, description, amount, tax, status, currency, payment_method,created_by, created_at, updated_by, updated_at
	from expenses where project_id = $1 order by expense_date desc`

	rows, err := db.QueryContext(ctx, query, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*Expense

	for rows.Next() {
		var expense Expense
		err := rows.Scan(
			&expense.ID,
			&expense.ProjectID,
			&expense.ExpenseDate,
			&expense.ExpenseCategory,
			&expense.Vendor,
			&expense.Description,
			&expense.Amount,
			&expense.Tax,
			&expense.Status,
			&expense.Currency,
			&expense.PaymentMethod,
			&expense.CreatedBy,
			&expense.CreatedAt,
			&expense.UpdatedBy,
			&expense.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		expenses = append(expenses, &expense)
	}

	return expenses, nil
}

func InsertExpense(expense NewExpense, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `
		insert into expenses (
            project_id,
            expense_date,
            expense_category,
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
		expense.ProjectID,
		expense.ExpenseDate,
		expense.ExpenseCategory,
		expense.Vendor,
		expense.Description,
		expense.Amount,
		expense.Tax,
		expense.Status,
		expense.Currency,
		expense.PaymentMethod,
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

func (u *Expense) UpdateExpense(updatedByUserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update expenses set
		project_id = $1,
		expense_date = $2,
		expense_category = $3,
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
		u.ExpenseDate,
		u.ExpenseCategory,
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

func GetExpenseById(ExpenseId string) (*Expense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, project_id, expense_date, expense_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, updated_by, updated_at
	from expenses where id = $1`

	row := db.QueryRowContext(ctx, query, ExpenseId)

	var expense Expense
	err := row.Scan(
		&expense.ID,
		&expense.ProjectID,
		&expense.ExpenseDate,
		&expense.ExpenseCategory,
		&expense.Vendor,
		&expense.Description,
		&expense.Amount,
		&expense.Tax,
		&expense.Status,
		&expense.Currency,
		expense.PaymentMethod,
		&expense.CreatedBy,
		&expense.CreatedAt,
		&expense.UpdatedBy,
		&expense.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &expense, nil
}
