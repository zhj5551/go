package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, _ := os.Getwd() // 打印所在的位置
	fmt.Println(pwd)
	err := os.Chdir("c:/users") // 切换目录
	s, _ := os.Getwd()
	fmt.Println(err, s)
}
