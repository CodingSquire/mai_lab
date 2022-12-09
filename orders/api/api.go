package api

import (
	"fmt"
	implmemory "orders/controllers/impl_memory"
	"orders/dtos"
	"orders/http"
)

func GetOrder(r *http.RouteContext) {
	var orderController implmemory.OrderMemController
	r.State(&orderController)

	order, ok := orderController.GetOrderById(r.Params("id"))

	if ok {
		r.SendJSON(order)
	} else {
		r.SendString(fmt.Sprintf("Failed to get, %q", r.Params("id")))
	}
}

func GetAllOrders(r *http.RouteContext) {
	var orderController implmemory.OrderMemController
	r.State(&orderController)

	orders := orderController.GetAllOrders()

	r.SendJSON(orders)
}

func UpdateOrder(r *http.RouteContext) {
	var orderController implmemory.OrderMemController
	r.State(&orderController)

	var order dtos.OrderPost
	err := r.DecodeJSON(&order)

	if err != nil {
		r.SendError(err)
		return
	}

	orderController.PatchOrderById(r.Params("id"), order.MakeOrder())

	r.SendString(fmt.Sprint("Updated"))
}

func DeleteOrder(r *http.RouteContext) {
	var orderController implmemory.OrderMemController
	r.State(&orderController)

	orderController.DeleteOrderById(r.Params("id"))

	r.SendString(fmt.Sprintf("Deleted, %q", r.Params("id")))
}

func PostOrder(r *http.RouteContext) {
	var orderController implmemory.OrderMemController
	r.State(&orderController)

	var order dtos.OrderPost
	err := r.DecodeJSON(&order)

	if err != nil {
		r.SendError(err)
		return
	}

	//fmt.Printf("Parsed: %+v\n", order)
	orderController.PostOrder(order.MakeOrder())
	r.SendString(fmt.Sprint("Posted"))
}
