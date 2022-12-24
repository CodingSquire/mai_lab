package serv

import (
	"context"
	user "mai_lab/app/repository"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server //входящий обращается в логику--incoming turns to logic
	us  *user.Users
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{} //открыть

	s.srv = http.Server{
		Addr:              addr,             //addres-port
		Handler:           h,                //muxer
		ReadTimeout:       30 * time.Second, //иначе бесконечно висящий клиент. Переполнение ресурса
		WriteTimeout:      30 * time.Second, //otherwise an infinitely hanging client. resource overflow
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx := context.Background()
	s.srv.Shutdown(ctx) //выключить сервер от контекста --shutdown server from context
}

func (s *Server) Start(us *user.Users) {
	s.us = us
	go s.srv.ListenAndServe() //без go сразу остановится//паралльное выполнение. Горутина.
	//without go, //parallel execution will immediately stop. Coroutines.
}
