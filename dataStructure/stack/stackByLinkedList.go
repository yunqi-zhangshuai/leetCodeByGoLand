package stack

import (
	"errors"
	"leetCodeByGoLand/dataStructure/linkedList/linkedList"
)

type ListStack struct {
	items  *linkedList.SingleList
	top    *linkedList.SingleNode
	bottom *linkedList.DoubleNode
}

func NewListStack() *ListStack {
	return &ListStack{items: linkedList.NewSingleList()}
}

func (ls *ListStack) Push(value interface{}) {
	ls.items.AddAtTail(linkedList.NewSingleNode(value))
	ls.top = ls.items.GetHead()
}

func (ls *ListStack) Pop() (interface{}, error) {

	if ls.Empty() {
		return nil, errors.New("stack is empty")
	}

	tmpPop := ls.top
	ls.top = tmpPop.Next
	return tmpPop.Item, nil
}

func (ls *ListStack) Empty() bool {
	return ls.bottom.Next() == nil
}
