package test

import (
	"leetCodeByGoLand/dataStructure/stack"
	"leetCodeByGoLand/util/testTool"
	"testing"
)

func TestStackByLinkedList(t *testing.T) {
	listStack := stack.NewListStack()
	listStack.Push("我")
	listStack.Push("是")
	listStack.Push("中")
	listStack.Push("国")
	listStack.Push("人")

	for {
		_, err := listStack.Pop()
		if err != nil {
			testTool.Test(t, err, testTool.BeNil)
			break
		}

	}

}
