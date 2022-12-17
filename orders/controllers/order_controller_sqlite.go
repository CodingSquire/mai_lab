package controllers

import (
	"database/sql"
	"orders/models"
	"sync"
)

type OrderSqliteController struct {
	db  *sql.DB
	mut sync.Mutex
}

// GetAllOrdersByUserId implements OrderController
func (o *OrderSqliteController) GetAllOrdersByUserId(userId string) []models.Order {
	o.mut.Lock()
	defer o.mut.Unlock()

	res := make([]models.Order, 1)

	return res
}

// DeleteOrderById implements OrderController
func (o *OrderSqliteController) DeleteOrderById(id string) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	return
}

// GetAllOrders implements OrderController
func (o *OrderSqliteController) GetAllOrders() []models.Order {
	o.mut.Lock()
	defer o.mut.Unlock()

	res := make([]models.Order, 1)

	return res
}

// GetOrderById implements OrderController
func (o *OrderSqliteController) GetOrderById(id string) (*models.Order, bool) {
	o.mut.Lock()
	defer o.mut.Unlock()

	return nil, false
}

// PatchOrderById implements OrderController
func (o *OrderSqliteController) PatchOrderById(id string, order *models.Order) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	return
}

// PostOrder implements OrderController
func (o *OrderSqliteController) PostOrder(order *models.Order) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	return
}

func NewSqliteController(db *sql.DB) OrderController {
	return &OrderSqliteController{
		db:  db,
		mut: sync.Mutex{},
	}
}
