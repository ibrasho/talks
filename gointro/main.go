package main

import (
	"fmt"
	"github.com/ibrasho/introtalk/model"
	"log"
)

func main() {
	user := model.NewUser(1, "username", "user@mail.com", "password")

	id := TestModel(user)
	fmt.Printf("ID: %v\n", id)

	if err := user.SendEmail(); err != nil {
		log.Println(err)
	}

	fmt.Printf("%v\n", user)
}

func TestModel(m model.M) int {
	return m.Id()
}
