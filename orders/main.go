package main

import (
	"log"
	"orders/api"
	"orders/dotenv"
	"orders/http"
	"orders/controllers/impl_memory"
)

func main() {
	config := dotenv.Config()

	app := http.NewApp()

	app.Manage(implmemory.CONTROLLERKEY, implmemory.NewOrderMemController())

	app.Get("/order/:id", api.GetOrder)
	app.Delete("/order/:id", api.DeleteOrder)
	app.Post("/order/:id", api.PostOrder)

	// app.Use(api.LoggerMiddleware)

	log.Fatal(app.Run(config.Get("PORT")))
}
