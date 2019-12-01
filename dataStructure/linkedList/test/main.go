package main

import (
	"fmt"
	"leetCodeByGoLand/dataStructure/linkedList"
)

func main() {
	//单向链表实现操作
	//singlyLinkList()
	//双向链表实现操作
	double := linkedList.Double{}
	double.AppendNode(1,3,4567)
	double.Traverse()

}

func singlyLinkList() {
	list := &linkedList.SinglyLinkList{}

	for i := 0; i < 10; i++ {
		list.Append(&linkedList.Node{i, nil})
	}
	fmt.Println("------------------------后续追加遍历-------------------------------")

	list.Traverse(nil)                                 //遍历链表
	list.BeforeInsert(&linkedList.Node{"www", nil}, 5) //在第几个之前插入链表
	fmt.Println("------------------------节点之前插入遍历------------------------------")

	list.Traverse(nil)
	list.RemoveNode(4)
	fmt.Println("-------------------------删除某个节点遍历-----------------------------")
	list.Traverse(nil)
}
