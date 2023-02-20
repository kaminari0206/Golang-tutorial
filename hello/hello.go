package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// fmt.Println(quote.Go())
	message := greetings.Hello("Kaminari")
	fmt.Println(message)
}

// Use the go mod edit command to edit the example.com/hello module to redirect Go tools
// from its module path (where the module isn't) to the local directory (where it is).
// $ go mod edit -replace example.com/greetings=../greetings
