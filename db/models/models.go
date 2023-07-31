package models

import (
	"gorm.io/gorm"
)


type Credentials struct{
	ID uint `gorm:"primaryKey"`
	isAdmin bool
	isAuthor bool
}

type Post struct{
	gorm.Model
	title string
	body string
	category Category `gorm:"embedded"`
	user User `gorm:"embedded"`
	comments []Comment `gorm:"embedded"`
}

type Comment struct {
	gorm.Model
	body string
	user User
	post Post
}

type Category struct {
	ID uint `gorm:"primaryKey"`
	name string
}

type User struct {
	gorm.Model
	username string
	password string
	name string
	credentials Credentials `gorm:"embedded"`
}

