package main

import "fmt"

//Max 求一个切片中的最大值和对应的下标
func Max(intArr []int) (maxVal, index int) {
	maxVal = intArr[0]
	index = 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			index = i
		}
	}
	return
}

func main() {
	intArr := []int{4, 5, 6, 2, 38, 8, 9, 5, 6, 16}
	maxVal, index := Max(intArr)
	fmt.Printf("最大值是%v, 下标是%d", maxVal, index)

}
