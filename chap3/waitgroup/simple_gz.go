package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	for _, f := range os.Args[1:] {
		compress(f)
	}
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
