package dtos

import (
	"orders/internal/models"
	"orders/rpc/orders"
)

// TwirpFromOrder converts a model.Order to a rpc.orders.Order
func TwirpFromOrder(order *models.Order) *orders.Order {
	return &orders.Order{
		Id:     order.ID,
		UserId: order.UserID,
		Item:   order.Item,
		Count:  int64(order.Count),
		Adress: order.Address,
		CreatedAt: order.CreatedAt.Unix(),
		UpdatedAt: order.UpdatedAt.Unix(),
	}
}

// TwirpFromOrders converts a slice of model.Order to a slice of rpc.orders.Order
func TwirpFromOrders(order []models.Order) []*orders.Order {
	var twirpOrders []*orders.Order
	for _, o := range order {
		twirpOrders = append(twirpOrders, TwirpFromOrder(&o))
	}
	return twirpOrders
}
