package bootstrap

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {
	initDB()

}
func initDB() {
	//fmt.Println("asdasdas",bootstrap.DataBase)

	host := GetString("database.connections.mysql.host")
	port := GetString("database.connections.mysql.port")
	username := GetString("database.connections.mysql.username")
	password := GetString("database.connections.mysql.password")
	database := GetString("database.connections.mysql.database")
	//logs.Info("info" + username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return GetString("database.connections.mysql.prefix") + defaultTableName
	//}
	url := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=true"
	logs.Info(url)
	DB, err = gorm.Open("mysql", url)
	if err != nil {
		logs.Info(err)
		panic("failed to connect database")
	} else {
		logs.Info("database is connected:", url)
	}
	DB.LogMode(GetBool("app.debug"))
	//defer db.Close()
}
