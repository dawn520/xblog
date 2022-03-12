package dao

import (
	"github.com/jinzhu/gorm"
	"xblog/model"
)

func GetUserById(id uint, column string) (user model.User, err2 error) {
	var err *gorm.DB
	if len(column) > 0 {
		err = DB.Select(column).First(&user, id)
	} else {
		err = DB.First(&user, id)
	}
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func GetUserByEmail(email string, column string) (user model.User, err2 error) {
	var err *gorm.DB
	if len(column) > 0 {
		err = DB.Select(column).Where("email = ?", email).First(&user)
	} else {
		err = DB.Where("email = ?", email).First(&user)
	}
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func GetUserByUsername(email string, column string) (user model.User, err2 error) {
	var err *gorm.DB
	if len(column) > 0 {
		err = DB.Select(column).Where("username = ?", email).First(&user)
	} else {
		err = DB.Where("username = ?", email).First(&user)
	}
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func CreateUser(user model.User) (res uint, err2 error) {
	err := DB.Create(&user)
	if err.Error != nil {
		return user.ID, err.Error
	}
	return user.ID, nil
}
