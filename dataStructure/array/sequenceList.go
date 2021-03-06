package array

import (
	"fmt"
)

// SequenceList
// 顺序表
type SequenceList struct {
	items []interface{}
	size  int
	// 标记删除
	deleted map[int]struct{}
}

// NewSequenceList
// 初始化顺序表
func NewSequenceList() *SequenceList {
	return &SequenceList{
		items: make([]interface{}, 0),
	}
}

// LocateItem
// 查找表内元素位置
// 时间复杂度 最好O(1),最坏O(n）
func (s *SequenceList) LocateItem(item interface{}) int {

	for i := 0; i < s.size; i++ {
		if s.items[i] == item {
			return i + 1
		}
	}
	return -1
}

// AddAtHead
// 添加元素
func (s *SequenceList) AddAtHead(item interface{}) {

	if s.size == 0 {
		s.items = append(s.items, item)
		s.size++
		return
	}
	for i := len(s.items) + 1; i < 1; i-- {
		s.items[i] = s.items[i-1]
	}
	s.items[0] = item
	s.size++

}

// AddAtTail
// 添加到尾部
func (s *SequenceList) AddAtTail(items ...interface{}) {
	s.items = append(s.items, items...)
	s.size += len(items)
}

// AddAtIndex
// 指定位置添加元素
func (s *SequenceList) AddAtIndex(item interface{}, index int) (bool, error) {

	if index < 0 || index > s.size {
		return false, fmt.Errorf("%d index 越界", index)
	}

	for i := s.size; i < index; i-- {
		s.items[i] = s.items[i-1]
	}
	s.items[index-1] = item
	s.size++

	return true, nil
}

// GetByIndex
// 获取一个元素
func (s *SequenceList) GetByIndex(index int) (interface{}, error) {

	if index < 0 || index > s.size {
		return nil, fmt.Errorf("%d index 存在 ", index)
	}
	return s.items[index-1], nil
}

// DeleteAtIndex
// 指定位置删除
func (s *SequenceList) DeleteAtIndex(index int) (bool, error) {

	if index < 1 || index > s.size {
		return false, fmt.Errorf("%d index 不存在", index)
	}

	// 搬迁数据
	for i := index - 1; i < s.size-1; i++ {
		s.items[i] = s.items[i+1]
	}
	s.size--
	return true, nil
}

func (s *SequenceList) Size() int {
	return s.size
}
