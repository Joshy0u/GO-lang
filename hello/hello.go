package main

import (
	"fmt"
	"go_tutorials/greetings"
	"log"
	// when running go mod init usually it would be example.com/WHATEVERMODULENAMEUWANT
	//but in this case, i already had it written as "go_tutorials", thats why i ran into an error.
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names

	names := []string{"Naruto", "Sasuke", "Hinata"}

	//request a greeting message
	message, err := greetings.Hellos(names)
	//if an error was returned, print it to the console and exit program
	if err != nil {
		log.Fatal(err)
	}

	//if no error, print the returned MAP of messages to console.
	fmt.Println(message)
}
