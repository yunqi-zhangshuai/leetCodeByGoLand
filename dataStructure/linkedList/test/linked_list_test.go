package main

import (
	"fmt"
	"leetCodeByGoLand/dataStructure/linkedList/linkedList"
	"testing"
)

func TestDoubleLinkList(t *testing.T) {
	//双向链表实现操作
	double := linkedList.DoubleList{}
	double.AppendNode(1, 3, 4, 5, 6, 7)
	double.Traverse()

}

// 单链表创建,追加节点,删除节点,遍历,在指定位置插入节点,获取指定位置node
func TestSingleList(t *testing.T) {
	list := linkedList.SingleList{}
	for i := 0; i < 3; i++ {
		list.AddAtTail(linkedList.NewSingleNode(i))
	}

	f := func(curNode *linkedList.SingleNode) {
		fmt.Println(curNode.Item)
	}
	list.Iterate(f)

	//fmt.Printf("删除%d节点\n", position)
	list.Iterate(f)
	fmt.Println("链表尾节点-----", list.GetTail())

	// 指定位置插入节点
	fmt.Println("-------指定在第2节点位置插入数据------")
	list.AddAtIndex(2, &linkedList.SingleNode{
		Item: 98,
	})

	var position uint32 = 4
	_, err := list.DeleteAtIndex(position)
	if err != nil {
		panic(err)
	}

	list.Iterate(f)

	node, err := list.Get(45)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(node)
	}

}

// 递归遍历链表
func TestTravelByRecursion(t *testing.T) {
	list := linkedList.SingleList{}
	list.BatchAppend(
		linkedList.NewSingleNode(1),
		linkedList.NewSingleNode(2),
		linkedList.NewSingleNode(3),
	)

	Recursion(list.GetHead())

}

func Recursion(head *linkedList.SingleNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Item)
	Recursion(head.Next)
}
