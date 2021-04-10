package test

import (
	"leetCodeByGoLand/dataStructure/array"
	"leetCodeByGoLand/util/testTool"
	"testing"
)

// 测试添加到 header 位置
func TestSequenceListAddAtHead(t *testing.T) {
	list := array.NewSequenceList()

	list.AddAtTail(1, 2, 3, 5)

	list.AddAtHead("我是头")

	item, err := list.GetByIndex(0)
	if err != nil {
		panic(err)
	}

	testTool.EqualTest(t, item, "我是头")
}
