package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Kieran Tierney"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre Aubameyang"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Dani Ceballos"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Darth Vader"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	//* reading a csv file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//* we set to a negative number because we aren't bothered if all fields aren't present for each record
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()

	//* the record is a slice of slices i.e. an array of arrays

	if err != nil {
		panic(err)
	}

	var posts []Post

	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
