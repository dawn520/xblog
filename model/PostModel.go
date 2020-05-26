package model

type Post struct {
	Model
	AuthorId       uint   `valid:"required"`
	Title          string `valid:"required,length(3|200)" `
	Digest         string `gorm:"type:text"` //摘要
	Content        string `gorm:"type:text" valid:"required"`
	Status         int
	CategoryId     int `valid:"required"`
	TagId          string
	ViewNumbers    int
	CommentNumbers int

	Author User `gorm:"foreignkey:AuthorId"`
}
