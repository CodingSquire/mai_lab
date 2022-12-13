package main

import (
	"log"
	"orders/api"
	"orders/controllers"
	"orders/dotenv"
	"orders/http"
)

func main() {
	config := dotenv.Config()

	app := http.NewApp()

	orderController := controllers.NewOrderMemController()
	ordersApi := api.NewOrdersApi(orderController)
	ordersApi.SetRoutes(app)

	// app.Use(api.LoggerMiddleware)

	log.Fatal(app.Run(config.Get("PORT")))
}
