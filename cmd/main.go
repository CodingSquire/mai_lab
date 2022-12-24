package main

import (
	"context"
	_ "github.com/lib/pq"
	"log"
	handler "mai_lab/api"
	start "mai_lab/app"
	user "mai_lab/app/repository"
	store_users "mai_lab/db"
	postgres "mai_lab/postgres"
	serv "mai_lab/serv"
	"sync"
)

func main() {
	ctx := context.Background() //глобальный стартовый контекст.--global start context.

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     "postgres_container",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Printf("No accees to database: %s", err.Error())
	}

	ust := store_users.NewUsers(db)
	a := start.NewApp(ust)
	us := user.NewUsers(ust)
	h := handler.NewRouter(us)
	srv := serv.NewServer(":8000", h)

	wg := &sync.WaitGroup{} //дожидаться всех горутин--wait for all goroutines
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
}
