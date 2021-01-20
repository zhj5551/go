package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "test2.txt"
	// O_TRUNC清空文件内容，O_WRONLY创建文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开文件失败：", err)
		return
	}
	defer file.Close()

	reader := bufio.NewWriter(file)
	writer := "aaaaaa；\r\n"
	for i := 0; i < 5; i++ {
		reader.WriteString(writer)
	}
	reader.Flush()
}
