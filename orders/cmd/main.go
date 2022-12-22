package main

import (
	"log"
	"orders/internal/api"
	"orders/internal/controllers"
	"orders/internal/db"
	"orders/internal/dotenv"
	"orders/internal/http"
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

	// orderController := controllers.NewOrderMemController()
	orderController := controllers.NewPgxController(db)
	ordersApi := api.NewOrdersApi(orderController)
	ordersApi.SetRoutes(app)

	// app.Use(api.LoggerMiddleware)

	log.Fatal(app.Run(config.Get("PORT")))
}
