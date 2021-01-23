package main

import "fmt"

func main() {
	type cat struct {
		name string
		age  int
	}
	interfaceChan := make(chan interface{}, 3)
	interfaceChan <- 10
	interfaceChan <- "ab"
	newCat := cat{
		name: "zs",
		age:  90,
	}

	interfaceChan <- newCat

	<-interfaceChan
	<-interfaceChan
	cat1 := <-interfaceChan
	// 下面的写法是错误的，编译不通过，因为通道是口接口，编译器并不知道cat1的类型
	fmt.Printf("cat1的类型%T, 值:%v\n", cat1, cat1)
	// 使用类型断言，即可解决上述问题
	a := cat1.(cat)
	fmt.Printf("cat1的类型%T, 值:%v\n", cat1, a.name)

}
