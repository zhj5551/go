package main

import (
	"fmt"
)

// Person 定义结构体
type Person struct {
	Name string
	Age  int
}

func main() {

	var p1 Person
	p1.Age = 10
	p1.Name = "小明"
	var p2 *Person = &p1 // 这里是关键-->画出示意图

	fmt.Println((*p2).Age) // 10 底层已经把(*p2).Age与p2.Age等价了
	fmt.Println(p2.Age)    // 10
	p2.Name = "tom~"
	fmt.Printf("p2.Name=%v p1.Name=%v \n", p2.Name, p1.Name)    // tom~ tom~
	fmt.Printf("p2.Name=%v p1.Name=%v \n", (*p2).Name, p1.Name) // tom~ tom~

	fmt.Printf("p1的地址%p\n", &p1)
	fmt.Printf("p2的地址%p p2的值%p\n", &p2, p2)

}
