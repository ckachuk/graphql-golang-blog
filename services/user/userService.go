package userService

import (
	"blog-graphql/db"
	"blog-graphql/graph/model"
	jwtService "blog-graphql/services/jwt"
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(ctx context.Context, username string, password string, name string) (*model.User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	idUser := uuid.NewString()
	user := model.User{ID: idUser, Username: username, Password: string(bytes), Name: name, Credentials: &model.Credentials{ID: uuid.NewString(), UserID: idUser, IsAuthor: false, IsAdmin: false}, Token: ""}

	db.DB.Create(&user)
	return &user, nil
}

func LoginService(ctx context.Context, username string, password string) (*model.User, error) {
	var user model.User

	query := db.DB.Where("username = ?", username).First(&user)
	if query.Error != nil {
		return nil, query.Error
	}

	value := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if value != nil {
		return nil, value
	}

	token, err := jwtService.GenerateToken(user.Username)

	if err != nil {
		return nil, err
	}
	db.DB.Model(&user).Update("token", token)

	return &user, nil
}


func UpdatePasswordService(ctx context.Context, username string, oldPassword string, newPassword string) (*model.User, error) {
	var user model.User

	query := db.DB.Where("username = ?", username).First(&user)
	if query.Error != nil {
		return nil, query.Error
	}

	value := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if value != nil {
		return nil, value
	}

	newPasswordHashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	if err != nil {
		return nil, err
	}

	result := db.DB.Model(&user).Update("password", string(newPasswordHashed))
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUsersService(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	result := db.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUserService(ctx context.Context, _id string) (*model.User, error) {
	var user *model.User

	result := db.DB.First(&user, "id = ?", _id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
