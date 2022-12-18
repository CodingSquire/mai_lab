package main

import (
	"context"
	handler "mai_lab/api"
	start "mai_lab/app"
	user "mai_lab/app/repository"
	store_users "mai_lab/db"
	serv "mai_lab/serv"
	"sync"
)

func main() {
	ctx := context.Background() //глобальный стартовый контекст.--global start context.

	ust := store_users.NewUsers()
	a := start.NewApp(ust)
	us := user.NewUsers(ust)
	h := handler.NewRouter(us)
	srv := serv.NewServer(":8000", h)

	wg := &sync.WaitGroup{} //дожидаться всех горутин--wait for all goroutines
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
}
