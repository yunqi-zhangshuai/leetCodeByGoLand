package main

import (
	"fmt"
	"leetCodeByGoLand/dataStructure/linkedList"
	"testing"
)

func TestSinglyLinkList(t *testing.T) {
	singlyLinkList(t)

}
func TestDoubleLinkList(t *testing.T) {
	//单向链表实现操作
	//双向链表实现操作
	double := linkedList.Double{}
	double.AppendNode(1, 3, 4, 5, 6, 7)
	double.Traverse()

}

func singlyLinkList(t *testing.T) {
	list := &linkedList.SinglyLinkList{}

	for i := 0; i < 10; i++ {
		list.Append(&linkedList.Node{Data: i})
	}
	t.Log("------------------------后续追加遍历-------------------------------")

	list.Traverse(nil)                                  //遍历链表
	list.BeforeInsert(&linkedList.Node{Data: "www"}, 5) //在第几个之前插入链表
	t.Log("------------------------节点之前插入遍历------------------------------")

	list.Traverse(nil)
	list.RemoveNode(4)
	t.Log("-------------------------删除某个节点遍历-----------------------------")
	list.Traverse(nil)
}

func TestSingleList(t *testing.T) {
	list := linkedList.SingleList{}
	node := linkedList.NewSingleNode(1)
	list.Append(node)
	node1 := linkedList.NewSingleNode(2)
	list.Append(node1)
	head := list.GetHead()

	p := head

	for p != nil {
		fmt.Println(p.Item)
		p = p.Next
	}

	fmt.Println(list.GetTail().Item)

}
