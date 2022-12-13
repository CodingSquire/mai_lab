package api

import (
	"fmt"
	"orders/controllers"
	"orders/dtos"
	"orders/http"
)

type OrdersApi struct {
	OrderController controllers.OrderController
}

// NewOrdersApi create service-like api object
func NewOrdersApi(orderController controllers.OrderController) *OrdersApi {	
	return &OrdersApi{
		OrderController: orderController,
	}
}

// SetRoutes inject OrdersApi routes
func (a *OrdersApi) SetRoutes(r http.HttpRouter) {
	order := r.Group("/order")

	//order.Use(LoggerMiddleware)
	
	order.Get("/:id", a.GetOrder)
	order.Get("", a.GetAllOrders)
	order.Delete("/:id", a.DeleteOrder)
	order.Post("", a.PostOrder)
	order.Patch("/:id", a.UpdateOrder)
}

// GetOrder returns order object if exists
func (a *OrdersApi) GetOrder(r *http.RouteContext) {
	order, ok := a.OrderController.GetOrderById(r.Params("id"))

	if ok {
		r.SendJSON(order)
	} else {
		r.SendString(fmt.Sprintf("Failed to get, %q", r.Params("id")))
	}
}

// GetAllOrders returns array of saved orders
func (a *OrdersApi) GetAllOrders(r *http.RouteContext) {
	orders := a.OrderController.GetAllOrders()

	r.SendJSON(orders)
}

// UpdateOrder setting new order params on order with :id
func (a *OrdersApi) UpdateOrder(r *http.RouteContext) {
	var order dtos.OrderPost
	err := r.DecodeJSON(&order)

	if err != nil {
		r.SendError(err)
		return
	}

	a.OrderController.PatchOrderById(r.Params("id"), order.MakeOrder())

	r.SendString(fmt.Sprint("Updated"))
}

// DeleteOrder deletes order with :id
func (a *OrdersApi) DeleteOrder(r *http.RouteContext) {
	a.OrderController.DeleteOrderById(r.Params("id"))

	r.SendString(fmt.Sprintf("Deleted, %q", r.Params("id")))
}

// PostOrder settings new order
func (a *OrdersApi) PostOrder(r *http.RouteContext) {
	var order dtos.OrderPost
	err := r.DecodeJSON(&order)

	if err != nil {
		r.SendError(err)
		return
	}

	//fmt.Printf("Parsed: %+v\n", order)
	a.OrderController.PostOrder(order.MakeOrder())
	r.SendString(fmt.Sprint("Posted"))
}
