package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Num 对切边翻转打印
func Num(slice []int) []int {
	mid := len(slice) / 2

	for i := 0; i > mid; i++ {
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}
	return slice

}

func main() {
	var slice []int
	num := 11
	slice = make([]int, num)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		slice[i] = rand.Intn(100)
	}
	fmt.Println(slice)
	newSlice := Num(slice)
	fmt.Println(newSlice)

}
