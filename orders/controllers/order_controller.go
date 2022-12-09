package controllers

import "orders/models"

type OrderController interface {
	GetAllOrders() []models.Order
	GetOrderById(id string) (*models.Order, bool)
	DeleteOrderById(id string)
	PatchOrderById(id string, order *models.Order)
	PostOrder(order *models.Order)
}
