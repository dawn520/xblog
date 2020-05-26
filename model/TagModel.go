package model

type Tag struct {
	Model
	Name  string
	Alias string
	Count uint32
	Sort  int32
}
