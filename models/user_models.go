package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserModel struct {
	DB *gorm.DB
}

func NewUserModel(DB *gorm.DB) *UserModel {
	return &UserModel{DB: DB}
}

func (user_model *UserModel) GetAllUsers(name, email, password string) ([]User, error) {
	users := []User{}
	if err := user_model.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
