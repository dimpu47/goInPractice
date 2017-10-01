package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}

// A struct for the foundation of base implementation.
type LocalFile struct {
	Base string // stores base path for file storage.
}

// opens a file stored locally or returns an error.
func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	return os.Open(p)
}

func (l LocalFile) Save(path string, body io.ReadSeeker) error {
	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)
	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	return err
}

func fileStore() (File, error) {
	return &LocalFile{Base: "."}, nil
}

func main() {
	content := `Lorem ipsum dimsum dolo ait asddk asoegvc,afgib sdfiob xciog aew lsri dzqwe.`
	body := bytes.NewReader([]byte(content))

	store, err := fileStore()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Storing file content...")
	err = store.Save("foo/bar", body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Retrieving file content...")
	c, err := store.Load("foo/bar")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	o, err := ioutil.ReadAll(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(o))

}
