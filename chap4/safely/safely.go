package safely

import (
	"log"
)

// GoDoer is a simple parameter less func
type GoDoer func()

// Safely.go runs a function as goroutine and handles any panics
func Go(todo GoDoer) {
	// run anonymous function
	go func() {
		// handle panics with usual deffering recovery pattern.
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic in safely.go: %s", err)
			}
		}()
		// Call GoDoer that was passed in.
		todo()
	}()
}
