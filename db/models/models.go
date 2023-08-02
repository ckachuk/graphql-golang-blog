package models

import (
	"github.com/google/uuid"
)



type Credentials struct{
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID int
	IsAdmin bool
	IsAuthor bool
}

type Post struct{
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title string
	Body string
	CategoryID int
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
	UserID int
	User User `gorm:"foreignKey:UserID;references:ID"`
	Comments []Comment 
}

type Comment struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Body string
	UserID int
	User User `gorm:"foreignKey:UserID;references:ID"`
	PostID int
	Post Post `gorm:"foreignKey:PostID;references:ID"`
}

type Category struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string
}

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string
	Password string
	Name string
	Credentials Credentials
}

