package main

import (
	"context"
	handler "mai_lab/api"
	start "mai_lab/app"
	user "mai_lab/app/repository"
	store_users "mai_lab/db"
	serv "mai_lab/serv"
)

func main() {
	ctx := context.Background()

	ust := store_users.NewUsers()
	a := start.NewApp(ust)
	us := user.NewUsers(ust)
	h := handler.NewRouter(us)
	srv := serv.NewServer(":8000", h) //http.DefaultServeMux

	go a.Serve(ctx, srv)

	<-ctx.Done()
}
