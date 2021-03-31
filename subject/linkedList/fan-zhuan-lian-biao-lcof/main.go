package main

import (
	"fmt"
	"leetCodeByGoLand/dataStructure/linkedList"
)

// 反转链表
func main() {
	// 初始化一个链表
	list := linkedList.SinglyLinkList{}

	var head *linkedList.Node
	for i := 0; i < 5; i++ {
		node := &linkedList.Node{Data: i}
		if head == nil {
			head = node // 注释
		}
		list.Append(node)
	}
	list.Traverse(head)
	fmt.Println("---------------------")
	// 迭代法
	//list.Traverse(reverseList1(head))
	// 数组法
	p := reverseListByArray(head)
	for p != nil {
		fmt.Println(p.Data)
		p = p.Next
	}
	return
	list.Traverse(reverseListByArray(head))
}

// 数组法反转链表,将链表遍历放到数组里,然后倒序遍历数组,串联链表
func reverseListByArray(head *linkedList.Node) *linkedList.Node {
	cur := head
	nodes := make([]*linkedList.Node, 0)
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	node := nodes[4]
	nodes[0].Next = nil
	p := node

	for i := 3; i >= 0; i-- {
		p.Next = nodes[i]
		p = p.Next
	}

	return node
}

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
//
//
// 示例:
// =====
//  输入: 1->2->3->4->5->NULL
//  输出: 5->4->3->2->1->NULL
// =====
//
//  限制：
//
//  0 <= 节点个数 <= 5000

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList1(head *linkedList.Node) *linkedList.Node {

	cur := head
	var pre, tmp *linkedList.Node

	for cur != nil {

		tmp = cur.Next //  tmp = 4

		cur.Next = pre //  cur.next = 2

		pre = cur //  pre = 4
		cur = tmp //  cur = 4
	}

	return pre
}
