package main

import (
	"context"
	"log"
	"os"
	"time"
	"users/internal/api/controllers"
	"users/internal/api/routers"
	"users/internal/application/services"
	"users/internal/contracts"
	"users/internal/infrastructure/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := getPort()
	var repository contracts.UserRepository
	if os.Getenv("MONGO_URI") != "" {
		client := getMongoClient()
		repository = repositories.NewMongoUserRepository(client.Database("users"))
	} else {
		repository = repositories.NewInMemoryUserRepository()
		log.Println("MONGO_URI not set, falling back to in-memory repository")
	}
	userService := services.NewUserService(repository)
	userController := controllers.NewUserController(userService)
	router := routers.NewRouter(userController)
	router.Run(port)
}

func getMongoClient() *mongo.Client {
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
	return client
}

func getPort() string {
	port := os.Getenv("PORT_MAIN")
	if port == "" {
		port = "5050"
	}
	return port
}
