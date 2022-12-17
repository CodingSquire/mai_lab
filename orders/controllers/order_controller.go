package controllers

import "orders/models"

type OrderController interface {
	GetAllOrders() []models.Order
	GetAllOrdersByUserId(userId string) []models.Order
	GetOrderById(id string) (*models.Order, bool)
	DeleteOrderById(id string) error
	PatchOrderById(id string, order *models.Order) error
	PostOrder(order *models.Order) error
}
