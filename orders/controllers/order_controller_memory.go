package controllers

import (
	"orders/models"
	"strconv"
	"sync"
	"time"
)

type OrderMemController struct {
	cache map[string]*models.Order
	mut   sync.Mutex
}

var globalIndex int

func NewOrderMemController() OrderController {
	return &OrderMemController{
		cache: make(map[string]*models.Order),
		mut:   sync.Mutex{},
	}
}

func (o *OrderMemController) GetAllOrders() []models.Order {
	o.mut.Lock()
	defer o.mut.Unlock()

	orders := make([]models.Order, 0)

	for _, order := range o.cache {
		orders = append(orders, *order)
	}

	return orders
}

func (o *OrderMemController) GetOrderById(id string) (*models.Order, bool) {
	o.mut.Lock()
	defer o.mut.Unlock()

	order, ok := o.cache[id]
	return order, ok
}

func (o *OrderMemController) DeleteOrderById(id string) {
	o.mut.Lock()
	defer o.mut.Unlock()

	delete(o.cache, id)
}

func (o *OrderMemController) PatchOrderById(id string, order *models.Order) {
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
}

func (o *OrderMemController) PostOrder(order *models.Order) {
	o.mut.Lock()
	defer o.mut.Unlock()

	timeNow := time.Now().Nanosecond()

	globalIndex++
	id := strconv.Itoa(globalIndex)

	order.Id = id
	order.CreatedAt = timeNow
	order.UpdatedAt = timeNow
	o.cache[id] = order
}
