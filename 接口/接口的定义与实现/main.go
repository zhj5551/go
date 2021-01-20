package main

import (
	"fmt"
)

// Humaner 定义人类接口
type Humaner interface {
	//方法
	Say()
}

// Student 学生结构体
type Student struct {
	name  string
	score int
}

// Say 定义方法
func (s *Student) Say() {
	fmt.Println("Student[%s,%d]瞌睡不断\n", s.name, s.score)
}

// Teacher 定义结构体
type Teacher struct {
	name  string
	group string
}

// Say 定义方法
func (t *Teacher) Say() {
	fmt.Println("Teacher[%s,%s] 诲人不倦\n", t.name, t.group)
}

// Mystr 自定义类型
type Mystr string

// Say 定义方法
func (str Mystr) Say() {
	fmt.Println("Mystr[%s] 统治醒醒，还有个bug\n", str)
}

// Whosay 参数为接口类型
func Whosay(i Humaner) {
	i.Say()
}
func main() {
	s := &Student{"学生", 88}
	t := &Teacher{"老师", "GO语言"}
	var tmp Mystr = "字符串"
	s.Say()
	t.Say()
	tmp.Say()
	//多态，条用同一接口不同的表现
	Whosay(s)
	Whosay(t)
	Whosay(tmp)

	//make()
	x := make([]Humaner, 3)
	x[0], x[1], x[2] = s, t, tmp
	for _, value := range x {
		value.Say()
	}
}
