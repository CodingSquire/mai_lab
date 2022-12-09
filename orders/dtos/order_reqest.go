package dtos

import "orders/models"

type OrderPost struct {
	Item   *string `json:"item"`
	UserId *string `json:"user_id"`
}

func (o *OrderPost) MakeOrder() *models.Order {
	order := &models.Order{}

	if o.Item != nil {
		order.Item = *o.Item
	}

	if o.UserId != nil {
		order.UserId = *o.UserId
	}

	return order
}
