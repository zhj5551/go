package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertsort(slice *[]int) *[]int {
	for i:=0;i<len(*slice);i++ {
		flag:=false
		index:=i
		num:=(*slice)[i]
		for j:=i+1;j<len(*slice);j++ {
			if num > (*slice)[j] {
				index = j		
				num = (*slice)[j]
				flag=true
			}
		}
		if flag {
			(*slice)[i],(*slice)[index] = (*slice)[index], (*slice)[i]
		}
	}	
	return slice
}
func  main() {
	var slice []int
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<100000000;i++ {
		slice = append(slice, rand.Intn(1000000))
	}
	fmt.Println("aaa")
	fmt.Println(len(slice))
	fmt.Printf("%v,%d\n", slice, len(slice))
	//slice1 := insertsort(&slice)
	//fmt.Println(*slice1)
}


