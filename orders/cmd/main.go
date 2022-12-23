package main

import (
	"log"
	"orders/internal/api"
	"orders/internal/controllers"
	"orders/internal/db"
	"orders/internal/dotenv"
	"orders/internal/http"
	"orders/rpc/orders"
)

func main() {
	config := dotenv.Config()

	db, err := db.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	app := http.NewApp()
	app.Use(api.LoggerMiddleware)

	orderController := controllers.NewPgxController(db)

	serverImpl := api.NewTwirpServer(orderController)
	twirpHandler := orders.NewOrdersServer(serverImpl)
	app.Handle(twirpHandler.PathPrefix(), twirpHandler)

	// orderController := controllers.NewOrderMemController()
	ordersApi := api.NewOrdersApi(orderController)
	ordersApi.SetRoutes(app)

	log.Fatal(app.Run(config.Get("PORT")))
}
