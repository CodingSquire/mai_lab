package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"mai_lab/internal/config"
	"mai_lab/internal/infrastructure/storage"
	"mai_lab/internal/services"
	"mai_lab/internal/transport/rest"
	"mai_lab/pkg/client/postgresql"

	"github.com/julienschmidt/httprouter"
)

func main() {

	log.Println("create router")
	router := httprouter.New()

	cfg := config.GetConfig()
	pgConfig := postgresql.NewPgConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port, cfg.PostgreSQL.Database,
	)

	// Attempt to connect to the database
	// if it does not connect, then we create a cache storage
	var userStorage storage.Storage
	pgClient, err := postgresql.NewClient(context.Background(), pgConfig)
	if err != nil {
		log.Fatalln(context.Background(), err)
		userStorage = storage.NewMemoryStorage()
	} else {
		userStorage = storage.NewPostgreStorage(pgClient)
	}

	userService := services.NewService(userStorage)
	userHandler := rest.NewHandler(userService)
	userHandler.Register(router)

	startHTTP(router, cfg)

}

func startHTTP(router *httprouter.Router, cfg *config.Config) {
	log.Println("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.HTTP.BindIP, cfg.HTTP.Port))
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("server is listening port %s:%s", cfg.HTTP.BindIP, cfg.HTTP.Port)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err = server.Serve(listener); err != nil {
		log.Fatalln(err)
	}

}
