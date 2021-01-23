package main

import "fmt"

func main() {

	intChan := make(chan int, 100)
	// 通过for循环向管道中添加100个元素
	for i := 0; i < 100; i++ {
		intChan <- i
	}

	// 在遍历时，如果管理已关闭，则会正常遍历数据，遍历完后，就会退出遍历close(chan)函数关闭
	//在遍历时，如果管理没有关闭，则会出现deadlock错误，
	close(intChan)
	//使用内置的函数close可以关闭管道，当管道关闭后，管道就只能读，不能写。
	// intChan <- 99
	// 遍历管道时，不能使用普通的for循环语句
	for v := range intChan {
		fmt.Println("v=", v)
	}
}
