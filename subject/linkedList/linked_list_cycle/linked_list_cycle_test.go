package linked_list_cycle

import (
	"leetCodeByGoLand/dataStructure/linkedList"
	"testing"
)

func TestHasCycle(t *testing.T) {
	// 构建一个环形链表

	list := linkedList.SingleList{}
	list.BatchAppend(
		linkedList.NewSingleNode("链表是否有环呢...."),
		linkedList.NewSingleNode(1),
		linkedList.NewSingleNode(2),
		linkedList.NewSingleNode(3),
	)

	list.GetTail().Next = list.GetHead()
	//list.Iterate(func(curNode *linkedList.SingleNode) {
	//	fmt.Println(curNode.Item)
	//})

	// 链表是否有环,map 法
	if hasCycleByMap(list.GetHead()) {
		t.Log("链表有环")
	}

	// 双指针

}
