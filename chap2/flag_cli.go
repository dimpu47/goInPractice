// working with cli applications the go way

package main

import (
	"flag"
	"fmt"
)

// create a new variable from a flag
var name = flag.String("name", "World", "A name to say hello to.")

// new variable to store flag value
var spanish bool

// set variable to flag value
func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use spanish language.")
}


func main() {

	// Parses the flags, placing values in variables
	flag.Parse()
	if spanish == true {
		if *name == "World" {
			*name = "Mundo"
		}
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
