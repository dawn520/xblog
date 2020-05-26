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
