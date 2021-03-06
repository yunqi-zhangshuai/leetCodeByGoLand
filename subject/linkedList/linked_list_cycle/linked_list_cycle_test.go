package linked_list_cycle

import (
	"leetCodeByGoLand/dataStructure/linkedList/linkedList"
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
		t.Log("链表有环---map")
	} else {
		t.Error("链表无环---map")
	}

	// 双指针
	if hasCycleFastSlowPoint(list.GetHead()) {
		t.Log("链表有环--快慢指针")
	} else {
		t.Error("链表无环--快慢指针")
	}

}
