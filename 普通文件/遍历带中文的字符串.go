package main

import "fmt"

func main() {
	str := "abcd张"
	fmt.Printf("%T\n", str)
	fmt.Println(len(str))
	for i := 0; i < len(str); i++ { //汉族占用不止一个字节
		fmt.Println(i, str[i])

	}
	fmt.Println()
	for i, v := range str { // 可以遍历汉字
		fmt.Println(i, string(v))

	}

}
