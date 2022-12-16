package main

import (
	"context"
	"os"
	"users/internal/api/controllers"
	"users/internal/api/routers"
	"users/internal/application/services"
	"users/internal/infrastructure/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)

	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	userRepository := repositories.NewMongoUserRepository(client.Database("users"))
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	router := routers.NewRouter(userController)
	router.Run(port)
}
