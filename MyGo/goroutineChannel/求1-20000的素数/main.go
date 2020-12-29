package main

import (
	"fmt"
	"time"
)

func suShu(num1 int) {

	flag := true
	for j := 2; j < num1; j++ {
		if num1%j == 0 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Print(num1, " ")
	}
	flag = true
}

func main() {
	// suShu(10000000)
	// fmt.Print("2 ")
	startTime := time.Now()
	for num := 2; num < 20000; num++ {
		suShu(num)
		time.Sleep(100)
	}
	//  计算程序运行的时间
	fmt.Println()
	fmt.Println((time.Now().UnixNano() - startTime.UnixNano()) / 1000)

}
