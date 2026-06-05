package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name provided")
	}

	//message := fmt.Sprintf(randomFormat(), name) (REAL ONE COMMENTED OUT)
	message := fmt.Sprint(randomFormat()) // FAKE ONE TO FAIL A TEST CASE
	//:= operator used for declaring and initializing a variable in one line.
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	//A map to assoiate name with messages.
	messages := make(map[string]string)
	//loop through the received slice of names, calling
	//the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil { //if the name is empty
			return nil, err
		}
		//In the map, associate the retrived message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// random Format returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string { //function starts lower case, making it only accessable to this go file.
	// A SLICE of message formats(like an array but more flexible(Dynamic))
	formats := []string{ //omit the number, that way its dynamic.
		"hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
	// the return line does the following:
	// 1: checks elements count
	// 2: generates random number len-1
	// 3: indexes by that randomly generated number, then returns.
}
