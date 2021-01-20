package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := `C:\Users\zhj5551\go\src\go\MyGo\test.txt`

	fileinfo, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Print(err)
		return
	}
	bs := []byte{0}
	num, err := fileinfo.Read(bs)
	fileinfo.Seek(4, io.SeekStart)
	num, err = fileinfo.Read(bs)
	fmt.Println(num, string(bs[:num]))
}
