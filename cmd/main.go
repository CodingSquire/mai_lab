package main

import (
	"fmt"
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

	} else {
		log.Println("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenErr != nil {
		log.Fatalln(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	log.Fatalln(server.Serve(listener))

}
