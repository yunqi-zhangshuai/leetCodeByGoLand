package stack

import (
	"errors"
)

type ArrayStack struct {
	items  []interface{}
	top    int
	bottom int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		items:  make([]interface{}, 0),
		bottom: -1,
	}
}

// Push
// 推入元素
func (as *ArrayStack) Push(val interface{}) {
	as.items = append(as.items, val)
	as.topIncr()
}

// Pop
// 弹出数据
func (as *ArrayStack) Pop() (interface{}, error) {
	if as.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	index := as.top
	defer as.topDecr()
	return as.items[index-1], nil
}

// IsEmpty
// 是否为空
func (as *ArrayStack) IsEmpty() bool {
	return as.top <= as.bottom
}

// topIncr
func (as *ArrayStack) topIncr() {
	as.top++
}

// topDecr
func (as *ArrayStack) topDecr() {
	as.top--
}
