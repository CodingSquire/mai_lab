package main

import (
	"encoding/json"
	"fmt"
	"mai_lab/pkg/service"
)

func main() {
	fmt.Println("Hello world")

	u := service.User{Login: "log", Email: "em", Mobile: "8-800", Name: "name", Password: "pass"}

	js, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	} else {
		service.CreateUser(js)
	}

	u1, err2 := service.GetUser("log")
	if err2 != nil {
		fmt.Println(u1.Email)
	} else {
		fmt.Println(err2)
	}

}
