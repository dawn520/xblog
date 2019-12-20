package model

type User struct {
	Model
	Username            string
	Phone               string
	Email               string
	Password            string
	RememberToken      string
}

func (User) TableName() string {
	return "user"
}