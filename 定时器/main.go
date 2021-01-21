package main

import (
	"fmt"
	"time"
)

/*
1.超时一次之后，就不再使用的定时器，time.After()。
2.每隔一段时间，就需要使用一次的定时器，time.Tick()。
3.阻塞住goroutinue的定时器，time.Sleep()，准确来说这个不算一个定时器
4.可以自由控制定时器启动和关闭的定时器，time.Ticker()。
*/

/*
1. 调用time.After(duration)，此函数马上返回，返回一个time.Time类型的Chan，不阻塞。后面你该做什么做什么，不影响。到了duration时间后，自动塞一个当前时间进去，底层是用NewTimer实现的。
2. 使用了time.After之后，只有在调用了<-tchan之后，才会真正的触发超时的操作，在此之前的代码是可以被立马执行的，因为time.After不会阻塞当前的goroutine
*/

// time.After() 当取出通道数据时，会阻塞duration时间
func after() {
	tchan := time.After(time.Second * 3)
	fmt.Println(time.Now().String(), "tchan type:%T", tchan)
	fmt.Println(time.Now().String(), "mark 1")
	fmt.Println(time.Now().String(), "tchan=", <-tchan) // channel里取出数据之后，会发现超时间是3秒
	fmt.Println(time.Now().String(), "mark 2")
}

// time.tick 每隔两秒运行一次for循环中的命令
func tick() {
	c := time.Tick(2 * time.Second)
	for next := range c {
		fmt.Printf("%v \n", next)
		fmt.Println("aaa")
	}
}

// time.Sleep() 阻塞duration时间后，代码继续运行
func sleep() {
	fmt.Println("aa")
	time.Sleep(time.Second)
	fmt.Println("bb")
}

// Ticker 和 Timer 类似，区别是：Ticker 中的 runtimeTimer 字段的 period 字段会赋值为 NewTicker(d Duration) 中的 d，表示每间隔 d 纳秒，定时器就会触发一次。除非程序终止前定时器一直需要触发，否则，不需要时应该调用 Ticker.Stop 来释放相关资源。
func ticker() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second) // 5s后触发time.Stop，关闭ticker
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

func main() {
	// after()
	// tick()
	ticker()
}
