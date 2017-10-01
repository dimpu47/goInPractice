package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine.")
	go func() {
		fmt.Println("Inside a function.")
	}()
	fmt.Println("Outside again.")

	runtime.Gosched()
}
