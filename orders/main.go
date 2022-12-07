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

	order.Get("/:id", api.GetOrder)
	order.Delete("/:id", api.DeleteOrder)
	order.Post("/:id", api.PostOrder)

	// app.Use(api.LoggerMiddleware)

	log.Fatal(app.Run(config.Get("PORT")))
}
