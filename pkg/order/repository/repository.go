package repository

import (
	"go_services_lab/models"

	"github.com/jmoiron/sqlx"
)

type Product interface {
	Create(product models.Product) (int, error)
	GetAll() ([]models.Product, error)
	LastOne() (models.Product, error)
}

type Order interface {
	Get(id int) (models.Order, error)
	GetAll() ([]models.Order, error)
	Amount(id int) (float32, error)
	Delete(id int) (int, error)
	Create(user_id int, products map[int]int) (int, error)
}

type OrderRepository struct {
	Product
	Order
}

func NewRepositoryOrder(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		Product: NewProductPostgres(db),
		Order:   NewOrderPostgres(db),
	}
}
