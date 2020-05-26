package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//复制结构体
func StructCopy(data interface{}, response interface{}) {
	a := &data
	aj, _ := json.Marshal(a)
	json.Unmarshal(aj, &response)
}

func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	s = strings.Replace(s, "\\", "/", -1)
	s = strings.Replace(s, "\\\\", "/", -1)
	i := strings.LastIndex(s, "/")
	path := string(s[0 : i+1])
	return path
}

/**
截取中文字符串
*/
func SubChineseString(str string, begin, length int) string {
	fmt.Println("Substring =", str)
	rs := []rune(str)
	lth := len(rs)
	fmt.Printf("begin=%d,end=%d,lth=%d\n", length, lth)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}
