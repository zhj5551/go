package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "test2.txt"
	// O_TRUNC清空文件内容，O_WRONLY创建文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开文件失败：", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	readerWriter := bufio.NewWriter(file)
	writer := "你好这是新追加的内容；\r\n"
	for i := 0; i < 5; i++ {
		readerWriter.WriteString(writer)
	}
	readerWriter.Flush()
}
