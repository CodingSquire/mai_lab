package main

import (
	"log"
	"github.com/innerave/mai_lab/orders/internal/api"
	"github.com/innerave/mai_lab/orders/internal/controllers"
	"github.com/innerave/mai_lab/orders/internal/db"
	"github.com/innerave/mai_lab/orders/internal/dotenv"
	"github.com/innerave/mai_lab/orders/internal/http"
	"github.com/innerave/mai_lab/orders/rpc/orders"
	"os"
)

func main() {
	dotenv.Config()

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

	log.Fatal(app.Run(os.Getenv("PORT")))
}
