package main

import (
	"context"
	"log"
	"os"
	"time"
	"users/internal/api/controllers"
	"users/internal/api/routers"
	"users/internal/application/services"
	"users/internal/infrastructure/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := os.Getenv("PORT_MAIN")
	if port == "" {
		port = "5050"
	}
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGO_URI")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repositories.NewMongoUserRepository(client.Database("users"))
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	router := routers.NewRouter(userController)
	router.Run(port)
}
