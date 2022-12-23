package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"users/internal/api/common"
	"users/rpc/orders"

	"github.com/google/uuid"
)

type OrderController interface {
	GetAllOrdersByUserId(w http.ResponseWriter, r *http.Request)
	GetOrderById(w http.ResponseWriter, r *http.Request)
	CreateOrderByUserId(w http.ResponseWriter, r *http.Request)
	GetAllOrders(w http.ResponseWriter, r *http.Request)
	UpdateOrderById(w http.ResponseWriter, r *http.Request)
	DeleteOrderById(w http.ResponseWriter, r *http.Request)
}

type orderController struct {
	ordersClient orders.Orders
}

// DeleteOrderById gets the order id from the context and deletes the order.
func (c *orderController) DeleteOrderById(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	orderId := r.Context().Value(common.ContextKeyParams).(map[string]string)["orderId"]
	orderIdParsed, err := uuid.Parse(orderId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = c.ordersClient.DeleteOrder(r.Context(), &orders.DeleteOrderRequest{
		Id: orderIdParsed.String(),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAllOrders gets all orders from the orders service.
func (c *orderController) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	orders, err := c.ordersClient.GetAllOrders(r.Context(), &orders.GetAllOrdersRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

// UpdateOrderById updates an order by id.
func (c *orderController) UpdateOrderById(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	orderId := r.Context().Value(common.ContextKeyParams).(map[string]string)["orderId"]
	orderIdParsed, err := uuid.Parse(orderId)
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

	order.Id = orderIdParsed.String()
	updatedOrder, err := c.ordersClient.UpdateOrder(r.Context(), &orders.UpdateOrderRequest{
		Order: &order,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedOrder)
}

// GetOrderByUserId gets order by id.
func (c *orderController) GetOrderById(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	orderId := r.Context().Value(common.ContextKeyParams).(map[string]string)["orderId"]
	orderIdParsed, err := uuid.Parse(orderId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := c.ordersClient.GetOrder(r.Context(), &orders.GetOrderRequest{
		Id: orderIdParsed.String(),
	})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
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

// GetAllOrdersByUserId implements OrderController
func (c *orderController) GetAllOrdersByUserId(w http.ResponseWriter, r *http.Request) {
	prepareResponse(w, r)
	id := r.Context().Value(common.ContextKeyParams).(map[string]string)["id"]
	userId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orders, err := c.ordersClient.GetAllOrdersByUserId(r.Context(), &orders.GetAllOrdersByUserIdRequest{
		UserId: userId.String(),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func NewOrderController() OrderController {
	ordersClient := orders.NewOrdersProtobufClient(os.Getenv("ORDERS_URL"), &http.Client{})

	return &orderController{
		ordersClient: ordersClient,
	}
}
