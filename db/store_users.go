package store_users

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	user "mai_lab/app/repository"
)

type Users struct { // хранилище с юзерами.
	//repository with users.
	m  map[uuid.UUID]user.User
	db *sqlx.DB //прикрутили бд
}

func NewUsers(db *sqlx.DB) *Users {
	return &Users{
		m:  make(map[uuid.UUID]user.User), // initialization map
		db: db,
	}
}

func (us *Users) Create(u user.User) (*uuid.UUID, error) {

	// uid := uuid.New()
	// u.ID = uid       //присваивание uuid--uuid assignment
	// us.m[u.ID] = u   //запомним uuid--remember uuid
	// return &uid, nil //вернем uuid--return uuid
	uid := uuid.New()

	query := fmt.Sprintf("INSERT INTO users (id, name, email, phone) VALUES ($1, $2, $3, $4) RETURNING id")
	row := us.db.QueryRow(query, uid, u.Name, u.Email, u.Phone)

	if err := row.Scan(&uid); err != nil {
		return &uuid.Nil, err
	}

	u.ID = uid     //присваивание uuid--uuid assignment
	us.m[u.ID] = u //запомним uuid--remember uuid

	log.Print(us.m)
	return &uid, nil
}

// implementation of what we expect from users
func (us *Users) Read(uid uuid.UUID) (*user.User, error) {
	var user user.User

	user, ok := us.m[uid] //получить user из map--get user from map
	if ok {
		return &user, nil //если получили его--if you received it
	}

	err := us.db.Get(&user, "SELECT * FROM users WHERE id=$1", uid)
	us.m[user.ID] = user //запомним uuid--remember uuid

	log.Print(us.m)
	return &user, err
}

func (us *Users) Delete(uid uuid.UUID) error {
	// return nil //не возвращает ошибку если не нашли--does not return an error if not found

	res, err := us.db.Exec("DELETE FROM users WHERE id=$1", uid)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count > 0 {
				delete(us.m, uid)
			} else {
				return errors.New("User with this uuid does not exist")
			}
		}

	}

	log.Print(us.m)

	return err
}

func (us *Users) SearchUsers(s string) (chan user.User, error) {

	chout := make(chan user.User, 100) //канал создаем--create a channel

	go func() {
		defer close(chout)
		// for _, u := range us.m { //перебор map--enumeration map
		// 	if strings.Contains(u.Name, s) { //если есть юзер--if there is a user
		// 		select {
		// 		case <-time.After(2 * time.Second):
		// 			return
		// 		case chout <- u: //то отправляем юзера в наш канал--then we send the user to our channel
		// 		}
		// 	}
		// }

		// if len(chout) == 0 {
		var user user.User

		rows, err := us.db.Queryx("SELECT * FROM users WHERE name ILIKE '%' || $1 || '%'", s)

		if err == nil {
			for rows.Next() {
				err = rows.StructScan(&user)
				if err == nil {
					us.m[user.ID] = user
					chout <- user
				}
			}
		}
		// }
	}()

	log.Print(us.m)
	return chout, nil
}
