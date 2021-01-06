package main

import (
	"fmt"
	"go/MyGo/工厂函数/model"
)

func main() {
	tom := model.Student(10, "tom")
	fmt.Println(*tom)

}
