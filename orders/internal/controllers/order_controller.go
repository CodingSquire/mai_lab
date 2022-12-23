package controllers

import "github.com/innerave/mai_lab/orders/internal/models"

type OrderController interface {
	GetAllOrders() []models.Order
	GetAllOrdersByUserId(userId string) []models.Order
	GetOrderById(id string) (*models.Order, error)
	DeleteOrderById(id string) error
	PatchOrderById(id string, order *models.Order) error
	PostOrder(order *models.Order) (*models.Order, error)
}
