package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Magnus")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Gladys", "Samanta", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	for _, message := range messages {
		fmt.Println(message)
	}
}
