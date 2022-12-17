package controllers

import (
	"database/sql"
	"log"
	"orders/models"
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

	rows, err := o.db.Query("SELECT id, item, adress, count, createdAt, updatedAt FROM orders WHERE userId = $1", userId)
	if err != nil {
		log.Fatal(err)
		return res
	}

	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.Id,
			&order.Item,
			&order.Adress,
			&order.Count,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
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
		log.Fatal(err)
		return
	}

	return
}

// GetAllOrders implements OrderController
func (o *OrderPgxController) GetAllOrders() []models.Order {
	o.mut.Lock()
	defer o.mut.Unlock()

	res := make([]models.Order, 1)

	rows, err := o.db.Query("SELECT id, item, adress, count, createdAt, updatedAt FROM orders")
	if err != nil {
		log.Fatal(err)
		return res
	}

	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.Id,
			&order.Item,
			&order.Adress,
			&order.Count,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
			return res
		}
		res = append(res, order)
	}

	return res
}

// GetOrderById implements OrderController
func (o *OrderPgxController) GetOrderById(id string) (*models.Order, bool) {
	o.mut.Lock()
	defer o.mut.Unlock()

	var order models.Order
	err := o.db.QueryRow("SELECT id, item, adress, count, createdAt, updatedAt FROM orders WHERE id = $1", id).Scan(
		&order.Id,
		&order.Item,
		&order.Adress,
		&order.Count,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	return &order, true
}

// PatchOrderById implements OrderController
func (o *OrderPgxController) PatchOrderById(id string, order *models.Order) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	if order.UserId != "" {
		_, err = o.db.Exec("UPDATE orders SET userId = $1 WHERE id = $2", order.UserId, id)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	if order.Item != "" {
		_, err = o.db.Exec("UPDATE orders SET item = $1 WHERE id = $2", order.Item, id)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	if order.Adress != "" {
		_, err = o.db.Exec("UPDATE orders SET adress = $1 WHERE id = $2", order.Adress, id)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	if order.Count != 0 {
		_, err = o.db.Exec("UPDATE orders SET count = $1 WHERE id = $2", order.Count, id)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	// FIXME: cockroach cannot? have triggers
	_, err = o.db.Exec("UPDATE orders SET updatedAt = NOW() WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

// PostOrder implements OrderController
func (o *OrderPgxController) PostOrder(order *models.Order) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	order.Id = cuid.New()

	_, err = o.db.Exec(
		"INSERT INTO orders (id, userId, item, adress, count, updatedAt) VALUES ($1, $2, $3, $4, $5, NOW())",
		order.Id,
		order.UserId,
		order.Item,
		order.Adress,
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
