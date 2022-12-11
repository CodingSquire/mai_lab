package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"log"
	"mai_lab/internal/config"
	"mai_lab/internal/user"
	"net"
	"net/http"
	"time"
)

func main() {

	log.Println("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	storage := user.NewStorage()

	user1 := user.User{
		ID:           uuid.UUID{},
		Name:         "kolya",
		Email:        "kol",
		Mobile:       "8-932",
		PasswordHash: "121",
	}

	user2 := user.User{
		ID:           uuid.UUID{},
		Name:         "kolya",
		Email:        "kol",
		Mobile:       "8-932",
		PasswordHash: "121",
	}

	storage.Create(context.Background(), user1)
	storage.Create(context.Background(), user2)

	log.Println("ID:  ", user1.ID)
	log.Println("ID2:  ", user2.ID)
	log.Println("register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config.Config) {
	log.Println("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		// TODO  socket ?
	} else {
		log.Println("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		log.Printf("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
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
