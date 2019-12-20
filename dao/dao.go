package dao

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"xblog/config"
)

var DB *gorm.DB
var err error

func init() {
	//fmt.Println("asdasdas",config.DataBase)

	host := config.GetString("database.connections.mysql.host")
	port := config.GetString("database.connections.mysql.port")
	username := config.GetString("database.connections.mysql.username")
	password := config.GetString("database.connections.mysql.password")
	database := config.GetString("database.connections.mysql.database")
	//logs.Info("info" + username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.GetString("database.connections.mysql.prefix") + defaultTableName
	}
	url := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
	DB, err = gorm.Open("mysql", url)
	if err != nil {
		panic("failed to connect database")
	} else {
		logs.Info("database is connected:", url)
	}
	DB.LogMode(config.GetBool("app.debug"))
	//defer db.Close()
}
