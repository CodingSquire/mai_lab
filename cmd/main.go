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
	pgClient, err := postgresql.NewClient(context.Background(), pgConfig)
	if err != nil {
		log.Fatalln(context.Background(), err)
	}

	//	storage := storage.NewMemoryStorage()
	storage := storage.NewPostgreStorage(pgClient)
	userService := services.NewService(storage)
	userHandler := rest.NewHandler(userService)
	userHandler.Register(router)

	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config.Config) {
	log.Println("start application")

	var listener net.Listener
	var listenErr error

	if cfg.HTTP.Type == "sock" {
		// TODO  socket ?
	} else {
		log.Println("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.HTTP.BindIP, cfg.HTTP.Port))
		log.Printf("server is listening port %s:%s", cfg.HTTP.BindIP, cfg.HTTP.Port)
	}

	if listenErr != nil {
		log.Fatalln(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))

}
