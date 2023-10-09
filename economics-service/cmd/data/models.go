package data

import (
	"context"
	"database/sql"
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
		NewProjectExpense: NewProjectExpense{},
		ProjectExpense:    ProjectExpense{},
		NewProjectIncome:  NewProjectIncome{},
		ProjectIncome:     ProjectIncome{},
	}
}

type Models struct {
	ProjectExpense    ProjectExpense
	NewProjectExpense NewProjectExpense
	NewProjectIncome  NewProjectIncome
	ProjectIncome     ProjectIncome
}

type ExpenseId struct {
	ExpenseId string `json:"expense_id"`
}

type NewProjectExpense struct {
	ProjectID       string    `json:"project_id"`
	ExpenseDate     time.Time `json:"expense_date"`
	ExpenseCategory string    `json:"expense_category"`
	Vendor          string    `json:"vendor"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Tax             float64   `json:"tax"`
	Status          string    `json:"status"`
	Currency        string    `json:"currency"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedBy       string    `json:"created_by"`
	ModifiedBy      string    `json:"modified_by"`
}

type NewProjectIncome struct {
	ProjectID      string    `json:"project_id"`
	IncomeDate     time.Time `json:"income_date"`
	IncomeCategory string    `json:"income_category"`
	Vendor         string    `json:"vendor"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Tax            float64   `json:"tax"`
	Status         string    `json:"status"`
	Currency       string    `json:"currency"`
	PaymentMethod  string    `json:"payment_method"`
	CreatedBy      string    `json:"created_by"`
	ModifiedBy     string    `json:"modified_by"`
}

type ProjectExpense struct {
	ExpenseID       string    `json:"expense_id"`
	ProjectID       string    `json:"project_id"`
	ExpenseDate     time.Time `json:"expense_date"`
	ExpenseCategory string    `json:"expense_category"`
	Vendor          string    `json:"vendor"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Tax             float64   `json:"tax"`
	Status          string    `json:"status"`
	Currency        string    `json:"currency"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	ModifiedBy      string    `json:"modified_by"`
	ModifiedAt      time.Time `json:"modified_at"`
}

type ProjectIncome struct {
	IncomeID       string    `json:"income_id"`
	ProjectID      string    `json:"project_id"`
	IncomeDate     time.Time `json:"income_date"`
	IncomeCategory string    `json:"income_category"`
	Vendor         string    `json:"vendor"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	Tax            float64   `json:"tax"`
	Status         string    `json:"status"`
	Currency       string    `json:"currency"`
	PaymentMethod  string    `json:"payment_method"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	ModifiedBy     string    `json:"modified_by"`
	ModifiedAt     time.Time `json:"modified_at"`
}

// -----------------------------------
// ----------- EXPENSES --------------
// -----------------------------------

func GetAllProjectExpenses() ([]*ProjectExpense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, project_id, expense_date, expense_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, modified_by, modified_at
	from project_expenses order by expense_date desc`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*ProjectExpense

	for rows.Next() {
		var expense ProjectExpense
		err := rows.Scan(
			&expense.ExpenseID,
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
			&expense.ModifiedBy,
			&expense.ModifiedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		expenses = append(expenses, &expense)
	}

	return expenses, nil
}

func GetAllProjectExpensesByProjectId(projectId string) ([]*ProjectExpense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, project_id, expense_date, expense_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, modified_by, modified_at
	from project_expenses where project_id = $1 order by expense_date desc`

	rows, err := db.QueryContext(ctx, query, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*ProjectExpense

	for rows.Next() {
		var expense ProjectExpense
		err := rows.Scan(
			&expense.ExpenseID,
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
			&expense.ModifiedBy,
			&expense.ModifiedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		expenses = append(expenses, &expense)
	}

	return expenses, nil
}

func InsertExpense(expense NewProjectExpense) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `
		insert into project_expenses (
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
            modified_by,
            modified_at
        )
        values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14 ) returning id`

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
		expense.CreatedBy,
		time.Now(),
		expense.ModifiedBy,
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}

func GetExpenseById(ExpenseId string) (*ProjectExpense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, project_id, expense_date, expense_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, modified_by, modified_at
	from project_expenses where id = $1`

	row := db.QueryRowContext(ctx, query, ExpenseId)

	var expense ProjectExpense
	err := row.Scan(
		&expense.ExpenseID,
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
		&expense.ModifiedBy,
		&expense.ModifiedAt,
	)

	if err != nil {
		return nil, err
	}

	return &expense, nil
}

// -----------------------------------
// ----------- INCOME ----------------
// -----------------------------------

func GetAllProjectIncomes() ([]*ProjectIncome, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, project_id, income_date, income_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, modified_by, modified_at
	from project_incomes order by income_date desc`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incomes []*ProjectIncome

	for rows.Next() {
		var income ProjectIncome
		err := rows.Scan(
			&income.IncomeID,
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
			&income.ModifiedBy,
			&income.ModifiedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		incomes = append(incomes, &income)
	}

	return incomes, nil
}

func GetAllProjectIncomesByProjectId(projectId string) ([]*ProjectIncome, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, project_id, income_date, income_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, modified_by, modified_at
	from project_incomes where project_id = $1 order by income_date desc`

	rows, err := db.QueryContext(ctx, query, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incomes []*ProjectIncome

	for rows.Next() {
		var income ProjectIncome
		err := rows.Scan(
			&income.IncomeID,
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
			&income.ModifiedBy,
			&income.ModifiedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		incomes = append(incomes, &income)
	}

	return incomes, nil
}

func InsertIncome(income NewProjectIncome) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `
		insert into project_incomes (
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
            modified_by,
            modified_at
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
		income.ModifiedBy,
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}

func GetIncomeById(IncomeId string) (*ProjectIncome, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, project_id, income_date, income_category, vendor, description, amount, tax, status, currency, payment_method, created_by, created_at, modified_by, modified_at
	from project_incomes where id = $1`

	row := db.QueryRowContext(ctx, query, IncomeId)

	var income ProjectIncome
	err := row.Scan(
		&income.IncomeID,
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
		&income.ModifiedBy,
		&income.ModifiedAt,
	)

	if err != nil {
		return nil, err
	}

	return &income, nil
}
