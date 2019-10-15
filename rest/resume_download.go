//* Resuming downloads after a network erro

package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//* create a local file to store the download
	file, err := os.Create("file.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//* download the remote file to the local file, retrying up to 100 times
	location := "https://example.com/file.zip"
	err = download(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	//* displays the size of the file after the download is complete
	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Got it with %v bytes downloaded", fi.Size())
}

func download(location string, file *os.File, retries int64) error {
	//* create a new GET request for the file being downloaded
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}
	//* start the local file to find the current file information
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	//* retrieve the size of the local file
	current := fi.Size()
	//* when the local file already has content, set a header requesting where the local file left off. Ranges have an index of 0, making the current length the index for the next needed byte
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		req.Header.Set("Range", "bytes="+start+"-")
	}
	//* an http client configured to explicitly check for timeout
	cc := &http.Client{Timeout: 5 * time.Minute}
	//* performs the request for the file or part if part of the file is already stored locally
	res, err := cc.Do(req)
	//* when checking for an error, tries the request again if the error was caused by a timeout
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	//* handle nonsuccess HTTP status codes
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errFmt := "Unsuccessful HTTP request. Status: %s"
		return fmt.Errorf(errFmt, res.Status)
	}
	//* if the server doesn't support serving partial files, set retries to 0
	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}

	//* copy the remote response to the local file
	_, err = io.Copy(file, res.Body)
	//* if a timeout occurs while copying the file, try retrieving the remaining content
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	return nil
}

func hasTimedOut(err error) bool {
	//* use a type switch to detect the type of underlying error
	switch err := err.(type) {
	//* a url error may be caused by an underlying net error that can be checked for a timeout
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	//* look for timeouts detected by the net package
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}
	//* some errors, w/out a custom type of varaiable to check against, can indicate a timeout
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}
