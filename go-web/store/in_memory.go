package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

//* there are two ways of getting the post: either by a unique ID or by the name of the author

//* map the uniqueID to a pointer to a post
var PostById map[int]*Post

//* map the author's name to a slice of pointers to posts i.e. an array of posts
var PostsByAuthor map[string][]*Post

func store(post Post) {
	//* create post
	PostById[post.Id] = &post
	//* append post to author's posts
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Rob Holding"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Alex Lacazette"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Hector Bellerin"}
	post4 := Post{Id: 4, Content: "Greeting Earthlings", Author: "Thanos"}
	post5 := Post{Id: 5, Content: "Je m'apelle Alex!", Author: "Alex Lacazette"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)
	store(post5)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Alex Lacazette"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["Hector Bellerin"] {
		fmt.Println(post)
	}
}
