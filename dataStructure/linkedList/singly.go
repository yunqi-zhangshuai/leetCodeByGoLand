package linkedList

import "fmt"

/*
单链表
*/

//节点
type Node struct {
	Data interface{} //数据快
	Next *Node       //下一个节点指针
}

type SinglyLinkList struct {
	head    *Node //头指针
	tail    *Node //尾指针
	size    int32 //链表长度
	current *Node //当前节点指针
}

//在链表末尾后面添加结点
func (l *SinglyLinkList) Append(nodeData *Node) {

	var cur *Node
	if l.current == nil {
		l.head = nodeData
		cur = l.head
	} else {
		cur = l.current
	}
	cur.Next = nodeData
	l.current = cur.Next
	l.tail = cur.Next
	l.size++
}

//遍历
func (l *SinglyLinkList) Traverse(p *Node) {
	var cur *Node
	if p == nil {
		cur = l.head
	} else {
		cur = p
	}
	for cur.Next != nil {
		cur = cur.Next
		fmt.Println("当前节点数据是---", cur.Data)
	}
}

//在某个节点之后插入数据
func (l *SinglyLinkList) BeforeInsert(n *Node, s int32) {

	current := l.head
	var i int32 = 1
	for current.Next != nil {
		if i == s {
			n.Next = current.Next //最重要的地方 ,先把要插入节点的next 指针指定为当前节点的next 指针,再将当前Next指针指向n
			current.Next = n
			l.size++
			break
		}
		current = current.Next
		i++
	}

}

//删除某个节点
func (l *SinglyLinkList) RemoveNode(s int32) {
	current := l.head
	var i int32 = 1

	for current.Next != nil {
		if i == s {
			current.Next = current.Next.Next
			l.size--
			break
		}
		current = current.Next
		i++
	}

}
