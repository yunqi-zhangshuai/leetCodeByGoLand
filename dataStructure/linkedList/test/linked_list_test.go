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

// 单链表创建,追加节点,删除节点,遍历,在指定位置插入节点,获取指定位置node
func TestSingleList(t *testing.T) {
	list := linkedList.SingleList{}
	for i := 0; i < 10; i++ {
		list.Append(linkedList.NewSingleNode(i))
	}

	f := func(curNode *linkedList.SingleNode) {
		fmt.Println(curNode.Item)
	}
	list.Iterate(f)

	var position uint32 = 1
	_, err := list.RemoveNode(position)
	if err != nil {
		panic(err)
	}

	fmt.Printf("删除%d节点\n", position)
	list.Iterate(f)
	fmt.Println("链表尾节点-----", list.GetTail().Item)

	// 指定位置插入节点
	fmt.Println("-------指定在第三节点位置插入数据------")
	list.InsertNode(3, &linkedList.SingleNode{
		Item: 98,
	})

	list.Iterate(f)

	node, err := list.GetNode(45)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(node)
	}

}
