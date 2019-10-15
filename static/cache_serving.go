//* we could use reverse proxy via <varnish> to cache files, but here we load a file from disk & serve it from memory
package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

//* data structure to store a file in memory
type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

//* map to store files in memory
var cache map[string]*cacheFile

//* mutex to handle race conditions while handling parallel cache changes
var mutex = new(sync.RWMutex)

func main() {
	//* make the map usable
	cache = make(map[string]*cacheFile)
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	//* loads from the cache if it's already populated
	mutex.RLock()
	v, found := cache[req.URL.Path]
	mutex.RUnlock()

	//* when the file isn't in the cache, starts the loading process
	if !found {
		//* maps can't be written to concurrently or be read while being written to. Using a mutex prevents this from happening
		mutex.Lock()
		defer mutex.Unlock()
		//* open the file to be cached making sure to defer the close
		fileName := "files" + req.URL.Path
		f, err := os.Open(fileName)
		defer f.Close()

		//* handles an error when a file can't be opened
		if err != nil {
			http.NotFound(res, req)
			return
		}

		//* copy the file to an in-memory buffer
		var b bytes.Buffer
		_, err = io.Copy(&b, f)
		//* handle error copying from file to memory
		if err != nil {
			http.NotFound(res, req)
			return
		}
		//* put the bytes into a Reader for later use
		r := bytes.NewReader(b.Bytes())

		//* populates the cache object and stores it for later
		info, _ := f.Stat()
		v := &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}
		cache[req.URL.Path] = v
	}
	//* serve the file from cache
	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)
}
