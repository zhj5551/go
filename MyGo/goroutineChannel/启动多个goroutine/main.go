package main

import (
	"fmt"
	"sync"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "a")
		time.Sleep(time.Second / 10)
	}
	wg.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "b")
		time.Sleep(time.Second / 10)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go a()
	wg.Add(1)
	go b()
	for i := 0; i < 10; i++ {
		fmt.Println(i, "main")
		time.Sleep(time.Second / 10)
	}
	wg.Wait()
}
