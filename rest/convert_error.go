package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//* converting an unexpected status to an error

type Error struct {
	HTTPCode int    `json:"="`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

//* The error method implements the error interface on the Error struct
func (e Error) Error() string {
	fs := "HTTP: %d, Code: %d, Message: %s"
	return fmt.Sprintf(fs, e.HTTPCode, e.Code, e.Message)
}

//* the get function should be used instead of http.Get to make requests
func get(u string) (*http.Response, error) {
	//* use http.Get to retrieve the resource and return any http.Get errors
	res, err := http.Get(u)
	if err != nil {
		return res, err
	}

	//* check if the response code was outside the 200 range of successful responses
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		//* check the response content type & return an error if it's not correct
		if res.Header.Get("Content-Type") != "application/json" {
			sm := "Unknown error, HTTP status: %s"
			return res, fmt.Errorf(sm, res.Status)
		}

		//* read the body of the response into a buffer
		b, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		//* Parse the JSON response and place the data into a struct and respond to any errors
		var data struct {
			Err Error `json:"error"`
		}
		err = json.Unmarshal(b, &data)
		if err != nil {
			sm := "Unable to parse json: %s, HTTP status: %s"
			return res, fmt.Errorf(sm, err, res.Status)
		}
		//* Add the HTTP status code to the error instance
		data.Err.HTTPCode = res.StatusCode

		//* return the custom error and the response
		return res, data.Err
	}
	//* when there is no error, return the response as expected
	return res, nil
}

func main() {
	res, err := get("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
