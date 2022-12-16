package main

import (
	"os"
	"users/internal/api/controllers"
	"users/internal/api/routers"
	"users/internal/application/services"
	"users/internal/infrastructure/repositories"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	userRepository := repositories.NewInMemoryUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	router := routers.NewRouter(userController)
	router.Run(port)
}
