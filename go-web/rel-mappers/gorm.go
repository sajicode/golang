package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    `sql:"index"`
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "House on fire!", Author: "Sia Furler"}
	fmt.Println(post)

	//* create a post
	Db.Create(&post)
	fmt.Println(post)

	//* create a comment and link it to newly created post
	comment := Comment{Content: "The Greatest!", Author: "Kendrick Lamar"}
	Db.Model(&post).Association("Comments").Append(comment)

	//* get the comments of a post
	var readPost Post
	//* look for the first post that has an Author called Sia Furler & push the result into the readPost variable
	Db.Where("author = $1", "Sia Furler").First(&readPost)
	var comments []Comment

	//* get the post model & then get the related comments
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
