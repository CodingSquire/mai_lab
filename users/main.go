package main

import (
	"os"
	"users/controllers"
	"users/repositories"
	"users/routers"
	"users/services"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	userRepository := repositories.NewInMemoryUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	router := routers.NewRouter(userController)
	router.Run(port)
}
