package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"mai_lab/internal/config"
	"mai_lab/internal/controller/rest"
	"mai_lab/internal/infrastructure/storage"
	"mai_lab/internal/services"
	"mai_lab/pkg/client/postgresql"
	"mai_lab/rpc"

	"github.com/julienschmidt/httprouter"
)

func main() {

	log.Println("create router")
	router := httprouter.New()

	cfg := config.GetConfig(".")
	pgConfig := postgresql.NewPgConfig(
		cfg.Username, cfg.Password,
		cfg.Host, cfg.Portdb, cfg.Database,
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

	twiprService := services.NewTwirpService(userStorage)
	rpc.NewUsersServer(twiprService)

	userService := services.NewService(userStorage)
	userHandler := rest.NewHandler(userService)
	userHandler.Register(router)

	startHTTP(router, cfg)

}

func startHTTP(router *httprouter.Router, cfg *config.Config) {
	log.Println("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.BindIP, cfg.Port))
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("server is listening port %s:%s", cfg.BindIP, cfg.Port)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err = server.Serve(listener); err != nil {
		log.Fatalln(err)
	}

}
