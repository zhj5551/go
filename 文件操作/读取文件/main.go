package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fileinfo, err := os.Open(`C:\Users\zhj5551\go\src\go\MyGo\test.txt`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileinfo.Close()
	// fmt.Println(fileinfo)
	// fmt.Println(fileinfo.Name())
	bs := make([]byte, 4, 4)
	// n, _ := fileinfo.Read(bs)
	// fmt.Println(string(bs[:n]))
	// n2, _ := fileinfo.Read(bs)
	// fmt.Println(string(bs[:n2]))
	// n3, err := fileinfo.Read(bs)
	// fmt.Println(string(bs[:n3]), err)
	// n4, err := fileinfo.Read(bs)
	// fmt.Println(string(bs[:n4]), err)
	// fmt.Println("=======")
	for {
		num, err := fileinfo.Read(bs)
		if err != nil || err == io.EOF {
			// fmt.Print(err)
			break
		}
		fmt.Print(string(bs[:num]))
		time.Sleep(time.Second)
	}
}
