package main

import (
	"fmt"
	"log"

	"github.com/dmilojkovic76/echo_server/go_echo_server/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("echo_server: ")
	log.SetFlags(0)

	// Create a new message
	init_mgs, err := greetings.Hello(true)

	// Check if the message was succesfully created
	// and log an error if it exists
	if err != nil {
		log.Fatal(err)
	}

	// Output the message
	fmt.Println(init_mgs)
}
