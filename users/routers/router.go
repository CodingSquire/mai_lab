package routers

import (
	"net/http"
	"users/controllers"
	"users/repositories"
	"users/services"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) SetupRoutes() {
	userRepository := repositories.NewInMemoryUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	http.HandleFunc("/users", userController.GetAllUsers)
	http.HandleFunc("/users/create", userController.CreateUser)
}

func (r *Router) Run(port string) {
	http.ListenAndServe(":"+port, nil)
}
