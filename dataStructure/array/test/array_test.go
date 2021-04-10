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

// 指定位置添加
func TestAddAtIndex(t *testing.T) {
	list := array.NewSequenceList()

	list.AddAtTail(1, 2, 3, 5)

	_, err := list.AddAtIndex("小明", 2)

	testTool.Test(t, err, testTool.BeNil)

	item, err := list.GetByIndex(2)
	testTool.Test(t, err, testTool.BeNil)

	testTool.EqualTest(t, item, "小明")

}

func TestDeleteAtIndex(t *testing.T) {
	list := array.NewSequenceList()

	list.AddAtTail(1, 3, 6, 5)

	_, err := list.DeleteAtIndex(2)

	item1, _ := list.GetByIndex(2)
	testTool.EqualTest(t, item1, 6)
	testTool.Test(t, err, testTool.BeNil)

	list.DeleteAtIndex(1)

	first, _ := list.GetByIndex(1)
	testTool.EqualTest(t, first, 6)

}
