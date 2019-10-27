package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:      1,
		Content: "Silicon Valley!",
		Author: Author{
			Id:   2,
			Name: "Steve Jobs",
		},
		Comments: []Comment{
			Comment{
				Id:      3,
				Content: "Windows!",
				Author:  "Bill Gates",
			},
			Comment{
				Id:      4,
				Content: "Home PCs!",
				Author:  "Steve Wozniak",
			},
		},
	}

	//* marshall structs to byte slice of JSON data
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	err = ioutil.WriteFile("post2.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
