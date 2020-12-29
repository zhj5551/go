package main

import "fmt"

func main() {
	array := [...]int32{3, 1, 66, 9, 29, 0, 15}
	n := len(array)
	fmt.Println("未排序前：", array)
	/*
	   冒泡排序
	*/
	for i := 0; i <= n-1; i++ {
		fmt.Println("第", i+1, "次冒泡")
		for j := i; j <= n-1; j++ {
			if array[i] > array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
				// break
			}
			fmt.Println(array)
		}
	}

	fmt.Println("最终结果", array)
}
