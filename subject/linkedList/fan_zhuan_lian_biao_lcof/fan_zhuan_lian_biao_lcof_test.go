package fan_zhuan_lian_biao_lcof

import (
	"fmt"
	"leetCodeByGoLand/dataStructure/linkedList/linkedList"
	"testing"
)

var (
	list = linkedList.SinglyLinkList{}
	head *linkedList.Node
)

func init() {
	for i := 0; i < 5; i++ {
		node := &linkedList.Node{Data: i}
		if head == nil {
			head = node // 注释
		}
		list.Append(node)
	}
}

// 迭代法反转链表
func TestIteration(t *testing.T) {

	fmt.Println("-------迭代法反转链表-------")
	// 原链表
	list.Traverse(head)
	fmt.Println("---------反转后--------")
	// 反转后
	list.Traverse(reverseListIteration(head))
}

// 反转链表 数组法
func TestByArray(t *testing.T) {

	fmt.Println("--------数组法反转链表--------")
	fmt.Println()
	fmt.Println("-原链表-")
	fmt.Println()

	list.Traverse(head)
	p := reverseListByArray(head)

	fmt.Println()
	fmt.Println("-反转后链表-")
	fmt.Println()

	list.Traverse(p)
}

// 递归法反转链表
func TestReverseListByRecursion(t *testing.T) {
	// 构建链表
	list := linkedList.SingleList{}
	list.BatchAppend(
		linkedList.NewSingleNode("测试1"),
		linkedList.NewSingleNode("测试2"),
		linkedList.NewSingleNode("测试3"),
	)

	fmt.Println("-------原链表------")
	f := func(curNode *linkedList.SingleNode) {
		fmt.Println(curNode.Item)
	}
	list.Iterate(f)

	// 递归
	nodeHead := reverseListByRecursion(list.GetHead())
	fmt.Println("----反转后链表-----")
	list.Traverse(nodeHead, func(node *linkedList.SingleNode) {
		fmt.Println(node.Item)
	})

}
