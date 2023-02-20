package greetings

import (
	"errors"
	"fmt"
)

//In Go, the := operator is a shortcut for declaring and initializing a variable in one line
// (Go uses the value on the right to determine the variable's type).
//Taking the long way, you might have written this as:

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
