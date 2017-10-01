package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser
	file, err := OpenCSV("data.csv")
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return
	}
	defer file.Close()

	// Do something file
}

func OpenCSV(fname string) (f *os.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			f.Close()
			err = r.(error)
		}
	}()

	f, err = os.Open(fname)
	if err != nil {
		fmt.Printf("Failed to open file.")
		return f, err
	}

	removeEmptyLines(f)

	return f, err
}

func removeEmptyLines(f *os.File) {
	panic(errors.New("Failed parse."))
}
