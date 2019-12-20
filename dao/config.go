package dao

import (
	"duu-common/model"
	"fmt"
)

func GetRecommendTypeList(page int32, limit int32) (data []model.RecommendType) {
	DB.Limit(limit).Offset((page - 1) * limit).Find(&data)
	fmt.Printf("%+v\n", data)
	return data
}

func UpdateRecommend(id int32, data model.RecommendType) (response int32) {
	data.Id = 0
	DB.Model(&data).Where("id = ?",id).Update(&data)
	return 1
}
