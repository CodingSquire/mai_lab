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

	storage := user.NewStorage()
	userService := user.NewService(storage)
	userHandler := user.NewHandler(userService)
	userHandler.Register(router)

	//u1 := user.CreateUserDTO{
	//	Name:     "kolya",
	//	Email:    "avr.nic",
	//	Mobile:   "1593",
	//	Password: "1332",
	//}
	//
	//u2 := user.CreateUserDTO{
	//	Name:     "kolya",
	//	Email:    "avr.nikves@gmail",
	//	Mobile:   "159332",
	//	Password: "133fefewf2",
	//}
	//
	//userService.CreateUser(context.Background(), u1)
	//userService.CreateUser(context.Background(), u2)

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
