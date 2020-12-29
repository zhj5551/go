package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "test2.txt" // 把新内容覆盖到此文件中
	// O_TRUNC清空文件内容，O_WRONLY创建文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewWriter(file)
	writer := "你好原内容已被覆盖掉；\r\n" // 要写入文件的内容
	for i := 0; i < 5; i++ {
		reader.WriteString(writer)
	}
	reader.Flush()
}
