package main

import (
	"errors"
	"time"

	"./safely"
)

func message() {
	println("Inside a goroutine.")
	panic(errors.New("Oops!"))
}

func main() {
	safely.Go(message)
	println("Outside goroutine")
	time.Sleep(1000)
}
