package main

import "fmt"

// 空接口类型的channel
type cat struct {
	name string
	age  int
}

func iterf() {
	interfaceChan := make(chan interface{}, 3)
	interfaceChan <- 10
	interfaceChan <- "ab"
}
func main() {
	// 声明chan通道
	var intChan chan int
	// 初始化通道
	intChan = make(chan int, 3)

	// intChan := make(chan int, 3) // 也可以直接初始化定义通道变量

	intChan <- 10
	<-intChan // 可以从通道中取出数据，没有变量接受
	intChan <- 11
	// intChan <- "aa" // 通道必须存放指定类型
	// intChan <- 11
	intChan <- 11 // 当往通道中添加的元素时，不能超过通道的容量时
	fmt.Printf("长度:%v, 容量:%v\n", len(intChan), cap(intChan))

	// 通道先进先出FIFO
	num1 := <-intChan
	fmt.Println("num1:", num1)
	num2 := <-intChan
	fmt.Println("num2:", num2)
	fmt.Printf("长度:%v, 容量:%v\n", len(intChan), cap(intChan))

	// num3 := <-intChan //在没有使用协程的情况下，如果channel数据取完了，再取就会报错 deadlock!
	// fmt.Println("num3:", num3)
	// fmt.Printf("长度:%v, 容量:%v\n", len(intChan), cap(intChan))
}
