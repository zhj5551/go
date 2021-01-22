package main

import (
	"fmt"
)

func main() {
	a := [...]string{"USA", "china", "india"}
	b := a
	b[0] = "Singapore"

	fmt.Println("a is", a)  // ["USA", "china", "india"]
	fmt.Println("b is ", b) // ["Singapore", "china", "india"]

	for i, v := range a {
		fmt.Println(i, v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	d := [7]int{1, 2, 3, 4, 5, 6, 7}
	var e []int = d[1:4]
	fmt.Println(e)

	f := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, v := range f {
		fmt.Println(v)
	}

	ff := [][][]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{1, 2},
			{3, 4},
		},
	}
	fmt.Println(ff)
	for _, v := range ff {
		fmt.Println(v)
	}

	//var m [][]int   //多维数组的切片，通过循环一下
	//for i, v :=range f{
	//  fmt.Println(v[1:3])
	//  m = append(m, f[i][1:3])
	//  fmt.Println(m)
	//}

	//fmt.Println("\n")
	//fmt.Println( m)

	u := [7]int{1, 2, 2, 3, 4, 5, 7}
	t := u[1:3]
	fmt.Println(t)
	fmt.Println("len", len(t))
	fmt.Println("cap", cap(t))
}
