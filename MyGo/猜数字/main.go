package main

import "fmt"

func main() {

	var val int
	for {
		fmt.Printf("请输入数字: ")
		fmt.Scanln(&val)
		if val < 100 {
			fmt.Printf("你输入的%d数字太小\n", val)
		} else if val > 100 {
			fmt.Printf("你输入的%d数字太大\n", val)
		} else {
			fmt.Printf("你输入的%d数字正确\n", val)
			break
		}
	}
}
