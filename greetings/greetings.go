package greetings

import "fmt"

//In Go, the := operator is a shortcut for declaring and initializing a variable in one line
// (Go uses the value on the right to determine the variable's type).
//Taking the long way, you might have written this as:

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
