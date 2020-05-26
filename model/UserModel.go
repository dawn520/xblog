package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username      string
	Nickname      string
	Phone         string
	Email         string
	Password      string
	RememberToken string
}

func (User) TableName() string {
	return "user"
}
