package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // 返回Time结构体
	fmt.Println(t.Date())
	fmt.Printf("%T %v\n", t.Unix(), t.Unix())
	fmt.Println(t.Zone())
	time.Sleep(1 * 10000000000) // 阻塞一个时间段，纳秒为单位
	fmt.Print("aaa")
}
