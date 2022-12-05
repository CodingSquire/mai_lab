package main

import (
	"os"
	"users/routers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := routers.NewRouter()
	router.SetupRoutes()
	router.Run(port)
}
