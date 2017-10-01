/*
using buffered queues for locking while preventing send ops from blocking while theres room in the queue.
Buffre Queues are also used fro constructing message queues and pipelines.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create buffered channel with space 1.
	lock := make(chan bool, 1)

	// Starts up six goroutine shring the locking channel.
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)

}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock.\n", id)
	// A worker requires lock by sending it a message. First worker to hit this will get the lock, rest will block.
	lock <- true
	fmt.Printf("%d got the lock.\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock.\n", id)
	// Releases lock by sending value and opening space on buffer for next func to acquire lock
	<-lock
}
