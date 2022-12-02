package main

import (
	"encoding/json"
	"fmt"
	"mai_lab/pkg/service"
)

func main() {
	fmt.Println("Hello world")

	u := service.User{Login: "log", Email: "em", Mobile: "8-800", Name: "name", Password: "pass"}
	u2 := service.User{Login: "Sixzer", Email: "6zer", Mobile: "8-999-999-99-99", Name: "Alex", Password: "qwe"}

	existingUser := service.User{Login: "log", Mobile: "8-800", Name: "name", Password: "pass"}

	js, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	} else {
		service.CreateUserFromJSON(js)
	}

	service.CreateUser(u2)

	service.CreateUser(existingUser)

	service.PrintUsers()

	u1, err2 := service.GetUser("log")
	if err2 == nil {
		fmt.Println(u1)
	} else {
		fmt.Println(err2)
	}

}
