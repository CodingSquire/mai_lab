package dtos

import "orders/internal/models"

type OrderPost struct {
	ID      *string `json:"id"`
	Item    *string `json:"item"`
	UserId  *string `json:"user_id"`
	Address *string `json:"address"`
	Count   *int    `json:"count"`
}

func (o *OrderPost) MakeOrder() *models.Order {
	order := &models.Order{}


	if o.ID != nil {
		order.ID = *o.ID
	}

	if o.Item != nil {
		order.Item = *o.Item
	}

	if o.UserId != nil {
		order.UserID = *o.UserId
	}

	if o.Address != nil {
		order.Address = *o.Address
	}

	if o.Count != nil {
		order.Count = *o.Count
	}

	return order
}
