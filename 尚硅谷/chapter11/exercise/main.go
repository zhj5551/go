package main

import "fmt"

// Usb 定义接口
type Usb interface {
	Say()
}

// Stu 定义结构体
type Stu struct {
}

// Say 实现接口方法
func (s *Stu) Say() { // *Stu指针实现的Usb接口的Say()方法，但Stu并没有实现
	fmt.Println("Say()")
}

func main() {

	var stu Stu = Stu{}
	fmt.Println(stu)
	// 错误！ 会报 Stu类型没有实现Usb接口 ,
	// 如果希望通过编译,  var u Usb = &stu
	/*
		var u Usb = stu // 会报 Stu类型没有实现Usb接口 ,
		u.Say()
		fmt.Println("here", u)
	*/
}
