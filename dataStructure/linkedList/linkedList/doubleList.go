package linkedList

import (
	"fmt"
	"runtime"
)

// DoubleNode  双向链表 node
type DoubleNode struct {
	prev *DoubleNode //上一个节点
	next *DoubleNode //指向下一个节点
	Data interface{} //数据
}

// SetPrev
// 设置前前趋节点
func (d *DoubleNode) SetPrev(node *DoubleNode) {
	d.prev = node
}

// SetNext
// 设置后趋node
func (d *DoubleNode) SetNext(node *DoubleNode) {
	d.next = node
}

func (d *DoubleNode) Next() *DoubleNode {
	return d.next
}

func (d *DoubleNode) Prev() *DoubleNode {
	return d.prev
}

func NewDoubleNode(pre, next *DoubleNode, data interface{}) *DoubleNode {
	return &DoubleNode{Data: data, prev: pre, next: next}
}

/*---------------------------------------------------------*/

// DoubleList
// 双向链表
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
		tail.next = node
		node.prev = tail
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
		cur = cur.next
	}

	fmt.Println("------分界线-----------")
	cur = l.tail
	for cur != nil {
		if cur.Data != nil {
			fmt.Println("从尾部遍历数据", cur.Data)
		}
		cur = cur.prev
	}
}

func (l *DoubleList) AddAtIndex(value interface{}, index int) {

	cur := l.head
	newNode := NewDoubleNode(nil, nil, value)

	var sort = 0
	for cur.next != nil {
		if sort == index {
			newNode.SetPrev(cur)
			newNode.SetNext(cur.Next())
			cur.Next().SetPrev(newNode)
			l.size++
			break
		}
		cur = cur.next
		sort++
	}

}

// IsEmpty
// 判断链表是否为空
func (l *DoubleList) IsEmpty() bool {
	return l.head.next == nil
}

// Size
// 链表长度
func (l *DoubleList) Size() int32 {
	return l.size

}

// Clear
// 删除链表
func (l *DoubleList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
	runtime.GC()
}

// NewDoubleList
// 初始化双向链表
func NewDoubleList() *DoubleList {
	return &DoubleList{}
}
