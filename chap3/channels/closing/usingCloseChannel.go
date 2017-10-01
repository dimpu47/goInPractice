// The best idiomatic way to close channels in go.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool) // additional bool chan indicating finish.
	until := time.After(5 * time.Second)

	go send(msg, done) // passing two channels to send func.

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true // on time out we indicate send to close chan
			fmt.Println("Timed Out.")
			os.Exit(0)
		}
	}
}

// ch is recieving, done is sending
func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			println("Done.")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
