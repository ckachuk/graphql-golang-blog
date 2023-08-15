package categoryService

import (
	"blog-graphql/db"
	"blog-graphql/graph/model"
	"context"

	"github.com/google/uuid"
)

func CreateCategoryService(ctx context.Context, name string) (*model.Category, error){
	category := model.Category{ID: uuid.NewString(), Name: name}

	result := db.DB.Create(&category)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func DeleteCategoryService(ctx context.Context, _id string)(*model.Category, error){
	var category model.Category
	query := db.DB.First(&category, "id = ?", _id)
	if query.Error != nil {
		return nil, query.Error
	}
	result := db.DB.Delete(model.Category{}, "id = ?", _id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}


func UpdateCategoryService(ctx context.Context, _id string, name string)(*model.Category, error){
	var category model.Category
	query := db.DB.First(&category, "id = ?", _id)
	if query.Error != nil {
		return nil, query.Error
	}
	result := db.DB.Model(&category).Update("name", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}


func GetCategoryService(ctx context.Context, _id string) (*model.Category, error) {
	var category *model.Category

	result := db.DB.First(&category, "id = ?", _id)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func GetCategoriesService(ctx context.Context) ([]*model.Category, error) {
	var categories []*model.Category

	result := db.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}