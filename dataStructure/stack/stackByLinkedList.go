package stack

import (
	"errors"
	"leetCodeByGoLand/dataStructure/linkedList/linkedList"
)

// ListStack
// 链表实现栈
type ListStack struct {
	items  *linkedList.SingleList
	top    *linkedList.SingleNode
	bottom *linkedList.SingleNode
}

// NewListStack
// 初始化栈
func NewListStack() *ListStack {
	return &ListStack{items: linkedList.NewSingleList()}
}

// Push
// 入栈
func (ls *ListStack) Push(value interface{}) {
	node := linkedList.NewSingleNode(value)
	ls.items.AddAtHead(node)
	ls.top = ls.items.GetHead()
}

// Pop
// 出栈
func (ls *ListStack) Pop() (interface{}, error) {

	if ls.Empty() {
		return nil, errors.New("stack is empty")
	}

	tmpPop := ls.top
	ls.items.DeleteAtIndex(1)
	ls.top = ls.items.GetHead()
	return tmpPop.Item, nil
}

func (ls *ListStack) Empty() bool {
	return ls.items.Size() <= 0
}
