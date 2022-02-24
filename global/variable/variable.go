package variable

import (
	"fmt"
	"log"
	"novel_learning/global/my_errors"
	"os"
	"strings"
)

var (
	BasePath string // 定义项目的根目录
	// 用户自行定义的其他全局变量
)

/* *******
定义BasePath这个关键变量<PS:调用variable文件会 自动执行文件下的init函数，有几个就执行几个>
* *******/
func init() {
	fmt.Println("开始执行variable文件")
	if path, err := os.Getwd(); err == nil {
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(path, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = path
		}
	} else {
		log.Fatal(my_errors.ErrorsBasePath)
	}
	fmt.Println("init函数执行完毕， BasePath=", BasePath)
}
