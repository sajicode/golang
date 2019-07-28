//* demo of how to decode a JSON response

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//* the strings at the end of each field in the structs
//* are called tags & they're the mechanism to provide
//* metadata about the field mapping between the JSON document
//* and the struct type. If tags are not present, the decoding
//* & encoding process will attempt to match against the field
//* names directly in a case-insensitive way. When a mapping can't
//* be made, the field in the struct value will contain its zero value

type (
	//* gResult maps to the result document received from the search
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	//* gResponse contains the top level document
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	//* Issue the search against google
	resp, err := http.Get(uri)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	//* Decode the JSON response into our struct type
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}

//* url is outdated
