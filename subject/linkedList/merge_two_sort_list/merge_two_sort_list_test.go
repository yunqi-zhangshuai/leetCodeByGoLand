package merge_two_sort_list

import (
	"fmt"
	"leetCodeByGoLand/dataStructure/linkedList/linkedList"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		l1 *linkedList.SingleNode
		l2 *linkedList.SingleNode
	}

	ls1 := linkedList.NewSingleList()
	ls1.BatchAppend(
		linkedList.NewSingleNode(1),
		linkedList.NewSingleNode(2),
		linkedList.NewSingleNode(4),
	)

	ls2 := linkedList.NewSingleList()
	ls2.BatchAppend(
		linkedList.NewSingleNode(1),
		linkedList.NewSingleNode(3),
		linkedList.NewSingleNode(4),
	)

	node := MergeTwoLists(ls1.GetHead(), ls2.GetHead())
	head := node
	for head != nil {
		fmt.Println(head.Item)
		head = head.Next
	}
}
