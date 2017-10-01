package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// time.After returns <-chan time.Time (recieve only channel that recieves time.Time objects) that after 30 seconds will revieve a message.
	done := time.After(68 * time.Second)
	// makes a new channel for passing bytes from stdin to stdout. will hold one message at a time, no size specification.
	echo := make(chan []byte)
	// start goroutine to read stdin & pass it to echo channel for communicating.
	go readStdin(echo)

	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed Out")
			os.Exit(0)
		}
	}

}

// takes write only <-chan and sends any recieved input to that channel
func readStdin(out chan<- []byte) {
	for {
		// Copies some data from stdin in data. Note: .Read() vlocks until it recieves data
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data // sends the buffered data over the channel
		}
	}
}
