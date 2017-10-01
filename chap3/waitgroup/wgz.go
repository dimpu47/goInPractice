package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup // waitgroup definition.
	var i int = -1
	var f string
	for i, f = range os.Args[1:] {
		wg.Add(i)
		go func(f string) {
			compress(f)
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Printf("Compresses %d files", i+1)

}

func compress(fname string) error {
	in, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(fname + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	// gzip.Writer compresses data and then writes it to underlying file.
	gzout := gzip.NewWriter(out)

	_, err = io.Copy(gzout, in)
	if err != nil {
		return err
	}

	return err
}
