package controllers

import (
	"encoding/json"
	"net/http"
	"users/internal/api/common"
	"users/rpc/orders"

	"github.com/google/uuid"
)

type OrderController interface {
	GetAllOrdersByUserId(w http.ResponseWriter, r *http.Request)
	CreateOrderByUserId(w http.ResponseWriter, r *http.Request)
	UpdateOrderByUserId(w http.ResponseWriter, r *http.Request)
	DeleteOrderByUserId(w http.ResponseWriter, r *http.Request)
}

type orderController struct {
	ordersClient orders.Orders
}

// CreateOrderByUserId implements OrderController
func (c *orderController) CreateOrderByUserId(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	id := r.Context().Value(common.ContextKeyParams).(map[string]string)["id"]
	userId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var order orders.Order
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orderRequest := orders.CreateOrderRequest{
		UserId: userId.String(),
		Item:   order.Item,
		Adress: order.Adress,
		Count:  order.Count,
	}
	newOrder, err := c.ordersClient.CreateOrder(r.Context(), &orderRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}

// DeleteOrderByUserId implements OrderController
func (*orderController) DeleteOrderByUserId(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetAllOrdersByUserId implements OrderController
func (*orderController) GetAllOrdersByUserId(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// UpdateOrderByUserId implements OrderController
func (*orderController) UpdateOrderByUserId(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewOrderController() OrderController {
	ordersClient := orders.NewOrdersProtobufClient("http://localhost:8080", &http.Client{})

	return &orderController{
		ordersClient: ordersClient,
	}
}
