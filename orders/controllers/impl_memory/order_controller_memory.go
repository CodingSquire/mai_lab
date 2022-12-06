package implmemory

import (
	"orders/models"
)

type OrderMemController struct {
	_memory map[string]*models.Order
}

const CONTROLLERKEY = "orderController"

func NewOrderMemController() *OrderMemController {
	return &OrderMemController{
		_memory: make(map[string]*models.Order),
	}
}

func (o *OrderMemController) GetAllOrders() []models.Order {
	orders := make([]models.Order, len(o._memory))

	for _, order := range orders {
		orders = append(orders, order)
	}

	return orders
}

func (o *OrderMemController) GetOrderById(id string) (*models.Order, bool) {
	order, ok := o._memory[id]
	return order, ok
}

func (o *OrderMemController) DeleteOrderById(id string) {
	delete(o._memory, id)
}

func (o *OrderMemController) PatchOrderById(id string, order *models.Order) {
	order.Id = &id
	o._memory[id] = order
}

func (o *OrderMemController) PostOrder(id string, order *models.Order) {
	order.Id = &id
	o._memory[id] = order
}
