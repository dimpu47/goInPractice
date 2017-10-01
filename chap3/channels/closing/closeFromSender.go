package main

import "time"

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch)
	for {
		select {
		case <-ch:
			println("Got message.")
		case <-timeout:
			println("Timed out")
			return
		default:
			println("**yawn**")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// send func closes channel ch, closed ch return nil which in select statements makes ch recieve false value even after ch is closed. [needs fix]
func send(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("Sent and closed.")
}
