package main

import (
	"fmt"
	"os"
)

func fileIsExit(s string) (bool, error) {
	_, err := os.Stat(s)
	if err == nil { // err等于nil，说明文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	filePath := "aa.txt"
	fmt.Println((fileIsExit(filePath)))

}
