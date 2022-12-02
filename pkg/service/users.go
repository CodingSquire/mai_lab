package service

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
)

// string for password hashing
const salt = "dfvjfewfewdvndojpbe"

// an alias that can be used anywhere in place of a  type name
type key string

type User struct {
	Login    key    `json:"login"`
	Email    string `json:"email" ,omitempty`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	Password string `json:"password"`
	//BirthDay date   `json:"birth_day", omitempty`
}

// function for password encryption
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

// an alias that can be used anywhere in place of a type name
type usersMap map[key]User

// structure for storing users
var users = make(usersMap)

// the function of verifying the existence of a user with this key
func (u usersMap) checkUserID(keyValue key) bool {
	if _, ok := u[keyValue]; ok {
		fmt.Println("User ", keyValue, " exist")
		return true
	}
	return false
}

// GetUser the function returns the user by key
func GetUser(keyValue key) (User, error) {
	if exist := users.checkUserID(keyValue); exist {
		return users[keyValue], nil
	} else {
		return User{}, errors.New("the User does not exist")
	}
}

// CreateUser the function adds a new user
func CreateUser(u User) bool {
	if !users.checkUserID(u.Login) {
		u.Password = generatePasswordHash(u.Password)
		users[u.Login] = u
		return true
	}
	return false
}

// CreateUserFromJSON CreateUser the function adds a new user from JSON
func CreateUserFromJSON(b []byte) bool {
	u := User{}
	err := json.Unmarshal(b, &u)
	if err == nil {
		return CreateUser(u)
	}
	return false
}

// PrintUsers the function displays information about users
func PrintUsers() {
	fmt.Println("Print users:")
	for _, value := range users {
		fmt.Println(value)
	}
}
