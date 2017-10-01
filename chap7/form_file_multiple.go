package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file_multi.html")
		t.Execute(w, nil)
	} else {
		// Parse form in request and handle any errors
		err := r.ParseMultipartForm(16 << 20)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		// Retrieves a slice keyed by input name,containing files from MultipartForm.
		data := r.MultipartForm
		files := data.File["files"]
		// Iterates over all files uplloaded to field files.
		for _, fh := range files {
			// Opens a file handler
			f, err := fh.Open()
			// Be sure to close and handle any errors while opening fh
			defer f.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			// Creates local file to store contents of fh
			out, err := os.Create("/tmp/" + fh.Filename)
			// Close and handle errors while creating
			defer out.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			// Copies uploaded file to locationon filesystem
			_, err = io.Copy(out, f)
			// Handles any errors copying uploaded files to local files`
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}

		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
