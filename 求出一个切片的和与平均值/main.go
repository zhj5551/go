package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Sum 求出一个切片中所有元素的和
func Sum(slice []int) int {
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum = sum + slice[i]
	}
	return sum
}

// Arv 求平均值
func Arv(arvSlice []int) float64 {
	sum := 0
	for _, val := range arvSlice {
		sum = sum + val
	}
	return float64(sum / len(arvSlice))
}
func main() {
	rand.Seed(time.Now().UnixNano())
	// var slice []int
	slice := make([]int, 11, 11)
	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(100)
	}
	fmt.Println(slice)

	fmt.Println(Sum(slice))
	fmt.Printf("%v", Arv(slice))

}
