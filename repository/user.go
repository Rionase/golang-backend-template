package repository

import (
	"golang-backend-template/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	AddUser(userData model.UserAuthBody) (model.User, error)
	CheckUserAvail(username string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u userRepository) AddUser(userData model.UserAuthBody) (model.User, error) {
	newUserData := model.User{
		Username: userData.Username,
		Password: userData.Password,
		Role:     userData.Role,
	}
	if err := u.db.Create(&newUserData).Error; err != nil {
		return model.User{}, err
	}
	return newUserData, nil
}

func (u userRepository) CheckUserAvail(username string) (model.User, error) {
	userData := model.User{}
	if err := u.db.Select("id, role, password").Where("username = ?", username).First(&userData).Error; err != nil {
		return model.User{}, err
	}
	return userData, nil
}
