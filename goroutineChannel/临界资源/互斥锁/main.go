package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10
var wg sync.WaitGroup
var mutex sync.Mutex

// Sale 定义售票函数
func Sale(name string) {

	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000000)) * time.Microsecond)

			fmt.Printf("%v %d\n", name, ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "票卖完了")
			break
		}
		mutex.Unlock()
	}

}
func main() {
	wg.Add(4)
	go Sale("售票口1")
	go Sale("售票口2")
	go Sale("售票口3")
	go Sale("售票口4")
	// time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Println("主线程结束")
}
