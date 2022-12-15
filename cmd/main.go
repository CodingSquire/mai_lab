package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"mai_lab/internal/config"
	"mai_lab/internal/user"
	"mai_lab/internal/user/storage"
	"mai_lab/pkg/client/postgresql"
	"net"
	"net/http"
	"time"
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

	//	storage := storage.NewStorage()
	storage := storage.NewPostgreStorage(pgClient)
	userService := user.NewService(storage)
	userHandler := user.NewHandler(userService)
	userHandler.Register(router)

	//	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	//	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

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
