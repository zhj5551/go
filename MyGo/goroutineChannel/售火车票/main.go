package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ticket = 10

// Sale 定义售票函数
func Sale(name string) {
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)

			fmt.Printf("%v %d\n", name, ticket)
			ticket--
		} else {
			fmt.Println("票卖完了")
			break
		}
	}
}
func main() {
	go Sale("售票口1")
	go Sale("售票口2")
	go Sale("售票口3")
	go Sale("售票口4")
	time.Sleep(3 * time.Second)
	fmt.Println("主线程结束")
}
