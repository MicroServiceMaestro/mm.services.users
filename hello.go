package main

import (
	"fmt"
	"log"
	"mm.com/microservices/maestro/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(2)

	// Request a greeting message.
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Println("slmf")
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
