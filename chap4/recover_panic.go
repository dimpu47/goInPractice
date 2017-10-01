/* Illustrates the most common pattern for panic recovery
using deffered closure function that checks for a panic
and recovers if it finds one.*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	// deffered closure for panic recovery.
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped Panic: %s (%T)\n", err, err)
		}
	}()

	yikes() // calls function that panic
}

// Emits a panic with error for a body.
func yikes() {
	panic(errors.New("Something bad happened."))
}
