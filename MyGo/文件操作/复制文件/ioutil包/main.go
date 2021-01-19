package main

import (
	"fmt"
	"io/ioutil"
)

// 使用ioutil包中的 ioutil.WriteFile()和 ioutil.ReadFile()，但由于使用一次性读取文件，再一次性写入文件的方式，所以该方法不适用于大文件，容易内存溢出。
func copyFile(srcFile, destFile string) (int, error) {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		fmt.Println("操作失败：", destFile)
		fmt.Println(err)
		return 0, err
	}

	return len(input), nil
}
