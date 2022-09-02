package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Bob")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
