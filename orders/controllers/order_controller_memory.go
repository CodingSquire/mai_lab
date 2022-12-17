package controllers

import (
	"orders/models"
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
		if order.UserId == userId {
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
func (o *OrderMemController) GetOrderById(id string) (*models.Order, bool) {
	o.mut.Lock()
	defer o.mut.Unlock()

	order, ok := o.cache[id]
	return order, ok
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

	gotOrder, ok := o.cache[id]

	if ok {
		timeNow := time.Now().Nanosecond()

		if order.Item != "" {
			gotOrder.Item = order.Item
		}

		if order.UserId != "" {
			gotOrder.UserId = order.UserId
		}

		gotOrder.UpdatedAt = timeNow

		o.cache[id] = gotOrder
	}

	return
}

// PostOrder setting new order in memory
func (o *OrderMemController) PostOrder(order *models.Order) (err error) {
	o.mut.Lock()
	defer o.mut.Unlock()

	timeNow := time.Now().Nanosecond()

	id := cuid.New()

	order.Id = id
	order.CreatedAt = timeNow
	order.UpdatedAt = timeNow
	o.cache[id] = order

	return
}
