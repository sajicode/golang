package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	//* create json file to store data
	jsonFile, err := os.Create("post3.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	//* create encoder with json file
	encoder := json.NewEncoder(jsonFile)
	//* encode struct into file
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}
