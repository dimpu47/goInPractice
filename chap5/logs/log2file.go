package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./logfile.txt")
	defer logfile.Close()

	f := log.Ldate | log.Lmicroseconds | log.Lshortfile | log.Llongfile
	logger := log.New(logfile, "example ", f)
	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error")
	logger.Println("This is end of this func.")
}
