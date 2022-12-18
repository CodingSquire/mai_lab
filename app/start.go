package start

import (
	"context"
	user "mai_lab/app/repository"
	"sync"
)

type App struct {
	us *user.Users
}

func NewApp(ust user.UserStore) *App { // init app, back billed app
	a := &App{
		us: user.NewUsers(ust),
	}
	return a
}

type HTTPServer interface { //start - внешняя приложуха. Не должна знать об api(внешний адаптер)
	//external attachment. Must not know about api(external adapter)
	Start(us *user.Users)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs HTTPServer) {
	defer wg.Done()
	hs.Start(a.us)
	<-ctx.Done()
	hs.Stop()
}
