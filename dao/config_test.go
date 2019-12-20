package dao

import (
	"duu-common/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUpdateRecommend(t *testing.T) {
	Convey("修改Recommend", t, func() {
		data := model.RecommendType{Name:"1天啊1"}
		So(UpdateRecommend(1,data), ShouldEqual,1)
	})
}
