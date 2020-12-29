package main

// 每次都在链表的最后添加链表

import "fmt"

// HeroNode 定义链表结构体
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

// Push 添加链表元素
func Push(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	temp.next = newHeroNode
}

// Show 显示链表内容
func Show(head *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		fmt.Printf("[%d, %s, %s]---->", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
	}

}

func main() {
	head := &HeroNode{}

	hero1 := &HeroNode{
		no:       1,
		name:     "songjiang",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       3,
		name:     "lujunyi",
		nickname: "卢俊义",
	}
	hero3 := &HeroNode{
		no:       2,
		name:     "lujunyi",
		nickname: "卢俊义",
	}
	Push(head, hero1)
	Push(head, hero2)
	Push(head, hero3)
	Show(head)

}
