package service

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

const salt = "dfvjfewfewdvndojpbe"

//type passwordEnconder interface {
//	generatePasswordHash(password string) string
//}

type key string

type User struct {
	Login    key    `json:"login"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile" , omitempty`
	Name     string `json:"name"`
	Password string `json:"password"`
	//BirthDay date   `json:"birth_day", omitempty`
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type usersMap map[key]User

var users = usersMap{}

// Можно добавить проверку по email
func (u usersMap) checkUserID(ID key) bool {
	if _, ok := u[ID]; ok {
		fmt.Println("User ", ID, " exist")
		return true
	} else {
		fmt.Println("User not found")
		return false
	}
}

func (u *User) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(u)
	if err != nil {
		return err
	}
	_, err = w.Write(js)
	return err
}

//func ReadJSON(w io.Writer) error {
//	var u User
//	err := json.Unmarshal(w, u)
//	if err != nil {
//		return err
//	}
//	_, err = w.Write(js)
//	return err
//}

func CreateUser(b []byte) bool {
	u := User{}
	err := json.Unmarshal(b, &u)
	if err == nil && !users.checkUserID(u.Login) {
		u.Password = generatePasswordHash(u.Password)
		users[u.Login] = u
		return true
	}
	return false
}

func GetUser(login key) (User, error) {
	if exist := users.checkUserID(login); exist {
		return users[login], nil
	} else {
		return User{}, errors.New("the User does not exist")
	}
}

//func NewUsers() usersMap {
//	return make(usersMap)
//}

//func main() {
//
//	//us := make(usersMap)
//
//	//	us.CreateUser()
//	//	us.CreateUser(11)
//	//	us.CreateUser(10)
//}
