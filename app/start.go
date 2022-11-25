package start

import (
	"context"
	user "mai_lab/app/repository"
)

type App struct {
	us *user.Users
}

func NewApp(ust user.UserStore) *App {
	a := &App{
		us: user.NewUsers(ust),
	}
	return a
}

type HTTPServer interface {
	Start(us *user.Users)
	Stop()
}

func (a *App) Serve(ctx context.Context, hs HTTPServer) {
	//	defer wg.Done()
	hs.Start(a.us)
	<-ctx.Done() //подождать.
	hs.Stop()
}
