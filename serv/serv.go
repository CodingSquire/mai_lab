package serv

import (
	"context"
	user "mai_lab/app/repository"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server //server embedding
	us  *user.Users
}

func NewServer(addr string, h http.Handler) *Server { //business logic outside
	s := &Server{}

	s.srv = http.Server{
		Addr:              addr, //==ListenAndServe()
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second, // or there will be a constantly hanging client
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) //add 2 sec
	s.srv.Shutdown(ctx)
	cancel()
}

func (s *Server) Start(us *user.Users) {
	s.us = us
	go s.srv.ListenAndServe() //ListenAndServeTLS() -certification
}
