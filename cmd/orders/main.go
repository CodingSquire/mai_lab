package main

import (
	"context"
	order_handler "go_services_lab/pkg/order/handler"
	order_repository "go_services_lab/pkg/order/repository"
	order_service "go_services_lab/pkg/order/service"
	postgres "go_services_lab/postgres"
	server "go_services_lab/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// cache_order := cache.New(5*time.Minute, 10*time.Minute)
	// cache_order.Set("tata", 12, cache.DefaultExpiration)
	// cache_order.Set("product1", &models.Product{1, "Banana", 12.}, cache.DefaultExpiration)
	// cache_order.Set("product2", &models.Product{2, "Apple", 16.}, cache.DefaultExpiration)
	// cache_order.Set("product3", &models.Product{3, "Orange", 20.}, cache.DefaultExpiration)
	// cache_order.Set("countProduct", 3, cache.DefaultExpiration)
	// cache_order.Set("order1", &models.Order{1, 1, models.Stores{{models.Product{1, "Banana", 12.}, 10}, {models.Product{2, "Apple", 16.}, 15}}}, cache.DefaultExpiration)
	// cache_order.Set("order2", &models.Order{2, 2, models.Stores{{models.Product{1, "Banana", 12.}, 2}, {models.Product{2, "Apple", 16.}, 10}, {models.Product{3, "Orange", 20.}, 25}}}, cache.DefaultExpiration)
	// cache_order.Set("countOrder", 2, cache.DefaultExpiration)
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     "postgres_container",
		Port:     "5432",
		Username: "postgres",
		Password: "qweasd",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Printf("No accees to database: %s", err.Error())
	}

	repository_order := order_repository.NewRepositoryOrder(db)
	service_order := order_service.NewServiceOrder(repository_order)
	handler_order := order_handler.NewHandlerOrder(service_order)

	server_order := new(server.Server)
	if err := server_order.Run("8000", handler_order.InitRoutesOrder()); err != nil {
		log.Printf("listen: %s\n", err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server_order.Shutdown(ctx); err != nil {
		log.Fatal("Server order forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
