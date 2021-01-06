package main

import (
	"fmt"
	"go/尚硅谷/使用结构体实现家庭收支记账软件/utils"
)

func main() {
	fmt.Println("面向对象的方式来完成.....")
	utils.NewMyFamilyAccount().MainMenu()
}
