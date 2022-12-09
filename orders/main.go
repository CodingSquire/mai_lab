package main

import (
	"log"
	"orders/api"
	"orders/controllers/impl_memory"
	"orders/dotenv"
	"orders/http"
)

func main() {
	config := dotenv.Config()

	app := http.NewApp()

	app.Manage(implmemory.NewOrderMemController())

	order := app.Group("/order")

	order.Use(api.LoggerMiddleware)

	order.Get("/:id", api.GetOrder)
	order.Get("", api.GetAllOrders)
	order.Delete("/:id", api.DeleteOrder)
	order.Post("", api.PostOrder)
	order.Patch("/:id", api.UpdateOrder)

	// app.Use(api.LoggerMiddleware)

	log.Fatal(app.Run(config.Get("PORT")))
}
