// Because msg is defined before closure. closure may reference it.
package main

import "fmt"

func main() {
	// Defines variable outside the closure
	var msg string

	defer func() {
		// Prints variable in defered closure.
		fmt.Println(msg)
	}()
	// sets value of variable.
	msg = "Hola Mundo!"
}
