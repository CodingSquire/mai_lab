package store_users

import (
	"database/sql"
	"github.com/google/uuid"
	user "mai_lab/app/repository"
	"strings"
	"time"
)

var _ user.UserStore = &Users{} //type matching user.syntax check.

type Users struct { // хранилище с юзерами.
	//repository with users.
	m map[uuid.UUID]user.User
}

func NewUsers() *Users {
	return &Users{
		m: make(map[uuid.UUID]user.User), // initialization map
	}
}

func (us *Users) Create(u user.User) (*uuid.UUID, error) {

	uid := uuid.New()
	u.ID = uid       //присваивание uuid--uuid assignment
	us.m[u.ID] = u   //запомним uuid--remember uuid
	return &uid, nil //вернем uuid--return uuid
}

// implementation of what we expect from users
func (us *Users) Read(uid uuid.UUID) (*user.User, error) {

	u, ok := us.m[uid] //получить user из map--get user from map
	if ok {
		return &u, nil //если получили его--if you received it
	}
	return nil, sql.ErrNoRows
}

func (us *Users) Delete(uid uuid.UUID) error {

	delete(us.m, uid)
	return nil //не возвращает ошибку если не нашли--does not return an error if not found
}

func (us *Users) SearchUsers(s string) (chan user.User, error) {

	chout := make(chan user.User, 100) //канал создаем--create a channel

	go func() {
		defer close(chout)
		for _, u := range us.m { //перебор map--enumeration map
			if strings.Contains(u.Name, s) { //если есть юзер--if there is a user
				select {
				case <-time.After(2 * time.Second):
					return
				case chout <- u: //то отправляем юзера в наш канал--then we send the user to our channel
				}
			}
		}
	}()

	return chout, nil
}
