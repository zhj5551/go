package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(s string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("test()--%v--%v\n", s, strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	wg.Done()
}

func test1() {
	for i := 5; i < 10; i++ {
		fmt.Println("test1--", i)
		time.Sleep(time.Second)
	}
	wg.Done()
}
func main() {

	// go test()
	wg.Add(1)
	go test("aaa")
	wg.Add(1)
	go test1()
	wg.Wait()
}
