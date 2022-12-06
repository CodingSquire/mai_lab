package models

import "time"

type Order struct {
	Id        *string
	Item      *string
	CreatedAt *int
	UpdatedAt *int
}

type OrderPost struct {
	Item string `json:"item"`
}

func (o *OrderPost) MakeOrder() *Order {
	id := "1"
	timeNow := time.Now().Nanosecond()

	return &Order {
		Id: &id,
		Item: &o.Item,
		CreatedAt: &timeNow,
		UpdatedAt: &timeNow,
	}
}
