package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Welcome:")
	log.SetFlags(0)

	name := []string{"krishna", "Akash", "Vijay"}
	messages, err := greetings.Multi(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
