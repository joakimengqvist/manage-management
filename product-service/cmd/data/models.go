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
		Product:    Product{},
		NewProduct: NewProduct{},
	}
}

type Models struct {
	Product    Product
	NewProduct NewProduct
}

type Product struct {
	ID            string    `json:"id"`
	Name          string    `json:"name" sql:"not null"`
	Description   string    `json:"description"`
	Category      string    `json:"category"`
	Price         float64   `json:"price"`
	TaxPercentage float64   `json:"tax_percentage"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	CreatedBy     string    `json:"created_by,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	UpdatedBy     string    `json:"updated_by,omitempty"`
}

type NewProduct struct {
	Name          string  `json:"name" sql:"not null"`
	Description   string  `json:"description"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	TaxPercentage float64 `json:"tax_percentage"`
	CreatedBy     string  `json:"created_by,omitempty"`
}

func InsertProduct(product NewProduct) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string

	stmt := `insert into products (name, description, category, price, tax_percentage, created_at, created_by, updated_at, updated_by)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		returning id`

	err := db.QueryRowContext(ctx, stmt,
		product.Name,
		product.Description,
		product.Category,
		product.Price,
		product.TaxPercentage,
		time.Now(),
		product.CreatedBy,
		time.Now(),
		product.CreatedBy,
	).Scan(&newID)

	if err != nil {
		return "error inserting product", err
	}

	return newID, nil
}

func UpdateProduct(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update products set
		name = $1,
		description = $2,
		category = $3,
		price = $4,
		tax_percentage = $5,
		updated_at = $6,
		updated_by = $7,
		where id = $8
	`

	_, err := db.ExecContext(ctx, stmt,
		product.Name,
		product.Description,
		product.Category,
		product.Price,
		product.TaxPercentage,
		time.Now(),
		product.UpdatedBy,
		product.ID,
	)

	if err != nil {
		log.Println("Error updating product", err)
		return err
	}

	return nil
}

func GetAllProducts() ([]*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, category, price, tax_percentage, created_at, created_by, updated_at, updated_by
	from products order by updated_at desc`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Category,
			&product.Price,
			&product.TaxPercentage,
			&product.CreatedAt,
			&product.CreatedBy,
			&product.UpdatedAt,
			&product.UpdatedBy,
		)
		if err != nil {
			log.Println("Error scanning for products", err)
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func GetProductById(id string) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, category, price, tax_percentage, created_at, created_by, updated_at, updated_by 
	from products where id = $1`

	var product Product
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Category,
		&product.Price,
		&product.TaxPercentage,
		&product.CreatedAt,
		&product.CreatedBy,
		&product.UpdatedAt,
		&product.UpdatedBy,
	)

	if err != nil {
		log.Println("GetProductById model", err)
		return nil, err
	}

	return &product, nil
}
