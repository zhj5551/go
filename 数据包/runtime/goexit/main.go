package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//创建新建的协程
	go func() {
		fmt.Println("goroutine开始。。。")

		//调用了别的函数
		fun()

		fmt.Println("goroutine结束。。")
	}() //别忘了()

	//睡一会儿，不让主协程结束
	time.Sleep(3 * time.Second)
	fmt.Println("退出mian线程")
}

func fun() {
	defer fmt.Println("defer。。。")
	runtime.Goexit() //退出当前 goroutine(但是defer语句会照常执行)
	fmt.Println("fun函数。。。")
}
