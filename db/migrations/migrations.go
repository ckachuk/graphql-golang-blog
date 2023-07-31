package migrations

import (
	"blog-graphql/db"
	"blog-graphql/db/models"
)

func Migrate() {
	db.DB.AutoMigrate(models.User{}, models.Credentials{}, models.Post{}, models.Comment{}, models.Category{})
}