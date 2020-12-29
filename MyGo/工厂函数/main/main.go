package main

import (
	"MyGo/工厂函数/model"
	"fmt"
)

func main() {
	tom := model.Student(10, "tom")
	fmt.Println(*tom)

}
