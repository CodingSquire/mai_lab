package controllers

import (
	"database/sql"
	"log"
	"orders/internal/models"
	"sync"

	"github.com/lucsky/cuid"
)

type OrderPgxController struct {
	db  *sql.DB
	mut sync.Mutex
}

// GetAllOrdersByUserId implements OrderController
func (o *OrderPgxController) GetAllOrdersByUserId(userId string) []models.Order {
	o.mut.Lock()
	defer o.mut.Unlock()

	res := make([]models.Order, 1)

	rows, err := o.db.Query("SELECT id, userId, item, adress, count, createdAt, updatedAt FROM orders WHERE userId = $1", userId)
	if err != nil {
		log.Print(err)
		return res
	}

	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.ID,
			&order.UserID,
			&order.Item,
			&order.Address,
			&order.Count,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			log.Print(err)
			return res
		}
		res = append(res, order)
	}

	return res
}

// DeleteOrderById implements OrderController
func (o *OrderPgxController) DeleteOrderById(id string) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	_, err = o.db.Exec("DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		log.Print(err)
		return
	}

	return
}

// GetAllOrders implements OrderController
func (o *OrderPgxController) GetAllOrders() []models.Order {
	o.mut.Lock()
	defer o.mut.Unlock()

	res := make([]models.Order, 0)

	rows, err := o.db.Query("SELECT id, userId, item, adress, count, createdAt, updatedAt FROM orders")
	if err != nil {
		log.Print(err)
		return res
	}

	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.ID,
			&order.UserID,
			&order.Item,
			&order.Address,
			&order.Count,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			log.Print(err)
			return res
		}
		res = append(res, order)
	}

	return res
}

// GetOrderById implements OrderController
func (o *OrderPgxController) GetOrderById(id string) (*models.Order, error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	var order models.Order
	err := o.db.QueryRow("SELECT id, userId, item, adress, count, createdAt, updatedAt FROM orders WHERE id = $1", id).Scan(
		&order.ID,
		&order.UserID,
		&order.Item,
		&order.Address,
		&order.Count,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// PatchOrderById implements OrderController
func (o *OrderPgxController) PatchOrderById(id string, order *models.Order) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()


	_, err = o.db.Exec(
		"UPDATE orders SET userId = $1, item = $2, adress = $3, count = $4, updatedAt = NOW() WHERE id = $5",
		order.UserID,
		order.Item,
		order.Address,
		order.Count,
		id,
	)

	if err != nil {
		log.Print(err)
		return
	}

	return
}

// PostOrder implements OrderController
func (o *OrderPgxController) PostOrder(order *models.Order) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	if order.ID == "" {
		order.ID = cuid.New()
	}

	_, err = o.db.Exec(
		"INSERT INTO orders (id, userId, item, adress, count) VALUES ($1, $2, $3, $4, $5)",
		order.ID,
		order.UserID,
		order.Item,
		order.Address,
		order.Count,
	)

	return
}

func NewPgxController(db *sql.DB) OrderController {
	return &OrderPgxController{
		db:  db,
		mut: sync.Mutex{},
	}
}
