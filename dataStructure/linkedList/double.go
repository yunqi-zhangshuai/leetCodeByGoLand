package linkedList

import "fmt"

type Node2 struct {
	Pre  *Node2      //上一个节点
	Next *Node2      //指向下一个节点
	Data interface{} //数据
}

type Double struct {
	head    *Node2 //头指针 //头指针不放数据
	last    *Node2 //尾指针
	size    int32  //链表长度
	current *Node  //当前节点指针
}

func (l *Double) AppendNode(values ...interface{}) {
	element := &Node2{Pre: l.last}
	//头指针单独处理
	if l.size <= 0 {
		l.head = element
		l.last = element
	}
	for _, value := range values {
		node := &Node2{Data: value, Pre: l.last}
		l.last.Next = node
		l.last = node

		l.size++
	}
}

func (l *Double) Traverse() {
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
		fmt.Println("从头部遍历数据", cur.Data)
	}

	cur = l.last
	for cur.Pre != nil {
		fmt.Println("从尾部遍历数据", cur.Data)
		cur = cur.Pre
	}
}
