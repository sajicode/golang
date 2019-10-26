package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Post struct
type Post struct {
	Id         int
	Content    string
	AuthorName string `db: author` //* this line tells the code to map AuthorName to the author in the db
}

// Db instantiation
var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// GetPost - a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author, from posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

// Create Post function
func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "I survived!", AuthorName: "Sia Furler"}
	post.Create()
	fmt.Println(post)
}
