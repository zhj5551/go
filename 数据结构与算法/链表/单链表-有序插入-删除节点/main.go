package main

// 插入链表，自动排列有序
import "fmt"

// HeroNode 定义链表的结构体
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

// Push 添加链表
func Push(head, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("链表号重复")
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

// Del 删除链表中的一个节点
func Del(head *HeroNode, id int) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			flag = false
			break
		} else if temp.next.no == id {
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("链表节点的id不存在")
		// return
	} else {
		temp.next = temp.next.next
	}
}

// Show 显示链表内容
func Show(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("[%d, %s, %s]---->", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
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
		no:       2,
		name:     "lujunyi",
		nickname: "卢俊义",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "linchong",
		nickname: "林冲",
	}
	hero4 := &HeroNode{
		no:       4,
		name:     "wusong",
		nickname: "武松",
	}
	hero5 := &HeroNode{
		no:       5,
		name:     "husanniang",
		nickname: "扈三娘",
	}
	Push(head, hero1)
	Push(head, hero5)
	Push(head, hero2)
	Push(head, hero4)
	Push(head, hero3)
	Del(head, 4)
	Del(head, 1)
	Del(head, 5)
	Del(head, 3)
	Del(head, 2)
	Del(head, 2)
	Show(head)

}
