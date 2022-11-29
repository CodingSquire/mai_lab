package main

import (
	"encoding/json"
	"fmt"
	"mai_lab/pkg/service"
)

func main() {
	fmt.Println("Hello world")

	users := service.NewUsers()

	u := service.User{Login: "log", Email: "em", Mobile: "8-800", Name: "name", Password: "pass"}

	js, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	} else {
		users.CreateUser(js)
	}

	u1, err := users.GetUser("log")
	if err != nil {
		fmt.Println(u1.Email)
	} else {
		fmt.Println(err)
	}

	for k, v := range users {
		fmt.Println(k, "value is", v)
	}

}
