package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orders/controllers"
	implmemory "orders/controllers/impl_memory"
	httpApp "orders/http"
	"orders/models"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(httpApp.PARAMS).(map[string]string)
	orderController := r.Context().Value(implmemory.CONTROLLERKEY).(controllers.OrderController)

	order, ok := orderController.GetOrderById(params["id"])

	if ok {
		fmt.Fprintf(w, "Got, %q", *order.Item)
	} else {
		fmt.Fprintf(w, "Failed to get, %q", params["id"])
	}
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(httpApp.PARAMS).(map[string]string)
	orderController := r.Context().Value(implmemory.CONTROLLERKEY).(controllers.OrderController)

	orderController.DeleteOrderById(params["id"])

	fmt.Fprintf(w, "Deleted, %q", params["id"])
}

func PostOrder(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(httpApp.PARAMS).(map[string]string)
	orderController := r.Context().Value(implmemory.CONTROLLERKEY).(controllers.OrderController)

    var order models.OrderPost
    err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Parsed: %+v\n", order)
	orderController.PostOrder(params["id"], order.MakeOrder())
	fmt.Fprintf(w, "Post, %q", params["id"])
}
