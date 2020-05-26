package dao

import (
	"github.com/jinzhu/gorm"
	"xblog/model"
)

func CreatePost(post model.Post) (res uint, err2 error) {
	err := DB.Create(&post)
	if err.Error != nil {
		return post.ID, err.Error
	}
	return post.ID, nil
}

func UpdatePost(post model.Post) (res uint, err2 error) {
	oldPost, err := FindPostById(post.ID, "")
	if err != nil {
		return 0, err
	}
	err3 := DB.Model(&oldPost).Updates(post)
	if err3.Error != nil {
		return 0, err3.Error
	}
	return oldPost.ID, nil
}

func FindPostById(id uint, column string) (post model.Post, err2 error) {
	var err *gorm.DB
	if len(column) > 0 {
		err = DB.Select(column).First(&post, id)
	} else {
		err = DB.First(&post)
	}
	if err.Error != nil {
		return post, err.Error
	}
	return post, nil
}

func GetPostList(limit uint32, page uint32) (items []model.Post, total uint32, err error) {
	err2 := DB.Limit(limit).Offset((page - 1) * limit).Preload("Author").Find(&items).Count(&total)
	if err2.Error != nil {
		return nil, 0, err2.Error
	}
	err2 = DB.Find(&items).Count(&total)
	if err2.Error != nil {
		return nil, 0, err2.Error
	}
	return items, total, nil
}
