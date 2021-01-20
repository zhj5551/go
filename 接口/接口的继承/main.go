package main

import "fmt"

//定义接口
type Humaner interface {
	//方法
	Say()
}
type Personer interface {
	//相当于写了say() 方法的继承
	Humaner
	//唱歌
	Sing(lyrics string)
}
type Student struct {
	name  string
	score int
}

func (s *Student) Say() {
	fmt.Printf("Student[%s,%d] 瞌睡不断\n", s.name, s.score) //Student[学生,80] 瞌睡不断
}

func (s *Student) Sing(lyrics string) {
	fmt.Printf("Student sing[%s]!!\n", lyrics) //Student sing[葫芦娃]!!
}
func main() {
	s := &Student{"学生", 80}
	//调Personer方法
	var p Personer
	p = s
	p.Say()
	p.Sing("葫芦娃")
}
