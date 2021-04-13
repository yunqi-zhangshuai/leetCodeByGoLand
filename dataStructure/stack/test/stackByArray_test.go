package test

import (
	"leetCodeByGoLand/dataStructure/stack"
	"leetCodeByGoLand/util/testTool"
	"testing"
)

func TestPop(t *testing.T) {
	stackArr := stack.NewArrayStack()
	stackArr.Push(1)
	stackArr.Push(2)
	stackArr.Push(3)
	stackArr.Push(4)
	//
	firstPop, err := stackArr.Pop()
	testTool.Test(t, err, testTool.BeNil)
	testTool.EqualTest(t, firstPop, 4)
	twoPop, err := stackArr.Pop()
	testTool.Test(t, err, testTool.BeNil)

	testTool.EqualTest(t, twoPop, 3)

}
