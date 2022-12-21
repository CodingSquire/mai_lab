package repository

import (
	"fmt"
	"go_services_lab/models"

	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product models.Product) (id int, err error) {
	query := fmt.Sprintf("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id")
	row := r.db.QueryRow(query, product.Name, product.Price)

	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ProductPostgres) GetAll() (products []models.Product, err error) {
	var product models.Product

	rows, err := r.db.Queryx("SELECT * FROM products")
	for rows.Next() {
		err = rows.StructScan(&product)
		if err == nil {
			products = append(products, product)
		}
	}

	return products, err
}

func (r *ProductPostgres) LastOne() (product models.Product, err error) {
	err = r.db.Get(&product, "SELECT * FROM products WHERE id IN (SELECT MAX(id) FROM 	products)")
	return product, err
}
