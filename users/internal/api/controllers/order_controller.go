package controllers

import (
	"net/http"
)

type OrderController interface {
	GetOrderById(w http.ResponseWriter, r *http.Request)
	GetAllOrders(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
}
