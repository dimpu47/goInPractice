package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in) // copies data to an os.Writer from os.Reader
}

func main() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(68 * time.Second)
	fmt.Println("Timed Out.I mean -- Done Sleeping.")
	os.Exit(0)
}
