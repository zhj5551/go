package main

import (
	"fmt"
	"time"
)

func main() {

	intChan := make(chan int, 10)
	// intChan <- 1
	// intChan <- 3
	// intChan <- 3
	go func() {
		for i := 0; i < 10; i++ {
			intChan <- i
			fmt.Println("写数据", i)
			time.Sleep(time.Second * 2)
		}
		close(intChan)
	}()

	// v1, ok := <-intChan
	// v2, ok := <-intChan
	// v3, ok := <-intChan
	// fmt.Println(v1, ok)
	// fmt.Println(v2, ok)
	// fmt.Println(v3, ok)
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(v, ok)
		// time.Sleep(time.Second / 5)
	}
	// time.Sleep(time.Second * 15)
}
