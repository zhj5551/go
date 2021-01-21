package main

import (
	"fmt"
	"time"
)

/*
1. Go 语言追求简洁优雅，所以，Go 语言不支持传统的 try…catch…finally  这种处理。
2. Go 中引入的处理方式为：defer, panic, recover
3. 这几个异常的使用场景可以这么简单描述：Go 中可以抛出一个 panic 的异常，然后在 defer 中通过 recover 捕获这个异常，然后正常处理
*/

func test() {
	//使用defer + recover来捕获和处理异常
	defer func() {
		err := recover() // recover()是go内置函数，可以捕获到异常
		if err != nil {  //err 不为空
			fmt.Println("err=", err)
		}
	}() //匿名函数
	num1 := 10
	num2 := 0
	res := num1 / num2       //这句会报错，
	fmt.Println("res=", res) // 当前函数以下的无法执行了，但不退出程序
}

func main() {
	test()
	for i := 0; i < 3; i++ {
		fmt.Println("main()下面的代码...")
		time.Sleep(time.Second)
	}
}
