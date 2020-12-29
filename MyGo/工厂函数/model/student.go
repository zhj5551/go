package model

// Student 定义学生结构体
type student struct {
	no   int
	name string
}

// Student 工厂函数
func Student(no int, name string) *student {
	return &student{
		no:   no,
		name: name,
	}
}

func (s *student) GetScore() string {
	return s.name
}
