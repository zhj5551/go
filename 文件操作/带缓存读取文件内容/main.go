package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := `test.txt`
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			// fmt.Printf("文件读取完毕%v", err)
			return
		}
		fmt.Print(str)
	}

}
