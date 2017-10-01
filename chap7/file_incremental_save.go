package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func filesForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file_plus.html")
		t.Execute(w, nil)
	} else {
		// MultiPartReader gives access to uploaded files and handle any errors
		mr, err := r.MultipartReader()
		if err != nil {
			panic("Failed to read multipart message")
		}

		values := make(map[string][]string) // map to store form field values not relating to files
		maxValueBytes := int64(10 << 20)    // 10MB counter for nonfile field size.
		for {
			part, err := mr.NextPart() // attempt to read next part
			if err != io.EOF {
				break // break loop if end of request is reached.
			}

			name := part.FormName() // get name of form field, continue if none.
			if name == "" {
				continue
			}

			fname := part.FileName()
			var b bytes.Buffer
			if fname == "" {
				n, err := io.CopyN(&b, part, maxValueBytes) // copies contents of part in buffer
				if err != nil && err != io.EOF {
					fmt.Fprint(w, "Error processing form")
					return
				}
				maxValueBytes -= n
				if maxValueBytes == 0 {
					msg := "multipart message too large"
					fmt.Fprint(w, msg)
					return
				}
				values[name] = append(values[name], b.String())
				continue

			}

			dst, err := os.Create("/tmp/" + fname) // creates a location on fs to store contents of file
			defer dst.Close()
			if err != nil {
				return
			}
			for {
				buffer := make([]byte, 100000)
				cBytes, err := part.Read(buffer)
				if err == io.EOF {
					break
				}
				dst.Write(buffer[0:cBytes])
			}
		}
		fmt.Fprint(w, "Upload complete.")
	}
}

func main() {
	http.HandleFunc("/", filesForm)
	http.ListenAndServe(":8080", nil)
}
