package model

type Category struct {
	Model
	Name        string
	Sort        int
	Count       int
	Alias       string
	Description string
	RootId      int
	ParentId    int
}
