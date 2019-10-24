package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//* we pass the underscore before we call the package bcos we are not meant to call the database driver directly.
//* This way, if we upgrade the version of the driver or change the implementation of the driver, we don't need to make changes to all our code

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	//* connect to the database
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

//* get all posts
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

//* get a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

//* create a post
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		return err
	}
	return
}

//* update a post
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set Content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

//* delete a post
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	post := Post{Content: "La Vache qui rui", Author: "Ratatouille"}
	post2 := Post{Content: "La Notre Damme", Author: "Humpty"}

	fmt.Println(post)
	err := post.Create()

	if err != nil {
		panic(err)
	}

	fmt.Println(post2)
	post2.Create()

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost2, _ := GetPost(post2.Id)
	fmt.Println(readPost2)

	readPost.Content = "Illigitimi non carborundum!"
	readPost.Author = "Plato"
	readPost.Update()

	posts, _ := Posts(2)
	fmt.Println(posts)

	readPost2.Delete()
}
