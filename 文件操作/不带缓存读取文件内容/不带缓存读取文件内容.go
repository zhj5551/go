package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := "test.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(content)
	fmt.Println(string(content))

}
