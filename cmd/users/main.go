package main

import (
	"context"
	user_handler "go_services_lab/pkg/user/handler"
	"go_services_lab/pkg/user/proto"
	user_repository "go_services_lab/pkg/user/repository"
	user_service "go_services_lab/pkg/user/service"
	postgres "go_services_lab/postgres"
	server "go_services_lab/server"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	// cache_user := cache.New(5*time.Minute, 10*time.Minute)
	// cache_user.Set("user1", &models.User{1, "Alexey", "lewka", "lewka007"}, cache.DefaultExpiration)
	// cache_user.Set("user2", &models.User{2, "Ivan", "vane4ka", "trueMan_"}, cache.DefaultExpiration)
	// cache_user.Set("user3", &models.User{3, "Masha", "tyan", "mashanyasha"}, cache.DefaultExpiration)
	// cache_user.Set("countUser", 3, cache.DefaultExpiration)

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     "postgres_container",
		Port:     "5432",
		Username: "postgres",
		Password: "qweasd",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Printf("No accees to database: %s", err.Error())
	}

	repository_user := user_repository.NewRepositoryUser(db)
	service_user := user_service.NewServiceUser(repository_user)
	handler_user := user_handler.NewHandlerUser(service_user)

	go func() {
		s := grpc.NewServer()
		srv := &user_repository.GRPCUsersServer{User: service_user}
		proto.RegisterUsersServer(s, srv)

		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal(err)
		}

		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	server_user := new(server.Server)
	if err := server_user.Run("8000", handler_user.InitRoutesUser()); err != nil {
		log.Printf("listen: %s\n", err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server_user.Shutdown(ctx); err != nil {
		log.Fatal("Server user forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
