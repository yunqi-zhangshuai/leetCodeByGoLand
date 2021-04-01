package main

import (
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
		list.Append(&linkedList.Node{i, nil})
	}
	t.Log("------------------------后续追加遍历-------------------------------")

	list.Traverse(nil)                                 //遍历链表
	list.BeforeInsert(&linkedList.Node{"www", nil}, 5) //在第几个之前插入链表
	t.Log("------------------------节点之前插入遍历------------------------------")

	list.Traverse(nil)
	list.RemoveNode(4)
	t.Log("-------------------------删除某个节点遍历-----------------------------")
	list.Traverse(nil)
}
