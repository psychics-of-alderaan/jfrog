package main

import (
	"fmt"
	"github.com/psychics-of-alderaan/ktoad/pkg/things"
)

func main() {
	fmt.Println("Hello, World!")

	// Create a new Ktoad
	k := things.Ktoad{"Ktoad", "Hello, World!"}

	fmt.Printf("%s says %q\n", k.Name, k.Greeting)

}
