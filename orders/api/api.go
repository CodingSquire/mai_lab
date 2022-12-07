package api

import (
	"encoding/json"
	"fmt"
	implmemory "orders/controllers/impl_memory"
	"orders/http"
	"orders/models"
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

func DeleteOrder(r *http.RouteContext) {
	var orderController implmemory.OrderMemController
	r.State(&orderController)

	orderController.DeleteOrderById(r.Params("id"))

	r.SendString(fmt.Sprintf("Deleted, %q", r.Params("id")))
}

func PostOrder(r *http.RouteContext) {
	var orderController implmemory.OrderMemController
	r.State(&orderController)

    var order models.OrderPost
    err := json.NewDecoder(r.Body()).Decode(&order)

	if err != nil {
		r.SendError(err)
		return
	}

	//fmt.Printf("Parsed: %+v\n", order)
	orderController.PostOrder(r.Params("id"), order.MakeOrder())
	r.SendString(fmt.Sprintf("Post, %q", r.Params("id")))
}
