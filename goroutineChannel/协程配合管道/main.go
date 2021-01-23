package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("往管道中写数据：", i)
		time.Sleep(time.Second / 6)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {

	for {
		// 循环的读取管道中的数据，如果管道中的数据取完了，ok就会变成false
		// 如果写的比较慢，读取数据时，会等着写的
		v, ok := <-intChan
		if !ok {

			break
		}
		fmt.Println("从管道中读取数据：", v)
		time.Sleep(time.Second / 28)
	}

	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	// time.Sleep(time.Second * 10)

	for {

		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
