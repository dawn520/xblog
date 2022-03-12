package testing

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"xblog/dao"
	"xblog/model"
)

func TestGetUserByUserName(t *testing.T) {
	Convey("通过用户名获取用户", t, func() {
		username := "xixi"
		So(dao.GetUserByUserName(username), ShouldEqual, model.User{})
	})
}
