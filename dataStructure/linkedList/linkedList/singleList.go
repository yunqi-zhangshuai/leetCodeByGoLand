package linkedList

import (
	"errors"
	"fmt"
)

// SingleNode
// 链表节点
type SingleNode struct {
	Item interface{}
	Next *SingleNode
}

// NewSingleNode
// new 链表节点
func NewSingleNode(item interface{}) *SingleNode {
	return &SingleNode{Item: item}
}

//----------------------------------------------------------------------

// SingleList
// 单链表
type SingleList struct {
	// 链表头尾节点
	head, tail *SingleNode
	size       uint32
}

// Append
// 追加到链表尾部
func (l *SingleList) Append(node *SingleNode) {
	if l.head == nil {
		l.head = node
		l.tail = node
		l.size++
		return
	}
	l.tail.Next = node
	l.size++
	l.tail = node
}

// GetHead
// 获取链表头
func (l *SingleList) GetHead() *SingleNode {
	return l.head
}

// GetTail
// 获取链表尾部节点
func (l *SingleList) GetTail() *SingleNode {
	return l.tail
}

// RemoveNode
// 删除指定位置链表节点
func (l *SingleList) RemoveNode(position uint32) (bool, error) {

	if position > l.size {
		return false, errors.New(fmt.Sprintf(
			"超过了链表的长度 %d", l.size,
		))
	}
	curNode := l.head

	if position == 1 {
		tmp := l.head
		l.head = tmp.Next
		tmp = nil
		l.size--
		return true, nil
	}

	var index uint32
	for curNode != nil {
		index++
		if index == position {
			tmpNext := curNode.Next
			curNode.Next = tmpNext.Next
			tmpNext = nil
			l.size--
			break
		}
		curNode = curNode.Next
	}

	return true, nil
}

// InsertNode
// 指定位置插入节点
func (l *SingleList) InsertNode(sort uint32, node *SingleNode) (bool, error) {
	cur := l.head

	if sort < 1 || sort > l.size {
		return false, errors.New(fmt.Sprintf("sort %d 边界错误", sort))
	}
	var index uint32
	for cur != nil {
		index++
		if index == sort {
			tmpNext := cur.Next.Next
			cur.Next = node
			cur.Next.Next = tmpNext
			l.size++
			return true, nil
		}
		cur = cur.Next
	}

	return true, nil
}

// Length
// 获取链表长度
func (l *SingleList) Length() uint32 {
	return l.size
}

// BatchAppend
// 批量追加链表节点
func (l *SingleList) BatchAppend(nodes ...*SingleNode) {
	for _, node := range nodes {
		l.Append(node)
	}
}

// Iterate
// 迭代
func (l *SingleList) Iterate(f func(curNode *SingleNode)) {

	curNode := l.head
	for curNode != nil {
		f(curNode)
		curNode = curNode.Next
	}
}

func (l *SingleList) GetNode(sort uint32) (*SingleNode, error) {
	var index uint32
	cur := l.head

	for cur != nil {
		index++
		if index == sort {
			return cur, nil
		}
		cur = cur.Next
	}
	return nil, errors.New(fmt.Sprintf("不存在此节点 %d", sort))
}

func (l *SingleList) Traverse(head *SingleNode, f func(node *SingleNode)) {
	cur := head
	for cur != nil {
		if f != nil {
			f(cur)
		}
		cur = cur.Next
	}
}
