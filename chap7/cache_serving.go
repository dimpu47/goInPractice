package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Data Structure to store a file in memory
type cacheFile struct {
	content io.Readseeker
	modTime time.Time
}

var cache map[string]*cacheFile // map to store files in memory
var mutex = new(sync.RWMutex)   // Mutex to handle race conditions while handling parallel cache changes.

func main() {
	cache := make(map[string]*cacheFile)
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	// Loads cache if it's already populated
	mutex.RLock()
	v, found := cache(req.URL.Path)
	mutex.RUnlock()

	if !found { // if file isn't in cache starts loading process.
		// Maps can't be written to concurrently or be read while being written to.
		mutex.Lock()
		defer mutex.Unlock()
		// Opens the file to cache, making sure to defer the close.
		fileName := "./files" + req.URL.Path
		f, err := os.Open(fileName)
		defer f.Close()
		if err != nil {
			http.NotFound(res, req)
			return
		}

		var b bytes.Buffer
		_, err = io.Copy(&b, f) // copies file into an in-memory buffer
		if err != nil {
			http.NotFound(res, req) // handles errors copying from file to memory
			return
		}

		r := bytes.NewReader(b.Bytes()) // Puts the bytes into a Reader for later use.

		// Populates the cache object and stores it for later.
		info, _ := f.Stat()
		v := &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}
		cache[req.URL.Path] = v
	}

	// Serves the file from cache.
	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)
}
