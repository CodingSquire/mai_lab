package implmemory

import "orders/models"

type OrderMemController struct {
	cache map[string]*models.Order
}

func NewOrderMemController() *OrderMemController {
	return &OrderMemController{
		cache: make(map[string]*models.Order),
	}
}

func (o *OrderMemController) GetAllOrders() []models.Order {
	orders := make([]models.Order, len(o.cache))

	for _, order := range orders {
		orders = append(orders, order)
	}

	return orders
}

func (o *OrderMemController) GetOrderById(id string) (*models.Order, bool) {
	order, ok := o.cache[id]
	return order, ok
}

func (o *OrderMemController) DeleteOrderById(id string) {
	delete(o.cache, id)
}

func (o *OrderMemController) PatchOrderById(id string, order *models.Order) {
	order.Id = &id
	o.cache[id] = order
}

func (o *OrderMemController) PostOrder(id string, order *models.Order) {
	order.Id = &id
	o.cache[id] = order
}
