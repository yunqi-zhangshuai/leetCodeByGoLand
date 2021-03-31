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
	list.Traverse(reverseList(head))
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
func reverseList(head *linkedList.Node) *linkedList.Node {

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
