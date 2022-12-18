package user

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
	Phone string
}

type UserStore interface {
	Create(u User) (*uuid.UUID, error)
	Read(uid uuid.UUID) (*User, error)
	Delete(uid uuid.UUID) error
	SearchUsers(s string) (chan User, error)
}

type Users struct {
	ustore UserStore
}

func NewUsers(ustore UserStore) *Users { //внешние запросы
	//external requests
	return &Users{
		ustore: ustore,
	}
}

func (us *Users) Create(u User) (*User, error) {
	id, err := us.ustore.Create(u)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}
	u.ID = *id // копия структуры - копия id
	//structure copy - id copy
	return &u, nil
}

func (us *Users) Read(uid uuid.UUID) (*User, error) {
	u, err := us.ustore.Read(uid)
	if err != nil {
		return nil, fmt.Errorf("read user error: %w", err)
	}
	return u, nil
}

func (us *Users) Delete(uid uuid.UUID) (*User, error) { // удаление по uuid
	//delete by uuid
	u, err := us.ustore.Read(uid) // сначала читаем по uuid
	//first read by uuid
	if err != nil {
		return nil, fmt.Errorf("search user error: %w", err)
	}
	return u, us.ustore.Delete(uid) //ошибка прочитанная из стора и ошибка из стора.
	//an error read from the store and an error from the store.
}

func (us *Users) SearchUsers(s string) (chan User, error) { //возврат канала с юзерами
	//return of the channel with users
	chin, err := us.ustore.SearchUsers(s)
	if err != nil {
		return nil, err
	}
	chout := make(chan User, 100) //буфер 100 человек. читатеть
	//buffer 100 people. read
	go func() {
		defer close(chout) // закрыть читающий канал
		//close reading channel
		for {
			select {
			case u, ok := <-chin: //читаем из входящего канала
				//read from incoming channel
				if !ok { //если канал закрылся - выходим
					//if the channel is closed - exit
					return
				}
				chout <- u
			}
		}
	}()
	return chout, nil
}
