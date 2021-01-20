package main

import "fmt"

func ercahshu(slice []int) *[]int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] { // 比较符合可以改变顺序
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return &slice
}

func main() {
	slice := []int{13, 6, 8, 4, 2}
	s1 := ercahshu(slice)
	fmt.Print(*s1)

}
