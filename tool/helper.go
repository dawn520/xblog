package tool

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
