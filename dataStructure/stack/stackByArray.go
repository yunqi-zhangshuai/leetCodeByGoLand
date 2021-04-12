package stack

import "leetCodeByGoLand/dataStructure/linkedList/linkedList"

type ListStack struct {
	item *linkedList.SingleList
	top  *linkedList.SingleNode
}

func NewListStack() *ListStack {
	return &ListStack{item: linkedList.NewSingleList()}
}

func (ls *ListStack) Push() {

}

func (ls *ListStack) Pop() {

}

func (ls *ListStack) Empty() bool {

	return true
}
