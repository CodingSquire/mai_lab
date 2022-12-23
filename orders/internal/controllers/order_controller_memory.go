package controllers

import (
	"errors"
	"github.com/innerave/mai_lab/orders/internal/models"
	"sync"
	"time"

	"github.com/lucsky/cuid"
)

type OrderMemController struct {
	cache map[string]*models.Order
	mut   sync.Mutex
}

// GetAllOrdersByUserId implements OrderController
func (o *OrderMemController) GetAllOrdersByUserId(userId string) []models.Order {
	o.mut.Lock()
	defer o.mut.Unlock()

	orders := make([]models.Order, 0)

	for _, order := range o.cache {
		if order.UserID == userId {
			orders = append(orders, *order)
		}
	}

	return orders
}

// NewOrderMemController gets new in-memory order controller
func NewOrderMemController() OrderController {
	return &OrderMemController{
		cache: make(map[string]*models.Order),
		mut:   sync.Mutex{},
	}
}

// GetAllOrders gets all orders as array
func (o *OrderMemController) GetAllOrders() []models.Order {
	o.mut.Lock()
	defer o.mut.Unlock()

	orders := make([]models.Order, 0)

	for _, order := range o.cache {
		orders = append(orders, *order)
	}

	return orders
}

// GetOrderById gets order with id
// or returns (, false) if one doesn't exists
func (o *OrderMemController) GetOrderById(id string) (*models.Order, error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	order, ok := o.cache[id]
	if !ok {
		return nil, errors.New("Order not found")
	}
	return order, nil
}

// DeleteOrderById deletes order with id if exists
func (o *OrderMemController) DeleteOrderById(id string) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	delete(o.cache, id)
	return
}

// PatchOrderById sets new order body on id
// Also update UpdatedAt field with Now()
func (o *OrderMemController) PatchOrderById(id string, order *models.Order) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	got_order, ok := o.cache[id]
	if !ok {
		err = errors.New("Order not found")
		return
	}

	order.ID = id
	order.CreatedAt = got_order.CreatedAt
	order.UpdatedAt = time.Now()
	o.cache[id] = order
	return
}

// PostOrder setting new order in memory
func (o *OrderMemController) PostOrder(order *models.Order) (*models.Order, error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	timeNow := time.Now()

	if order.ID == "" {
		order.ID = cuid.New()
	}
	order.CreatedAt = timeNow
	order.UpdatedAt = timeNow
	o.cache[order.ID] = order

	return order, nil
}
