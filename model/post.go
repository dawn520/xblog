package model

type Post struct {
	Model
	AuthorId       int
	Title          string
	Digest         string `gorm:"type:text()"`
	Content        string `gorm:"type:text()"`
	status         int
	CategoryId     int
	TagId          string
	ViewNumbers    int
	CommentNumbers int
}
