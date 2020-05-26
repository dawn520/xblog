package config

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/ini.v1"
	"log"
	"strings"
	"xblog/utils"
)

var Config map[string]interface{}

func init() {

	cfg, err := ini.Load(utils.GetCurrentPath() + "config.ini")
	log.Printf(utils.GetCurrentPath() + "config.ini")
	if err != nil {
		logs.Error("error: %v\n", err)
	}
	/**
	APP CONFIG
	*/
	app := map[string]interface{}{
		"listening": cfg.Section("common").Key("listening").MustString(":17171"),
		"debug":     true,
		"key":       "asdxcasda",
		"protoName": "/proto.CommonService",
	}

	/**
	dataBase config
	*/
	database := map[string]interface{}{
		"default": "mysql",
		"connections": map[string]interface{}{
			"mysql": map[string]interface{}{
				"host":      cfg.Section("mysql").Key("host").MustString("127.0.0.1"),
				"port":      cfg.Section("mysql").Key("port").MustString("3306"),
				"database":  cfg.Section("mysql").Key("database").MustString("test"),
				"username":  cfg.Section("mysql").Key("username").MustString("root"),
				"password":  cfg.Section("mysql").Key("password").MustString("root"),
				"charset":   "utf8mb4",
				"collation": "utf8mb4_unicode_ci",
				"prefix":    cfg.Section("mysql").Key("prefix").MustString(""),
				"strict":    false,
			},
		},
	}

	/**
	jwt config
	*/
	jwt := map[string]interface{}{
		"key": "asdsadadasdqwe1123",
		"ttl": 24 * 3600 * 30,
	}

	methodMiddleware := map[string]interface{}{
		"auth": []string{
			"/proto.CommonService/CreatePost",
		},
		"valid": []string{
			"/proto.CommonService/CreatePost",
		},
	}

	Config = map[string]interface{}{
		"app":              app,
		"database":         database,
		"jwt":              jwt,
		"methodMiddleware": methodMiddleware,
	}

}

func GetString(a string) string {
	return Get(a).(string)
}
func GetBool(a string) bool {
	return Get(a).(bool)
}
func GetInt(a string) int {
	return Get(a).(int)
}

func GetInt64(a string) int64 {
	return Get(a).(int64)
}

func GetArrayString(a string) []string {
	return Get(a).([]string)
}

func Get(key string) interface{} {
	//fmt.Printf("%v", Config)

	keyArray := strings.Split(key, ".")
	a := getValue(Config, keyArray, 0)
	b := a[keyArray[len(keyArray)-1]]
	return b
}

func getValue(a map[string]interface{}, keyArray []string, i int) map[string]interface{} {
	if i < len(keyArray)-1 {
		value := keyArray[i]
		return getValue(a[value].(map[string]interface{}), keyArray, i+1)
	}
	return a
}
