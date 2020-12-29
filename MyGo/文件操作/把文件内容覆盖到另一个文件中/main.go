package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 1, 目标文件可以不存在
	// 1，读取一个文件中的内容
	// 2，把读取的内容写到另一个文件中

	file1Path := "a.jpg"
	file2Path := "b.jpg"

	contents, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(file2Path, contents, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}
