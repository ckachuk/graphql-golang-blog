// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type Comment struct {
	ID        string  `json:"_id"`
	Body      string  `json:"body"`
	User      *User   `json:"user,omitempty"`
	Post      *Post   `json:"post"`
	CreateAt  *string `json:"createAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type Credentials struct {
	ID       string `json:"_id"`
	IsAdmin  bool   `json:"isAdmin"`
	IsAuthor bool   `json:"isAuthor"`
}

type Post struct {
	ID        string     `json:"_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Category  *Category  `json:"category,omitempty"`
	User      *User      `json:"user"`
	Comments  []*Comment `json:"comments"`
	CreateAt  *string    `json:"createAt,omitempty"`
	UpdatedAt *string    `json:"updatedAt,omitempty"`
}

type User struct {
	ID          string       `json:"_id"`
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Name        string       `json:"name"`
	Credentials *Credentials `json:"credentials"`
	Token       string       `json:"token"`
}
