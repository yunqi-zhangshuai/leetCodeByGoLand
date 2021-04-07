package linkedList

import (
	"fmt"
)

// DoubleNode  双向链表 node
type DoubleNode struct {
	Prev *DoubleNode //上一个节点
	Next *DoubleNode //指向下一个节点
	Data interface{} //数据
}

// SetPrev
// 设置前前趋节点
func (d *DoubleNode) SetPrev(node *DoubleNode) {
	d.Prev = node
}

// SetNext
// 设置后趋node
func (d *DoubleNode) SetNext(node *DoubleNode) {
	d.Next = node
}

func NewDoubleNode(pre, next *DoubleNode, data interface{}) *DoubleNode {
	return &DoubleNode{Data: data, Prev: pre, Next: next}
}

type DoubleList struct {
	head *DoubleNode //头指针 //头指针不放数据
	tail *DoubleNode //尾指针
	size int32       //链表长度
}

func (l *DoubleList) AppendNode(values ...interface{}) {

	//头指针单独处理
	if l.size <= 0 {
		element := NewDoubleNode(nil, nil, nil)

		l.head = element
		l.tail = element
	}
	for _, value := range values {
		node := NewDoubleNode(l.tail, nil, value)
		tail := l.tail
		tail.Next = node
		node.Prev = tail
		l.tail = node

		l.size++
	}
}

func (l *DoubleList) Traverse() {
	cur := l.head
	for cur != nil {
		if cur.Data != nil {
			fmt.Println("从头部遍历数据", cur.Data)
		}
		cur = cur.Next
	}

	fmt.Println("------分界线-----------")
	cur = l.tail
	for cur != nil {
		if cur.Data != nil {
			fmt.Println("从尾部遍历数据", cur.Data)
		}
		cur = cur.Prev
	}
}

func (l *DoubleList) InsertNode(value interface{}, size int) {

	cur := l.head
	newNode := &DoubleNode{Data: value}

	var csize = 0
	for cur.Next != nil {
		if csize == size {
			newNode.Prev = cur
			newNode.Next = cur.Next
			cur.Next.Prev = newNode
			l.size++
			break
		}
		cur = cur.Next
		csize++
	}

}

// NewDoubleList
// 初始化双向链表
func NewDoubleList() *DoubleList {
	return &DoubleList{}
}
